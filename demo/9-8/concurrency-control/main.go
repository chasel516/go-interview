//上面已经展示了如何使用尝试发送实现峰值限制。 同样地，我们也可以使用使用尝试机制来实现速率限制，但需要前面刚提到的定时器实现的配合。 速率限制常用来限制吞吐和确保在一段时间内的资源使用不会超标。
//下面的例子借鉴了官方Go维基中的例子 � 。 在此例中，任何一分钟时段内处理的请求数不会超过200。
package main

import "fmt"
import "time"

type Request interface{}

func handle(r Request) { fmt.Println(r.(int)) }

const RateLimitPeriod = time.Minute
const RateLimit = 200 // 任何一分钟内最多处理200个请求11|
func handleRequests(requests <-chan Request) {
	quotas := make(chan time.Time, RateLimit)

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
	// time.Sleep(time.Minute)
	for i := 0; ; i++ {
		requests <- i
	}
}

//上例的代码虽然可以保证任何一分钟时段内处理的请求数不会超过200，但是如果在开始的一分钟
//内没有任何请求，则接下来的某个瞬时时间点可能会同时处理最多200个请求（试着将
//time.Sleep行的注释去掉看看）。 这可能会造成卡顿情况。我们可以将速率限制和峰值限制一并使用来避免出现这样的情况。
