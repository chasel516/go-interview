package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func init() {
	go Ticker1(func() {
		log.Print("current go routine num: ", runtime.NumGoroutine())
	}, time.Second)
}

func main() {
	var ch chan int //只声明而没有通过make或者字面量的形式完成初始化
	for i := 0; i < 100000; i++ {
		go func(index int) {
			ch <- 1
			//<-ch
			fmt.Println("end", index)
		}(i)
	}
	<-ch
	//ch <- 1
}

func Ticker1(f func(), d time.Duration) {
	ticker := time.NewTicker(d)
	for {
		select {
		case <-ticker.C:
			go f()
		}
	}
}
