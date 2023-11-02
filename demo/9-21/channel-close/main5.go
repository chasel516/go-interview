package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	const sendNem = 5
	ch := make(chan int)
	stop := make(chan struct{})
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		go receive2(ch, i)
	}

	for i := 0; i < sendNem; i++ {
		wg.Add(1)
		go send2(ch, stop, &wg, i)
	}

	//让消息持续发送一段时间
	time.Sleep(time.Second)

	//通知发送者关闭
	close(stop)

	wg.Wait()
	//发送者全部关闭后,关闭数据channel
	close(ch)

	time.Sleep(time.Second)

}

// 多个发送者对一个接收者时，由接收者关闭一个额外通道来通知发送者停止发送数据
func send2(ch chan int, stop chan struct{}, wg *sync.WaitGroup, index int) {
	defer wg.Done()
	for {
		select {
		case <-stop:
			fmt.Println("退出send", index)
			//不能在有多个发送者的情况下关闭channel，当其中一个发送者关闭channel后其他发送者在还没来得及收到退出消息时可能还会继续发送数据
			//ch.Close()
			return

		case ch <- rand.Intn(10):
		}
	}

}

func receive2(ch chan int, index int) {
	//直到ch的缓冲队列为空且已关闭才会退出循环
	for v := range ch {
		fmt.Println("receive v=", v)
	}

	fmt.Println("退出receive-", index)
}
