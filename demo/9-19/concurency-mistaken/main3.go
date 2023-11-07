package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ch <- 1
	go fmt.Println(<-ch)

	//go func() {
	//	fmt.Println(<-ch)
	//}()
	time.Sleep(time.Second)
}
