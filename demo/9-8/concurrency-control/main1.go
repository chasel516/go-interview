package main

import (
	"log"
)
import "time"

type Request interface{}

func handle(r Request) {
	log.Println("发起请求", r)
}

const RateLimitPeriod = time.Second
const RateLimit = 1 // 任何一秒钟内最多处理1个请求
func handleRequests(requests <-chan Request) {
	quotas := make(chan time.Time, RateLimit)

	//相当于以恒定的速率放入令牌
	go func() {
		tick := time.NewTicker(RateLimitPeriod / RateLimit)
		defer tick.Stop()
		for t := range tick.C {
			select {
			case quotas <- t:
			default:
			}
		}
	}()

	for r := range requests {
		<-quotas
		go handle(r)
	}
}

func main() {
	requests := make(chan Request)
	go handleRequests(requests)
	for i := 0; i < 100; i++ {
		requests <- i
	}
}
