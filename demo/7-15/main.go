package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	//testCancelSleep()
	testCancelAfter()
}

func testCancelSleep() {
	t := time.Now()
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Cancelled", time.Now())
				return
			default:
				fmt.Println("Sleep begin", time.Now())
				time.Sleep(time.Second * 10)
				fmt.Println("Sleep end", time.Now())
			}
		}
	}(ctx)
	fmt.Println("Cancel begin", time.Now())
	cancel()
	fmt.Println("Cancel end", time.Now())
	wg.Wait()
	fmt.Println("time cost", time.Since(t).Milliseconds())
}

func testCancelAfter() {
	t := time.Now()
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			fmt.Println("for执行了")
			select {
			case <-ctx.Done():
				fmt.Println("Cancelled", time.Now())
				return
			case <-time.After(time.Second * 10):
				fmt.Println("time.After done", time.Now())
				return
			}
		}
	}(ctx)

	//defer cancel()
	time.Sleep(time.Second)
	fmt.Println("Cancel begin", time.Now())
	cancel()
	fmt.Println("Cancel end", time.Now())
	wg.Wait()
	fmt.Println("time cost", time.Since(t).Milliseconds())
}
