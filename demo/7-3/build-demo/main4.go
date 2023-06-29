package main

import (
	"fmt"
	"time"
)

func main() {
	x := 1
	y := 2
	go func() {
		fmt.Println(x, y)
	}()
	//time.Sleep(time.Millisecond)
	y = 3
	time.Sleep(time.Millisecond)
}
