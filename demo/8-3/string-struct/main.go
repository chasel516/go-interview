package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s1 := "abc"
	//s2 := "中文"
	//fmt.Println(len(s1), unsafe.Sizeof(s1))
	//fmt.Println(len(s2), unsafe.Sizeof(s2))

	//fmt.Println(UnicodeLen(s1), UnicodeLen(s2))

	//for i, b := range s2 {
	//	fmt.Println(i, b)
	//}
	//fmt.Println("======================")
	//for i, c := range []rune(s2) {
	//	fmt.Println(i, c)
	//}
	//
	//fmt.Println("======================")
	//for i, c := range []rune(s2) {
	//	fmt.Printf("i:%d;c:%s", i, string(c))
	//	fmt.Printf("i:%d;c:%c", i, c)
	//}

	ptr := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	fmt.Println(ptr)
	//字符串Data的地址
	fmt.Println(unsafe.Pointer(ptr.Data))

	//获取Data地址所指向的数据
	fmt.Println((*[3]byte)(unsafe.Pointer(ptr.Data))) //&[97 98 99]

	//s1[1] = 'd'
	s1 = "adcd"
}

func UnicodeLen(s string) int {
	r := []rune(s)
	return len(r)
}
