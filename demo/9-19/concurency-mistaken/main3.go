package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	//<-ch
	//go fmt.Println(<-ch)

	//go func() {
	//	fmt.Println(<-ch)
	//}()
	//ch <- 1

	//close(ch)
	////ch <- 1
	for {
		select {
		case c := <-ch:
			fmt.Println("执行了...", c)
			time.Sleep(time.Second)
		}
	}

	//判断通道是否关闭
	for {
		select {
		case c, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println("执行了...", c)
			time.Sleep(time.Second)
		}
	}

	//time.Sleep(time.Second)
}

func request() int {
	c := make(chan int, 2)
	for i := 0; i < 3; i++ {
		x := i
		go func() {
			c <- x // 每次调用都会有一个协程阻塞在这里
		}()
	}
	return <-c
}
