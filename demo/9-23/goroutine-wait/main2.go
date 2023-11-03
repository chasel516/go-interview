package main

import (
	"fmt"
	"time"
)

func main() {
	//无缓冲的channel
	ch := make(chan struct{})
	var v1, v2 int
	go func() {
		v2 = getData1()
		//获取完数据发送消息，通知主协程继续执行
		ch <- struct{}{}
	}()
	v1 = getData()
	//能读取说明	getData1()已经执行结束了
	<-ch
	fmt.Println(v1, v2)
}

func getData() int {
	time.Sleep(2 * time.Second)
	return 1
}

func getData1() int {
	time.Sleep(1 * time.Second)
	return 2
}
