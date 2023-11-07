package main

import (
	"channel-close/channel"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch := channel.NewChannel()
	stop := make(chan struct{})
	wg := sync.WaitGroup{}
	go receive1(ch, stop)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go send1(ch, stop, &wg, i)
	}
	wg.Wait()
	ch.Close()

	time.Sleep(time.Second)

}

// 多个发送者对一个接收者时，由接收者关闭一个额外通道来通知发送者停止发送数据
func send1(ch *channel.Channel, stop chan struct{}, wg *sync.WaitGroup, index int) {
	defer wg.Done()
	for {
		select {
		case <-stop:
			fmt.Println("退出send", index)
			return
		case ch.C <- rand.Intn(10):
		}
	}

}

func receive1(ch *channel.Channel, stop chan struct{}) {
	i := 0
	//close(ch.C)
	//直到ch的缓冲队列为空且已关闭才会退出循环
	for v := range ch.C {
		if i == 10 {
			close(stop)
			fmt.Println("close")
		}
		fmt.Println("receive v=", v)
		i++

	}

	fmt.Println("退出receive")
}
