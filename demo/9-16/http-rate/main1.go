package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	response := make(chan int, 1)
	done := make(chan bool)
	go func() {
		response <- request(1, done)
	}()
	go func() {
		response <- request(2, done)
	}()
	go func() {
		response <- request(3, done)
	}()
	//获取最先响应的请求
	res := <-response
	close(done)
	fmt.Println(res)
	time.Sleep(time.Second)
}

func request(param int, done chan bool) int {
	requestDone := make(chan bool, 1)
	resp := 0
	cost := rand.Int63n(100)
	go func() {
		fmt.Println("param", param, "cost", cost)
		time.Sleep(time.Duration(cost) * time.Millisecond)
		resp = param
		//标记请求完成
		requestDone <- true

	}()

	select {
	case <-done:
		fmt.Println("请求被取消", param)
		return 0
	case <-requestDone:
		return resp

	}
}
