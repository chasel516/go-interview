package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func init() {
	go Ticker2(func() {
		log.Print("current go routine num: ", runtime.NumGoroutine())
	}, time.Second)
}

func main() {
	for i := 0; i < 100000; i++ {
		go func(index int) {
			select {}
			fmt.Println("end", index)
		}(i)
	}
	time.Sleep(time.Minute)
}

func Ticker2(f func(), d time.Duration) {
	ticker := time.NewTicker(d)
	for {
		select {
		case <-ticker.C:
			go f()
		}
	}
}
