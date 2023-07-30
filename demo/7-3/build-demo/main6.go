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
	x = 2
	z := x + y
	fmt.Println(z)
	time.Sleep(time.Millisecond)
}
