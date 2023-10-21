package main

import (
	"log"
	"sync"
	"sync/atomic"
)

type SpinLock struct {
	state int32
}

func (s *SpinLock) Lock() {
	//相当于不停的尝试抢占锁
	for !atomic.CompareAndSwapInt32(&s.state, 0, 1) {
		// 自旋，不执行任何操作
	}
}

func (s *SpinLock) Unlock() {
	atomic.StoreInt32(&s.state, 0)
}

func main() {
	var cnt int
	spinLock := SpinLock{}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			spinLock.Lock()
			cnt++
			spinLock.Unlock()
			wg.Done()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			spinLock.Lock()
			cnt++
			spinLock.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	log.Println("cnt:", cnt)
}
