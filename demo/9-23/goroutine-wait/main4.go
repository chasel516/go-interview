package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Ltime)
}
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
			log.Println(x)
		}()
	}

	//获取全部资源，获取不到阻塞
	sema.Acquire(ctx, 3)
	fmt.Println("All workers done")
}
