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
				fmt.Println(1)
				time.Sleep(time.Second)
			}
		}
	}(ctx)

	// 当你不再需要这个goroutine时
	cancel()
	select {}
}
