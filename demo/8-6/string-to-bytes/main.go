package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main() {
	s := strings.Repeat("1", 1<<20)
	ptr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println("s pointer:", unsafe.Pointer(ptr.Data))
	StringSlice(s)
}
func StringSlice(s string) {
	s1 := []byte(s[:20])
	ptr := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	//字符串字节数组的的地址
	fmt.Println("s1:", unsafe.Pointer(ptr.Data))
	s2 := string(s1)
	ptr = (*reflect.StringHeader)(unsafe.Pointer(&s2))
	//字符串字节数组的的地址
	fmt.Println("s2:", unsafe.Pointer(ptr.Data))

	s3 := *(*string)(unsafe.Pointer(&s1))
	ptr = (*reflect.StringHeader)(unsafe.Pointer(&s3))
	//字符串字节数组的的地址
	fmt.Println("s3:", unsafe.Pointer(ptr.Data))
}
