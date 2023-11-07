package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	Ticker(func() {
		fmt.Println("goroutine num:", runtime.NumGoroutine())
	}, time.Second)
}
func main() {
	//ch := make(chan int)
	//<-ch
	//go fmt.Println(<-ch)

	//go func() {
	//	fmt.Println(<-ch)
	//}()
	//ch <- 1

	//close(ch)
	////ch <- 1
	//for {
	//	select {
	//	case c := <-ch:
	//		fmt.Println("执行了...", c)
	//		time.Sleep(time.Second)
	//	}
	//}

	//判断通道是否关闭
	//for {
	//	select {
	//	case c, ok := <-ch:
	//		if !ok {
	//			return
	//		}
	//		fmt.Println("执行了...", c)
	//		time.Sleep(time.Second)
	//	}
	//}

	//time.Sleep(time.Second)

	for i := 0; i < 10000; i++ {
		go request()
	}
	time.Sleep(time.Minute)
}

func request() {
	c := make(chan int, 2)
	for i := 0; i < 3; i++ {
		x := i
		go func() {
			c <- x // 每次调用都会有一个协程阻塞在这里
		}()
	}
}

func Ticker(f func(), d time.Duration) {
	go func() {
		ticker := time.NewTicker(d)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				go f()
			}
		}
	}()
}
