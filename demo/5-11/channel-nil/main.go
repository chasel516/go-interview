package main

import (
	"fmt"
	"time"
)

func main() {
	//channelIsNil()
	//f1()
	//f2()
	f3()

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
	go func() {
		fmt.Println("开始发送")
		c <- 1
		fmt.Println("阻塞在发送") //永远不会执行
	}()
	go func() {
		fmt.Println("开始接收")
		fmt.Println(<-c)
		fmt.Println("阻塞在接收") //永远不会执行
	}()
	time.Sleep(time.Second)
}

func f3() {
	var c chan int
	close(c) // panic: close of nil channel
}
