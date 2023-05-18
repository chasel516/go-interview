package main

import (
	"github.com/panjf2000/ants/v2"
	"log"
	"runtime"
	"sync/atomic"
	"time"
)

func init() {
	Ticker(func() {
		log.Print("current go routine num: ", runtime.NumGoroutine())
	}, 2*time.Minute)
}
func main() {
	// 创建一个Ants池，最多允许5个goroutine同时执行
	p, _ := ants.NewPool(5)
	defer p.Release()

	// 定义一个计数器，用于统计处理的任务数
	var counter int64

	// 向Ants池提交100个任务
	for i := 0; i < 100; i++ {
		x := i
		err := p.Submit(func() {
			// 模拟任务处理时间
			time.Sleep(time.Second)
			log.Println("执行任务", x)
			// 增加计数器的值
			atomic.AddInt64(&counter, 1)
		})
		if err != nil {
			log.Println("Submit error", err)
		}
	}

	// 等待任务完成(可能不包含最后一批任务)
	p.Waiting()
	// 输出计数器的值
	log.Println("counter:", counter)

	//当Running的任务数为0说明任务全部执行完
	for p.Running() > 0 {
		time.Sleep(100 * time.Millisecond)
	}

}

// 监控系统当前的协程数量：
// 启动一个定时器
func Ticker(f func(), d time.Duration) {
	ticker := time.NewTicker(d)
	for {
		select {
		case <-ticker.C:
			go f()
		}
	}
}
