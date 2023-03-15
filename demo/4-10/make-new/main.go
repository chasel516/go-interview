package main

import "fmt"

func main() {
	//声明切片slice1并申请地址，为申请的10个元素都赋上元素类型的零值
	//var slice1 = make([]int, 10)
	//fmt.Println(slice1, len(slice1), cap(slice1)) //10 10

	//var slice2 = new([]int)
	//fmt.Println(slice2, len(*slice2), cap(*slice2))

	//s1 := make([]int, 0)
	//fmt.Println("s1", len(s1), cap(s1))
	//s2 := new([]int)
	//s3 := *new([]int)
	//var s4 []int
	//var s5 = []int{}
	////
	//fmt.Println("s1 is nil?", s1 == nil)  //false
	//fmt.Println("s2 is nil?", *s2 == nil) //true
	//fmt.Println("s3 is nil?", s3 == nil)  //true
	//fmt.Println("s4 is nil?", s4 == nil)  //true
	//fmt.Println("s5 is nil?", s5 == nil)  //false

	a1 := *new([10]int)
	a2 := [10]int{}
	fmt.Println(a1, a2)
}
