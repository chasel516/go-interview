package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("a.len", len(a), "a.cap", cap(a))
	b := []int{1, 2, 3}
	fmt.Println("b.len", len(b), "b.cap", cap(b))
	b = append(b, 4)
	fmt.Println("b.len", len(b), "b.cap", cap(b))
}
