package main

import (
	"github.com/phper95/pkg/routine"
	"log"
	"sync/atomic"
	"time"
)

const PoolNameTest = "test-pool"

func init() {
	routine.InitPoolWithName(PoolNameTest, 1, 5, 0)
}
func main() {

	// 定义一个计数器，用于统计处理的任务数
	var counter int64
	for i := 0; i < 100; i++ {
		x := i
		routine.GetPool(PoolNameTest).Put(func() {
			// 模拟任务处理时间
			time.Sleep(time.Second)
			log.Println("执行任务", x)
			// 增加计数器的值
			atomic.AddInt64(&counter, 1)
		})
		for routine.GetPool(PoolNameTest).QueueLen() > 0 {
			time.Sleep(100 * time.Millisecond)
		}
		// 输出计数器的值
		log.Println("counter:", counter)
	}
}
