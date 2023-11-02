package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("end")
	time.Sleep(time.Second)
}
