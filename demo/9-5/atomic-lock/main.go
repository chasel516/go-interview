package main

import (
	"context"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func init() {
	log.SetFlags(log.Ltime)
}
func main() {
	//waitGroup()
	//semaphoreWait()
	rateLimit()

}

func waitGroup() {
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		x := i
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			time.Sleep(time.Second)
			log.Println(x)
		}()
	}
	wg.Wait()
}

func semaphoreWait() {
	sema := semaphore.NewWeighted(3)
	ctx := context.Background()
	for i := 0; i < 3; i++ {
		x := i
		sema.Acquire(ctx, 1)
		go func() {
			defer func() {
				sema.Release(1)
			}()
			time.Sleep(time.Second)
			log.Println(x)
		}()
	}

	//获取全部资源，获取不到阻塞
	sema.Acquire(ctx, 3)
}

func rateLimit() {
	sema := semaphore.NewWeighted(3)
	wg := sync.WaitGroup{}
	ctx := context.Background()
	for i := 0; i < 100; i++ {
		x := i
		sema.Acquire(ctx, 1)
		wg.Add(1)
		go func() {
			defer func() {
				sema.Release(1)
				wg.Done()
			}()
			time.Sleep(time.Second)
			log.Println(x)
		}()
	}
	wg.Wait()
}
