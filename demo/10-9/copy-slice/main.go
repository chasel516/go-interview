package main

import "fmt"

func main() {
	//f1()
	//f2()
	//f3()
	f4()
	//f5()
}

func f1() {
	s1 := []int{3: 100}
	s2 := s1[:4]
	fmt.Println(s2)
	s1[0] = 1
	fmt.Println(s2)
	s2[0] = 2
	fmt.Println(s1[0])
}

func f2() {
	var s1 []int
	s1 = []int{1, 2}
	s2 := []int{3, 4, 5}
	affectItems := copy(s2, s1)
	fmt.Println(s2, affectItems) //[1 2 5] 2
}

func f3() {
	var s1 []int
	s2 := []int{3, 4, 5}
	affectItems := copy(s2, s1)
	fmt.Println(s2, affectItems)
}

func f4() {
	s1 := []int{3: 100}
	s2 := []int{3, 4, 5}
	affectItems := copy(s2, s1)
	fmt.Println(s2, affectItems)
}

func f5() {
	s1 := []int{1, 2, 3}
	s2 := make([]int, len(s1))
	copy(s2, s1)
	fmt.Println(s2)
}
