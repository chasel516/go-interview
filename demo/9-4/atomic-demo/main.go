package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//runtime.GOMAXPROCS(1)
	//counter1()
	//counter2()
	//counter3()
	counter4()
}

func counter1() {
	cnt := 0
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			cnt++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(cnt)
}

func counter2() {
	cnt := 0
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			lock.Lock()
			defer func() {
				lock.Unlock()
				wg.Done()
			}()
			cnt++

		}()
	}
	wg.Wait()
	fmt.Println(cnt)
}

func counter3() {
	cnt := 0
	wg := sync.WaitGroup{}
	//声明一个带缓冲的channel，这里的作用跟锁类似，利用了channel本身是并发安全的特性
	ch := make(chan struct{}, 1)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			ch <- struct{}{}
			defer func() {
				<-ch
			}()
			cnt++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(cnt)
}

func counter4() {
	var cnt int64
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&cnt, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(cnt)
}
