package main

import (
	"fmt"
	"unsafe"
)

type Person1 struct {
	Name string
	age  int
}

type Person2 struct {
	Name string
}

type Person3 struct {
	uncompare [0]func()
	Name      string
}
type Person4 struct {
	Name string
	age  int
}

func main() {
	p1 := Person1{Name: "test"}
	p2 := Person2{Name: "test"}
	p3 := Person3{Name: "test"}
	fmt.Println("p1.size:", unsafe.Sizeof(p1), "p2.size:", unsafe.Sizeof(p2), "p3.size:", unsafe.Sizeof(p3)) //p1.size: 24 p2.size: 16 p3.size: 16

	//m := map[Person2]int{}
	//m1 := map[Person3]int{} //编译错误

	//set1 := map[int]struct-demo{}{1: {}}
	//set2 := map[int]bool{1: false}
	//fmt.Println("set1.size:", unsafe.Sizeof(set1[1]), "set2.size:", unsafe.Sizeof(set2[1])) //set1.size: 0 set2.size: 1

	//fmt.Println("p1==p2?", p1 == p2)

	//p4 := Person4{Name: "test"}
	//fmt.Println("p1==p4?", p1 == p4)
	p5 := struct {
		Name string
		age  int
	}{}

	p6 := struct {
		Name string
		age  int
	}{}
	p5 = p6
	fmt.Println("p5==p6?", p5 == p6)
}
