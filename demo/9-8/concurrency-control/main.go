package main

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"time"
)

var limit = rate.Every(time.Second)

var limiter = rate.NewLimiter(limit, 3)

// 参数r  每秒可以向 Token 桶中产生多少 token
// 参数b Token 桶的容量大小
// 允许最大突发流程是100，限制每秒10个请求
// var limiter = rate.NewLimiter(10, 100)
func main() {
	for i := 0; i < 100; i++ {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
		err := limiter.Wait(ctx)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		log.Println("发起请求", i)
	}
}
