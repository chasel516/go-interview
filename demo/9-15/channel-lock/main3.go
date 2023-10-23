package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

// 限制同时执行的任务数
var w = runtime.GOMAXPROCS(-1)
var semaC = make(chan struct{}, w) //信号量

func main() {
	counter := int32(0)
	for i := 0; i < 1000; i++ {
		// 如果没有信号量可用，会阻塞在这里，直到某个task被释放
		semaC <- struct{}{}
		// 启动 goroutine
		go func(i int) {
			defer func() {
				<-semaC
			}()
			time.Sleep(20 * time.Millisecond) // 模拟一个耗时操作
			atomic.AddInt32(&counter, 1)
		}(i)
	}
	// 获取所有的信号量(获取不到阻塞)，确保所有的goroutine执行完成
	for i := 0; i < w; i++ {
		semaC <- struct{}{}
	}
	fmt.Println(counter)
}
