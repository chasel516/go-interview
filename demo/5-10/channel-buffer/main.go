package main

import (
	"fmt"
	"time"
)

func main() {
	//f1()
	//f2()
	f3()
	//var ch chan struct{}
	//
	//for i := 0; i < 3; i++ {
	//	go test()
	//}
	//fmt.Println("永久阻塞")
	//<-ch
}

func f1() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)
}

func f2() {
	ch := make(chan int)
	fmt.Println(<-ch)
	go func() {
		ch <- 1
	}()
}

func f3() {
	ch := make(chan int)
	go func() {
		fmt.Println(<-ch)
	}()

	go func() {
		ch <- 1
	}()
	time.Sleep(time.Second)
}

func test() {
	for {
		time.Sleep(time.Second)
	}
}
