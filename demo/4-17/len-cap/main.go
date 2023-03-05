package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a [5]int
	fmt.Println("a.len", len(a), "a.cap", cap(a))
	b := []int{1, 2, 3}
	fmt.Println("b.len", len(b), "b.cap", cap(b))
	b = append(b, 4)
	fmt.Println("b.len", len(b), "b.cap", cap(b))

	m := map[string]int{}
	m["k1"] = 1
	fmt.Println("m.len", len(m))

	ch := make(chan int, 10)
	fmt.Println("ch.len", len(ch), cap(ch))

	//数组的长度和容量在编译阶段被估值
	const arrLen = len(a)
	const arrCap = cap(a)

	//通过反射可以修改切片的长度
	s := make([]string, 1, 6)

	//注意，len要比cap小
	reflect.ValueOf(&s).Elem().SetLen(2)
	fmt.Println("s.len", len(s), "s.cap", cap(s))
	//cap只能小于等于当前值
	reflect.ValueOf(&s).Elem().SetCap(5)
	fmt.Println("s.len", len(s), "s.cap", cap(s))

	str := "go面试"
	fmt.Println("str.len", len(str)) //unicode编码中，一个中文字符占3个字节
	fmt.Println("str.runeLen", len([]rune(str)))

}
