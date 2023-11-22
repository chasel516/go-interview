package main

import (
	"fmt"
)

var globalVar string

func main() {
	_ = globalVar
}

func f1() {
	n1 := 1
	fmt.Println(n1)

	n2 := 1
	_ = n2
}
func f2() {
	var i interface{}
	var n3 int
	i = n3
	_ = i

}

func f3() *int {
	var n4 int
	n4 = 42
	return &n4 // x的地址被返回，导致x逃逸到堆上
}
func f4() {
	var x [1024]int // x的大小超过栈的容量，导致x逃逸到堆上
	x[0] = 42
}

func f5() {
	s1 := make([]byte, 1, 65*1024) // >64k
	_ = s1

	s2 := make([]byte, 1, 64*1024) // 64k
	_ = s2
}

func f6() {
	//同理，string占8字节，当其容量大于8*1024时也会发生逃逸
	//golang在1.3版本之后使用了连续栈来取代了分段栈，这样可以提高栈的访问效率，但是也限制了栈的大小。golang在1.4版本中将连续栈的初始大小设为2kb，
	//如果栈的大小超过了64kb，就会触发栈的扩容，这样会影响性能和内存管理。为了避免频繁的栈扩容，golang采用了一种策略，
	//就是当对象的大小超过64kb时，就直接在堆上分配，而不是在栈上。这样可以减少栈的压力，也可以让堆上的对象由GC来回收。
	s3 := make([]string, 1, 8*1024+1)
	_ = s3

	s4 := make([]string, 1, 8*1024)
	_ = s4
}

func f7() func() int {
	var z int
	z = 42
	return func() int {
		return z // z被分配到一个闭包中，导致z逃逸到堆上
	}
}
func f8() interface{} {
	var str1 string // str1被分配到一个接口中，导致m逃逸到堆上
	return str1
}

func f9() string {
	var str2 string
	return str2
}

func f10() {
	s5 := []interface{}{1, 2}
	s5[0] = 0 //对接口类型的切片元素赋值，切片元素会发生逃逸
}
func f11() {
	s6 := []int{1, 2}
	s6[0] = 0
}
func f12() {
	s7 := make([]interface{}, 2)
	s7[0] = 1 //对接口类型的切片元素赋值，切片元素会发生逃逸
}

func f13() {
	m1 := make(map[int]interface{}, 2)
	m1[0] = 1 //对接口类型的切片元素赋值，切片元素会发生逃逸
}

func f14() {
	m3 := make([]*int, 2)
	i := 1
	m3[0] = &i //切片元素是指针类型时，对切片元素赋值会发生逃逸
}

func f15() {
	m2 := make(map[int]int, 2)
	m2[0] = 1
}

func f16() {
	m3 := make(map[int][]int, 2)
	m3[0] = []int{1} //map元素为切片时，元素赋值时会发生逃逸
}

func f18() {

}

func fn1(param any) any {
	return param
}
