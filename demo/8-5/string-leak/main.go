package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

var s = strings.Repeat("1", 1<<20)

func main() {
	ptr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	//字符串字节数组的的地址
	fmt.Println("s pointer:", unsafe.Pointer(ptr.Data))

	//Assign()
	//AssignPointer()
	//StringSlice()
	//Repeat()
	StringSlice1(s)
	StringSlice2(s)
	//StringSliceUseBuilder(s)
	//f1(s)
	//f2(&s)
}

func Assign() {
	s2 := s
	ptr := (*reflect.StringHeader)(unsafe.Pointer(&s2))
	//字符串字节数组的的地址
	fmt.Println("Assign:", unsafe.Pointer(ptr.Data))

	//将原始字符串置空
	s := "" //字符串s的底层指向改变
	ptr = (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println("s pointer:", unsafe.Pointer(ptr.Data))

	ptr = (*reflect.StringHeader)(unsafe.Pointer(&s2))
	//字符串字节数组的的地址
	fmt.Println("Assign:", unsafe.Pointer(ptr.Data))
	_ = s2
}

func AssignPointer() {

	s2 := &s

	ptr := (*reflect.StringHeader)(unsafe.Pointer(s2))
	//字符串字节数组的的地址
	fmt.Println("AssignPointer", unsafe.Pointer(ptr.Data))
	_ = s2
}

func StringSlice() {
	s2 := s[:20]
	ptr := (*reflect.StringHeader)(unsafe.Pointer(&s2))
	//字符串字节数组的的地址
	fmt.Println("StringSlice", unsafe.Pointer(ptr.Data))
	_ = s2

}

func Repeat() {

	//s2 := strings.Repeat(s, 1)
	s2 := strings.Repeat(s, 2)
	ptr := (*reflect.StringHeader)(unsafe.Pointer(&s2))
	//字符串字节数组的的地址
	fmt.Println("Repeat", unsafe.Pointer(ptr.Data))
	_ = s2

}

func f1(s string) string {
	ptr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	//字符串字节数组的的地址
	fmt.Println("f1:", unsafe.Pointer(ptr.Data))
	return s
}

func f2(s *string) *string {
	ptr := (*reflect.StringHeader)(unsafe.Pointer(s))
	//字符串字节数组的的地址
	fmt.Println("f2:", unsafe.Pointer(ptr.Data))
	return s
}

func StringSlice1(s string) string {
	s1 := []byte(s[:20])
	ptr := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	//字符串字节数组的的地址
	fmt.Println("StringSlice1:", unsafe.Pointer(ptr.Data))
	s2 := string(s1)
	ptr = (*reflect.StringHeader)(unsafe.Pointer(&s2))
	//字符串字节数组的的地址
	fmt.Println("StringSlice1:", unsafe.Pointer(ptr.Data))
	return s2
}

func StringSlice2(s string) string {
	s1 := string([]byte(s[:20]))
	ptr := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	//字符串字节数组的的地址
	fmt.Println("StringSlice2:", unsafe.Pointer(ptr.Data))

	return s1
}

func StringSlice3(s string) string {
	s1 := (" " + s[:20])[1:]

	ptr := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	//字符串字节数组的的地址
	fmt.Println("StringSlice3:", unsafe.Pointer(ptr.Data))

	return s1
}

func StringSliceUseBuilder(s string) string {
	var b strings.Builder
	b.Grow(20)
	b.WriteString(s[:20])
	s1 := b.String()

	ptr := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	//字符串字节数组的的地址
	fmt.Println("StringSliceUseBuilder:", unsafe.Pointer(ptr.Data))

	return s1
}
