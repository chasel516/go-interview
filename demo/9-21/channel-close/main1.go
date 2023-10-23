package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
		close(ch1)
	}()
	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- 2
		close(ch2)
	}()
	for {
		select {
		case x, ok := <-ch1:
			fmt.Println(x, ok)
			//if !ok {
			//	ch1 = nil
			//} else {
			//	fmt.Println(x)
			//}
		case x, ok := <-ch2:
			fmt.Println(x, ok)
			//if !ok {
			//	ch2 = nil
			//} else {
			//	fmt.Println(x)
			//}
		}
		if ch1 == nil && ch2 == nil {
			break
		}
	}
	fmt.Println("program end")
}
