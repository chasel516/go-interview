package main

import (
	"fmt"
	"time"
)

func main() {
	channelIsNil()
	//f1()
	//f2()
	//f3()
	//f4()
	//f5()
	//f6()
	f7()

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
	ch := make(chan int)
	close(ch)
	go func() {
		<-ch
	}()
	ch <- 1
}

func f2() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for x := range ch {
		fmt.Println(x)
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
