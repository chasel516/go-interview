package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//NonReentrant()
	//OnlyOneWriteLock()
	//lock := sync.Mutex{}
	//for {
	//	//未解锁将在此处阻塞
	//	lock.Lock()
	//	fmt.Println("LockandUnlockIndiffGoroutine-main:获取写锁")
	//	go LockandUnlockIndiffGoroutine(&lock)
	//	time.Sleep(1)
	//}

	//BlockWhenWriteLock()
	//BlockWhenWriteLock1()
	//BlockWhenReadLock()

	ControlGoroutineOrder()
	//ControlGoroutineOrder1()
	//ControlGoroutineOrder2()

	//UnlockCausePanic()

}

// 1. 不可重入性
func NonReentrant() {
	rw := sync.RWMutex{}
	rw.RLock()
	fmt.Println("获取读锁")
	rw.Lock() //读锁未解锁的情况下操作获取写锁会导致panic
	fmt.Println("获取写锁")
	rw.RUnlock()
	rw.Unlock()
}

func OnlyOneWriteLock() {
	lock := sync.RWMutex{}
	//lock := sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()
	fmt.Println("获取写锁")
	lock.Lock() //写锁未解锁的情况下会阻塞
	fmt.Println("获取写锁")
	defer lock.Unlock()
}

func LockandUnlockIndiffGoroutine(lock *sync.Mutex) {
	time.Sleep(time.Second)
	lock.Unlock()
	fmt.Println("LockandUnlockIndiffGoroutine:释放写锁")
}

func BlockWhenWriteLock() {
	lock := sync.RWMutex{}
	lock.Lock()
	fmt.Println("BlockWhenWriteLock 加锁")
	go func() {
		fmt.Println("BlockWhenWriteLock -goroutine开始执行")
		time.Sleep(3 * time.Second)
		fmt.Println("BlockWhenWriteLock -goroutine执行结束")
		lock.Unlock()
	}()
	//写锁解锁前会阻塞
	lock.Lock()
	fmt.Println("BlockWhenWriteLock结束")

}

func BlockWhenWriteLock1() {
	lock := sync.RWMutex{}
	lock.Lock()
	fmt.Println("BlockWhenWriteLock1 加锁")
	go func() {
		fmt.Println("BlockWhenWriteLock1 -goroutine开始执行")
		time.Sleep(3 * time.Second)
		fmt.Println("BlockWhenWriteLock1 -goroutine执行结束")
		lock.Unlock()
	}()
	//写锁解锁前会阻塞
	lock.RLock()
	fmt.Println("BlockWhenWriteLock1结束")

}

func BlockWhenReadLock() {
	lock := sync.RWMutex{}
	lock.RLock()
	fmt.Println("BlockWhenReadLock 加锁")
	go func() {
		fmt.Println("BlockWhenReadLock -goroutine开始执行")
		time.Sleep(3 * time.Second)
		fmt.Println("BlockWhenReadLock -goroutine执行结束")
		lock.RUnlock()
	}()
	//读锁解锁前会阻塞
	lock.Lock()
	fmt.Println("BlockWhenReadLock结束")
}

func ControlGoroutineOrder() {
	lock := sync.RWMutex{}
	lock.Lock()
	go func() {
		lock.Lock()
		fmt.Println("ControlGoroutineOrder:第一个协程执行")
		lock.Unlock()
	}()
	time.Sleep(time.Second)
	go func() {
		lock.Unlock()
		fmt.Println("ControlGoroutineOrder:第二个协程执行")
	}()
	time.Sleep(time.Millisecond)
}

func ControlGoroutineOrder1() {
	lock := sync.RWMutex{}
	lock.Lock()
	go func() {
		lock.RLock()
		fmt.Println("ControlGoroutineOrder1:第一个协程执行")
		lock.RUnlock()
	}()
	time.Sleep(time.Second)
	go func() {
		lock.Unlock()
		fmt.Println("ControlGoroutineOrder1:第二个协程执行")
	}()
	time.Sleep(time.Millisecond)
}

func ControlGoroutineOrder2() {
	lock := sync.RWMutex{}
	lock.RLock()
	go func() {
		lock.Lock()
		fmt.Println("ControlGoroutineOrder2:第一个协程执行")
		lock.Unlock()
	}()
	time.Sleep(time.Second)
	go func() {
		lock.RUnlock()
		fmt.Println("ControlGoroutineOrder2:第二个协程执行")
	}()
	time.Sleep(time.Millisecond)
}

func UnlockCausePanic() {
	lock := sync.Mutex{}
	//lock := sync.RWMutex{}
	lock.Unlock()
	//lock.RUnlock()

}
