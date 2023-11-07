package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//1. 如何确保通道关闭后不再向通道中发送数据？
	//2. 如何确保通道关闭时，通道中的数据全部处理完了？
	ch := make(chan int)
	wg := sync.WaitGroup{}
	go send(ch)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go receive(ch, &wg, i)
	}
	wg.Wait()

}

// 一个发送者对多个接收者时，由发送者关闭通道
func send(ch chan int) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Millisecond * 500)
		ch <- i
	}
	close(ch)
	fmt.Println("close")
}

func receive(ch chan int, wg *sync.WaitGroup, index int) {
	defer wg.Done()

	//直到ch的缓冲队列为空且已关闭才会退出循环
	for v := range ch {
		fmt.Printf("receive-%d:v=%d\n", index, v)
	}
	fmt.Println("退出receive-", index)
}
