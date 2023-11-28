package main

import (
	"fmt"
	"reflect"
)

func main() {}

// 函数返回值为接口类型且发生了类型转化会导致返回值发生逃逸
func f10() interface{} {
	var str1 string // str1被分配到一个接口中，导致m逃逸到堆上
	return str1
}

func f11() interface{} {
	var str2 interface{}
	return str2 //直接返回一个接口类型，没有发生类型转化则不会逃逸

}

func f12() {
	var str3 interface{}
	str3 = "test" //interface只做类型推断时不会逃逸
	_ = str3

}

func f13() interface{} {
	var str4 interface{}
	str4 = "test" //从string类型转interface，发生了逃逸
	return str4
}

// 容器类型的元素为接口类型时，对接口类型进行赋值会导致容器的元素发生逃逸；
func f14() {
	s5 := []interface{}{1, 2}
	s5[0] = 0 //对接口类型的切片元素赋值，切片元素会发生逃逸
}

func f15() {
	m3 := make([]*int, 2)
	i := 1 //切片元素是指针类型时，对切片元素赋值时，该切片元素会发生逃逸
	m3[0] = &i
}

func f16() {
	s6 := []int{1, 2} //非接口类型的切片元素没有发生逃逸
	s6[0] = 0
}

func f17() {
	m1 := make(map[int]interface{}, 2)
	m1[0] = 1 //对接口类型的map元素赋值，元素会发生逃逸
}

func f18() {
	m2 := make(map[int]int, 2)
	m2[0] = 1 //非接口类型的map的元素没有发生逃逸
}

func f19() {
	m3 := make(map[int][]int, 2)
	m3[0] = []int{1} //map元素为切片时，元素赋值时会发生逃逸
}

func f20() {
	ch := make(chan interface{})
	ch <- 1 //接口类型的通道元素发送到通道时也会发生逃逸

	ch1 := make(chan *int)
	ch1 <- new(int)
}

func f21() {
	var n21 int
	reflect.TypeOf(n21)
	reflect.ValueOf(n21) //反射获取值的时候会发生逃逸
}

func f22() {
	var str5 = "hello" //str5作为标准库函数的入参导致逃逸
	str6 := fmt.Sprintf("%s world", str5)
	str6 += " "
	var str7 = "test" //str5作为标准库函数的入参导致逃逸
	fmt.Println(str7)
}
