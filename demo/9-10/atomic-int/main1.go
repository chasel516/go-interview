package main

import (
	"log"
	"sync"
	"sync/atomic"
)

// 自旋锁本质上是一种无锁的结构，在需要改变数据的时候，反复判断数据是否和原数据一致
// 一致则替换，不一致时说明已经被其他Goroutine修改，则再次进入for循环判断。
// 自旋的代价仅仅是再进行一次for循环进行判断，相比于将协程挂起进行等待代价更小。
type spinLock int32

func main() {
	var wg sync.WaitGroup
	var sl spinLock

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer func() {
				wg.Done()
				sl.Unlock()
				log.Println("i= ", index, "Unlock")
			}()
			sl.Lock()
			log.Println("i= ", index, "Lock")

		}(i)
	}

	wg.Wait()
}

func (sl *spinLock) Lock() {
	for {
		//自旋判断lockFlag是否等于0
		if atomic.CompareAndSwapInt32((*int32)(sl), 0, 1) {
			return
		}
	}
}

func (sl *spinLock) Unlock() {
	atomic.StoreInt32((*int32)(sl), 0)
}
