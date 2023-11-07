package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var cnt int32 = 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			//wg.Add(1) //add方法的调用不应该在协程内
			atomic.AddInt32(&cnt, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("cnt=", atomic.LoadInt32(&cnt))
}
