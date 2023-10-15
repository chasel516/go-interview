package main

import (
	"log"
	"sync"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags)
}

var mu sync.Mutex
var cond = sync.NewCond(&mu)
var pendingTask = 0

func main() {
	producer()

	//启动多个消费者消费
	for i := 0; i < 10; i++ {
		consumer()
	}
	select {}

}

func producer() {
	go func() {
		for {
			//wait的调用必须加锁
			cond.L.Lock()
			//任务太多时，暂停生产任务，等待消费者消费任务
			if pendingTask == 10 {
				log.Println("too many task,waiting consume...  pendingTask:", pendingTask)
				cond.Wait()
			}
			cond.L.Unlock()
			log.Println("task produce pendingTask:", pendingTask)
			pendingTask++
			cond.Signal()

		}
	}()
}

func consumer() {
	go func() {
		for {
			//没有任务时等待生产者生产任务
			cond.L.Lock()
			if pendingTask == 0 {
				log.Println("no task,waiting produce...  pendingTask:", pendingTask)
				cond.Wait()
			}
			cond.L.Unlock()
			time.Sleep(time.Second)
			log.Println("task consumer pendingTask:", pendingTask)
			pendingTask--
			cond.Signal()
		}
	}()
}
