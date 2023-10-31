package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch)
	close1(ch)
}

func close1(ch chan int) (closed bool) {
	defer func() {
		if recover() != nil {
			fmt.Println("has been closed")
		}
	}()
	close(ch)
	return true
}
