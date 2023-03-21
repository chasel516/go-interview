package main

import (
	"fmt"
)

func main() {
	//a1 := [3]int{}
	//a2 := [3]int{}
	//fmt.Println(a1 == a2) //true
	//
	//a3 := [3][]int{}
	//a4 := [3][]int{}
	//fmt.Println(a3 == a4) //编译不通过
	////
	//m := map[float64]int{}
	//m[1.1] = 1
	//m[1.2] = 2
	//m[0.3] = 5
	//m[0.30000000000000001] = 6
	//fmt.Printf("m:%+v\n", m)
	////
	//fmt.Println(math.Float64bits(0.3))
	//fmt.Println(math.Float64bits(0.30000000000000001))
	//
	//m[math.NaN()] = 3
	//m[math.NaN()] = 4
	//fmt.Println("m[math.NaN]=", m[math.NaN()])
	//fmt.Println("math.NaN() == math.NaN()?", math.NaN() == math.NaN())
	//for k, v := range m {
	//	fmt.Println("k:", k, "v:", v)
	//}
	p1 := new(int)
	p2 := new(int)
	fmt.Println(p1, p2)
	//
	m1 := map[*int]int64{}
	m1[p1] = 1
	m1[p2] = 2
	fmt.Printf("m1:%+v \n", m1)
}
