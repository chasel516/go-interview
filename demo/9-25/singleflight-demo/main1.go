package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/marusama/cyclicbarrier"
)

func init() {
	log.SetFlags(log.LstdFlags)
}
func main() {
	//f1()
	f2()
}

func f1() {
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Millisecond * 500)
			log.Printf("协程 %d 执行结束\n", id)
		}(i)
	}
	wg.Wait()

	for i := 3; i < 6; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Millisecond * 500)
			log.Printf("协程 %d 执行结束\n", id)
		}(i)
	}
	wg.Wait()
}

func f2() {
	ctx := context.Background()
	// 创建循环障栅，设置等待的协程数量为3
	cb := cyclicbarrier.New(3)
	for i := 0; i < 3; i++ {
		go func(id int) {
			time.Sleep(time.Second * time.Duration(id))
			log.Printf("协程 %d 到达栅栏处\n", id)
			//使用 Await() 方法等待其他协程到达障栅处
			err := cb.Await(ctx)
			if err != nil {
				log.Println(err)
			}
			log.Printf("协程 %d 执行结束\n", id)
		}(i)
	}

	//可以继续复用
	for i := 3; i < 6; i++ {
		go func(id int) {
			time.Sleep(time.Second * time.Duration(id))
			log.Printf("协程 %d 到达栅栏处\n", id)
			//使用 Await() 方法等待其他协程到达障栅处
			err := cb.Await(ctx)
			if err != nil {
				log.Println(err)
			}
			log.Printf("协程 %d 执行结束\n", id)
		}(i)
	}

	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	fmt.Println(<-sig)
}
