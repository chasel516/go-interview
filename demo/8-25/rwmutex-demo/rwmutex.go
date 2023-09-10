package main

import (
	"sync/atomic"
	"unsafe"
)

type Mutex struct {
	//表示当前互斥锁的状态,他会被转成二进制位
	//其中包含锁是否为锁定状态、正在等待被锁唤醒的协程数量、两个和饥饿模式有关的标志
	state int32
	//当长时间未获取到锁时，就会使用信号量进行同步。
	//如果加锁操作进入信号量同步阶段，则信号量计数值减1。
	//如果解锁操作进入信号量同步阶段，则信号量计数值加1。
	//当信号量计数值大于0时，意味着有其他协程执行了解锁操作，这时加锁协程可以直接退出。
	//当信号量计数值等于0时，意味着当前加锁协程需要陷入休眠状态
	sema uint32 //用于控制锁状态的信号量
}

type RWMutex struct {
	w           Mutex        // 互斥锁
	writerSem   uint32       // 信号量，写锁等待读取完成
	readerSem   uint32       // 信号量，读锁等待读取完成
	readerCount atomic.Int32 // 当前正在进行读操作的数量
	readerWait  atomic.Int32 // 写操作被阻塞时，等待读操作的数量
}

func (m *Mutex) Lock() {
	// 先尝试快速上锁，当前 state 为 0，说明没人锁。CAS 上锁后直接返回
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		if race.Enabled {
			race.Acquire(unsafe.Pointer(m))
		}
		return
	}
	// 正常情况下会自旋尝试抢占锁一段时间，而不会立即进入休眠状态，
	//这样做的好处是可以让互斥锁在频繁加锁与释放时也能获得较高的效率。
	//锁只有在正常模式下才能够进入自旋状态，
	//runtime_canSpin函数会判断当前是否能进入自旋状态
	m.lockSlow()
}

// 在下面4种情况下，自旋状态立即终止：
// （1）程序在单核CPU上运行。
// （2）逻辑处理器P小于或等于1。
// （3）当前协程所在的逻辑处理器P的本地队列上有其他协程待运行
// （4）自旋次数超过了设定的阈值
func (m *Mutex) lockSlow() {
	var waitStartTime int64
	starving := false
	awoke := false // 被唤醒标记，如果是被别的 goroutine 唤醒的那么后面会置 true
	iter := 0
	old := m.state
	for {

		//runtime_canSpin判断当前是否能进入自旋状态
		if old&(mutexLocked|mutexStarving) == mutexLocked && runtime_canSpin(iter) {
			// Active spinning makes sense.
			// Try to set mutexWoken flag to inform Unlock
			// to not wake other blocked goroutines.
			if !awoke && old&mutexWoken == 0 && old>>mutexWaiterShift != 0 &&
				atomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken) {
				awoke = true
			}
			//进入自旋状态后，runtime_doSpin函数调用的procyield函数是一段汇编代码，
			//会执行30次PAUSE指令占用CPU时间
			runtime_doSpin()
			iter++
			old = m.state
			continue
		}
		new := old
		// Don't try to acquire starving mutex, new arriving goroutines must queue.
		if old&mutexStarving == 0 {
			new |= mutexLocked
		}
		if old&(mutexLocked|mutexStarving) != 0 {
			new += 1 << mutexWaiterShift
		}
		// The current goroutine switches mutex to starvation mode.
		// But if the mutex is currently unlocked, don't do the switch.
		// Unlock expects that starving mutex has waiters, which will not
		// be true in this case.
		if starving && old&mutexLocked != 0 {
			new |= mutexStarving
		}
		if awoke {
			// The goroutine has been woken from sleep,
			// so we need to reset the flag in either case.
			if new&mutexWoken == 0 {
				throw("sync: inconsistent mutex state")
			}
			new &^= mutexWoken
		}

		//CAS 更新，如果 m.state 不等于 old，说明有人也在抢锁，那么 for 循环发起新的一轮竞争
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			if old&(mutexLocked|mutexStarving) == 0 {
				break // locked the mutex with CAS
			}
			// If we were already waiting before, queue at the front of the queue.
			queueLifo := waitStartTime != 0
			if waitStartTime == 0 {
				waitStartTime = runtime_nanotime()
			}
			runtime_SemacquireMutex(&m.sema, queueLifo, 1)
			starving = starving || runtime_nanotime()-waitStartTime > starvationThresholdNs
			old = m.state
			if old&mutexStarving != 0 {
				// If this goroutine was woken and mutex is in starvation mode,
				// ownership was handed off to us but mutex is in somewhat
				// inconsistent state: mutexLocked is not set and we are still
				// accounted as waiter. Fix that.
				if old&(mutexLocked|mutexWoken) != 0 || old>>mutexWaiterShift == 0 {
					throw("sync: inconsistent mutex state")
				}
				delta := int32(mutexLocked - 1<<mutexWaiterShift)
				if !starving || old>>mutexWaiterShift == 1 {
					// Exit starvation mode.
					// Critical to do it here and consider wait time.
					// Starvation mode is so inefficient, that two goroutines
					// can go lock-step infinitely once they switch mutex
					// to starvation mode.
					delta -= mutexStarving
				}
				atomic.AddInt32(&m.state, delta)
				break
			}
			// 有人释放了锁，然后当前 goroutine 被 runtime 唤醒了，设置 awoke true
			awoke = true
			iter = 0
		} else {
			old = m.state
		}
	}

	if race.Enabled {
		race.Acquire(unsafe.Pointer(m))
	}
}

func (m *Mutex) Unlock() {
	if race.Enabled {
		_ = m.state
		race.Release(unsafe.Pointer(m))
	}

	// 如果当前锁处于普通的锁定状态，即没有进入饥饿状态和唤醒状态，也没有多个协程因为抢占锁陷入堵塞，
	//那么Unlock方法在修改mutexLocked状态后立即退出（快速路径）。
	//否则，进入慢路径调用unlockSlow方法
	new := atomic.AddInt32(&m.state, -mutexLocked)
	if new != 0 {
		// Outlined slow path to allow inlining the fast path.
		// To hide unlockSlow during tracing we skip one extra frame when tracing GoUnblock.
		m.unlockSlow(new)
	}
}

func (m *Mutex) unlockSlow(new int32) {
	//判断锁是否重复释放。锁不能重复释放，否则会在运行时报错
	if (new+mutexLocked)&mutexLocked == 0 {
		fatal("sync: unlock of unlocked mutex")
	}
	if new&mutexStarving == 0 {
		old := new
		for {
			// 如果锁当前未处于饥饿状态且当前mutexWoken已设置，
			//则表明有其他申请锁的协程准备从正常状态退出，
			//这时锁释放后不用去当前锁的等待队列中唤醒其他协程，而是直接退出。
			//如果唤醒了等待队列中的协程，
			//则将唤醒的协程放入当前协程所在逻辑处理器P的runnext字段中，
			//存储到runnext字段中的协程会被优先调度。

			if old>>mutexWaiterShift == 0 || old&(mutexLocked|mutexWoken|mutexStarving) != 0 {
				return
			}
			// Grab the right to wake someone.
			new = (old - 1<<mutexWaiterShift) | mutexWoken
			if atomic.CompareAndSwapInt32(&m.state, old, new) {
				runtime_Semrelease(&m.sema, false, 1)
				return
			}
			old = m.state
		}
	} else {
		//如果锁当前处于饥饿状态，则进入信号量同步阶段，
		//到全局哈希表中寻找当前锁的等待队列，以先入先出的顺序唤醒指定协程
		//在饥饿模式下，通过将runtime_Semrelease函数第2个参数设置为true，
		//可以让当前协程会让渡自己的执行权利，让被唤醒的协程直接运行
		runtime_Semrelease(&m.sema, true, 1)
	}
}

const rwmutexMaxReaders = 1 << 30

func (rw *RWMutex) RLock() {
	if race.Enabled {
		_ = rw.w.state
		race.Disable()
	}
	//读取操作先通过原子操作将readerCount加1，如果readerCount≥0就直接返回，
	//所以如果只有获取读取锁的操作，那么其成本只有一个原子操作。
	//当readerCount<0时，说明当前有写锁，当前协程将借助信号量陷入等待状态，
	//如果获取到信号量则立即退出，没有获取到信号量时的逻辑与互斥锁的逻辑相似。
	if rw.readerCount.Add(1) < 0 {
		// A writer is pending, wait for it.
		runtime_SemacquireRWMutexR(&rw.readerSem, false, 0)
	}
	if race.Enabled {
		race.Enable()
		race.Acquire(unsafe.Pointer(&rw.readerSem))
	}
}

func (rw *RWMutex) RUnlock() {
	if race.Enabled {
		_ = rw.w.state
		race.ReleaseMerge(unsafe.Pointer(&rw.writerSem))
		race.Disable()
	}
	//读锁解锁时，如果当前没有写锁，则其成本只有一个原子操作并直接退出
	if r := rw.readerCount.Add(-1); r < 0 {
		rw.rUnlockSlow(r)
	}
	if race.Enabled {
		race.Enable()
	}
}

// Lock locks rw for writing.
// If the lock is already locked for reading or writing,
// Lock blocks until the lock is available.
func (rw *RWMutex) Lock() {
	if race.Enabled {
		_ = rw.w.state
		race.Disable()
	}
	// 读写锁申请写锁时要调用Lock方法，必须先获取互斥锁，因为它复用了互斥锁的功能。
	rw.w.Lock()
	//接着readerCount减去rwmutexMaxReaders阻止后续的读操作。
	//但获取互斥锁并不一定能直接获取写锁，如果当前已经有其他Goroutine持有互斥锁的读锁，
	//那么当前协程会加入全局等待队列并进入休眠状态，当最后一个读锁被释放时，会唤醒该协程。
	r := rw.readerCount.Add(-rwmutexMaxReaders) + rwmutexMaxReaders
	// Wait for active readers.
	if r != 0 && rw.readerWait.Add(r) != 0 {
		runtime_SemacquireRWMutex(&rw.writerSem, false, 0)
	}
	if race.Enabled {
		race.Enable()
		race.Acquire(unsafe.Pointer(&rw.readerSem))
		race.Acquire(unsafe.Pointer(&rw.writerSem))
	}
}

func (rw *RWMutex) Unlock() {
	if race.Enabled {
		_ = rw.w.state
		race.Release(unsafe.Pointer(&rw.readerSem))
		race.Disable()
	}

	// 将readerCount加上rwmutexMaxReaders，表示不会堵塞后续的读锁
	r := rw.readerCount.Add(rwmutexMaxReaders)
	if r >= rwmutexMaxReaders {
		race.Enable()
		fatal("sync: Unlock of unlocked RWMutex")
	}
	// 依次唤醒所有等待中的读锁。当所有的读锁唤醒完毕后会释放互斥锁
	for i := 0; i < int(r); i++ {
		runtime_Semrelease(&rw.readerSem, false, 0)
	}
	// 调用互斥锁的Unlock方法释放互斥锁
	rw.w.Unlock()
	if race.Enabled {
		race.Enable()
	}
}

// TryRLock tries to lock rw for reading and reports whether it succeeded.
//
// Note that while correct uses of TryRLock do exist, they are rare,
// and use of TryRLock is often a sign of a deeper problem
// in a particular use of mutexes.
func (rw *RWMutex) TryRLock() bool {
	if race.Enabled {
		_ = rw.w.state
		race.Disable()
	}
	for {
		c := rw.readerCount.Load()
		if c < 0 {
			if race.Enabled {
				race.Enable()
			}
			return false
		}
		if rw.readerCount.CompareAndSwap(c, c+1) {
			if race.Enabled {
				race.Enable()
				race.Acquire(unsafe.Pointer(&rw.readerSem))
			}
			return true
		}
	}
}

// 如果当前有写锁正在等待，则调用rUnlockSlow判断当前是否为最后一个被释放的读锁，
// 如果是则需要增加信号量并唤醒写锁。
func (rw *RWMutex) rUnlockSlow(r int32) {
	if r+1 == 0 || r+1 == -rwmutexMaxReaders {
		race.Enable()
		fatal("sync: RUnlock of unlocked RWMutex")
	}
	// A writer is pending.
	if rw.readerWait.Add(-1) == 0 {
		// The last reader unblocks the writer.
		runtime_Semrelease(&rw.writerSem, false, 1)
	}
}

// TryLock tries to lock rw for writing and reports whether it succeeded.
//
// Note that while correct uses of TryLock do exist, they are rare,
// and use of TryLock is often a sign of a deeper problem
// in a particular use of mutexes.
func (rw *RWMutex) TryLock() bool {
	if race.Enabled {
		_ = rw.w.state
		race.Disable()
	}
	if !rw.w.TryLock() {
		if race.Enabled {
			race.Enable()
		}
		return false
	}
	if !rw.readerCount.CompareAndSwap(0, -rwmutexMaxReaders) {
		rw.w.Unlock()
		if race.Enabled {
			race.Enable()
		}
		return false
	}
	if race.Enabled {
		race.Enable()
		race.Acquire(unsafe.Pointer(&rw.readerSem))
		race.Acquire(unsafe.Pointer(&rw.writerSem))
	}
	return true
}

// RLocker returns a Locker interface that implements
// the Lock and Unlock methods by calling rw.RLock and rw.RUnlock.
func (rw *RWMutex) RLocker() Locker {
	return (*rlocker)(rw)
}

type rlocker RWMutex

func (r *rlocker) Lock()   { (*RWMutex)(r).RLock() }
func (r *rlocker) Unlock() { (*RWMutex)(r).RUnlock() }
