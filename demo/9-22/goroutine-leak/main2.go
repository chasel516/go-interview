package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		defer func() {
			fmt.Println("子协程退出")
		}()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				// do something
			}
		}
	}(ctx)

	// 当你不再需要这个goroutine时
	cancel()
	time.Sleep(3 * time.Second)
}
