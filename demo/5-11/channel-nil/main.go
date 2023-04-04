package main

import (
	"fmt"
)

func main() {
	//channelIsNil()
	//f1()
	f2()
	//f3()
	//f4()
	//f5()
	//f6()
	//f7()

	var c chan int
	var c2 chan<- int
	var c3 <-chan int
	c2 = c
	c3 = c
	c = c2  //编译不通过
	c2 = c3 //编译不通过
	fmt.Println(c2, c3)

}
func channelIsNil() {
	var ch chan struct{}
	var ch1 = make(chan struct{})
	fmt.Println("ch == nil:", ch == nil)
	fmt.Println("ch1 == nil:", ch1 == nil)
	close(ch1)
	fmt.Println("关闭后ch1 == nil:", ch1 == nil)
}

func f1() {
	var c chan int
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}

func f2() {
	var c chan int
	for {
		select {
		case c <- 1:
			fmt.Println("send nil channel")
		}
	}
}

func f3() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		select {
		case x, ok := <-ch:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Println(x)
		}
	}
}

func f4() {
	var c chan int
	close(c) // panic: close of nil channel
}

func f5() {
	c := make(chan int)
	defer close(c)
	go func() {
		c <- 1
		close(c)
	}()

	fmt.Println(<-c)

}
func f6() {
	c := make(chan int)
	defer close(c)
	go func() {
		c <- 1
	}()

	fmt.Println(<-c)
}

func f7() {
	c := make(chan int)
	go func() {
		c <- 1
		close(c)
	}()

	fmt.Println(<-c)
}
