package main

import "fmt"

func main() {
	//f6()
	//f7()
	f8()
}

func f6() {
	s1 := []int{1, 2, 3}
	s2 := make([]int, 0, len(s1))
	s2 = append(s2, s1...)
	fmt.Println(s2)
}
func f7() {
	s1 := []int{1, 2, 3}
	s2 := make([]int, len(s1))
	s2 = append(s2, s1...)
	fmt.Println(s2)
}

func f8() {
	var s1 []int
	s2 := append(s1[:0:0], s1...)
	fmt.Println(s2 == nil)
}
