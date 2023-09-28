package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
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
			fmt.Println(x)
		}()
	}

	//获取全部资源，获取不到阻塞
	sema.Acquire(ctx, 3)
}

func waightGroup() {
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		x := i
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			time.Sleep(time.Second)
			fmt.Println(x)
		}()
	}
	wg.Wait()
}

func rateLimit() {
	sema := semaphore.NewWeighted(3)
	ctx := context.Background()
	for i := 0; i < 100; i++ {
		x := i
		sema.Acquire(ctx, 1)
		go func() {
			defer func() {
				sema.Release(1)
			}()
			time.Sleep(time.Second)
			fmt.Println(x)
		}()
	}

}
