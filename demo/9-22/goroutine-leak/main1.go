package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func init() {
	go Ticker(func() {
		log.Print("current go routine num: ", runtime.NumGoroutine())
	}, time.Second)
}

func main() {
	ch := make(chan int)
	for i := 0; i < 100000; i++ {
		go func(index int) {
			ch <- 1
			fmt.Println("end", index)
		}(i)
	}
	<-ch
	select {}
}

func Ticker(f func(), d time.Duration) {
	ticker := time.NewTicker(d)
	for {
		select {
		case <-ticker.C:
			go f()
		}
	}
}
