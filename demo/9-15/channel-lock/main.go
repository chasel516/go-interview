package main

import (
	"fmt"
	"sync"
)

func main() {
	mutex := make(chan struct{}, 1) // 容量必须为1
	counter := 0
	increase := func() {
		mutex <- struct{}{} // 加锁
		counter++
		<-mutex // 解锁
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			increase()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("counter", counter)
}
