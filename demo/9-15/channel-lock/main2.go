package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"log"
	"runtime"
	"sync/atomic"
	"time"
)

var weight = runtime.GOMAXPROCS(-1)
var sema = semaphore.NewWeighted(int64(weight)) //信号量

func main() {
	fmt.Println("weight:", weight)
	ctx := context.Background()
	counter := int32(0)
	for i := 0; i < 1000; i++ {
		// 如果没有信号量可用，会阻塞在这里，直到某个task被释放
		if err := sema.Acquire(ctx, 1); err != nil {
			break
		}
		// 启动 goroutine
		go func(i int) {
			defer sema.Release(1)
			time.Sleep(100 * time.Millisecond) // 模拟一个耗时操作
			atomic.AddInt32(&counter, 1)
		}(i)
	}
	// 获取所有的信号量(获取不到阻塞)，确保所有的goroutine执行完成
	if err := sema.Acquire(ctx, int64(weight)); err != nil {
		log.Printf("获取所有的worker失败: %v", err)
	}
	fmt.Println(counter)
}
