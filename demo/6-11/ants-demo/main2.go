package main

import (
	"ants-demo/gopool"
	"log"
	"sync/atomic"
	"time"
)

func main() {
	p := gopool.NewPool(5, 5, 1) // 限制同时启动5个协程
	// 定义一个计数器，用于统计处理的任务数
	var counter int64

	// 向gopool提交100个任务
	for i := 0; i < 100; i++ {
		x := i
		p.Submit(func() {
			// 模拟任务处理时间
			time.Sleep(time.Second)
			log.Println("执行任务", x)
			// 增加计数器的值
			atomic.AddInt64(&counter, 1)
			//panic(111)
		})

	}

	for p.Running() > 0 {
		time.Sleep(100 * time.Millisecond)
	}

	// 输出计数器的值
	log.Println("counter:", counter)
}
