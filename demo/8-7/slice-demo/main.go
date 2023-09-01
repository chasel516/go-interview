package main

import "log"

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	//arr := [3]int{1, 2, 3}
	//log.Printf("arr:%p", &arr)
	//f1(arr)
	s := []int{1, 2, 3}
	log.Printf("s:%p", s)
	//f2(s)
	//f3(s)
	f4(&s)
	log.Printf("s:%v ptr: %p", s, s)

}

func f1(a [3]int) {
	log.Printf("f1 arr:%p", &a)
}

func f2(param []int) {
	log.Printf("f2 s:%p", param)
	param = append(param, 4)
	log.Printf("f2 重新赋值后: s:%p", param)
}

func f3(param []int) {
	log.Printf("f3 param:%p", param)
	s1 := param
	s1 = append(s1, 7)
	//s1[0] = 0
	copy(param, s1)
	log.Printf("f3 重新赋值后: param: %v ; ptr :  %p", param, param)
}

func f4(param *[]int) {
	log.Printf("f4 s:%p", *param)
	*param = append(*param, 4)
	log.Printf("f4 重新赋值后: s:%p", *param)
}

const N = 128

var x = []int{N - 1: 789}

func f5(param *[]int) {

	y := make([]int, N)
	*(*[N]int)(y) = *(*[N]int)(x)
}
