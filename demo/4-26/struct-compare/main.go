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

	set1 := map[int]struct{}{1: {}}
	set2 := map[int]bool{1: false}
	fmt.Println("set1.size:", unsafe.Sizeof(set1[1]), "set2.size:", unsafe.Sizeof(set2[1])) //set1.size: 0 set2.size: 1

	// fmt.Println("p1==p2?", p1 == p2)

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

/*
可以看出 Go 语言中 unsafe.Sizeof() 函数：

1、对不同长度的字符串，unsafe.Sizeof() 函数的返回值都为 16，这是因为 string 类型对应一个结构体，该结构体有两个域，第一个域指向该字符串的指针，第二个域为字符串的长度，每个域占 8 个字节，但是并不包含指针指向的字符串的内容，这就解释了unsafe.Sizeof() 函数对 string 类型的返回值始终是16。
2、对不同长度的数组，unsafe.Sizeof() 函数的返回值随着数组中的元素个数的增加而增加，这是因为unsafe.Sizeof() 函数总是在编译期就进行求值，而不是在运行时，这就意味着，unsafe.Sizeof() 函数的返回值可以赋值给常量，在编译期求值，意味着可以获得数组所占的内存大小，因为数组总是在编译期就指明自己的容量，并且在以后都是不可变的。
3、对所含元素个数不同的切片，unsafe.Sizeof() 函数的返回值都为 24，这是因为对切片来说，unsafe.Sizeof() 函数返回的值对应的是切片的描述符，而不是切片所指向的内存的大小，因此都是24。

*/
