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

	//转字节切片
	s1 := []byte(s)
	ptr = (*reflect.StringHeader)(unsafe.Pointer(&s1))
	fmt.Println("s1:", unsafe.Pointer(ptr.Data))

	//强转
	s2 := string(s1)
	ptr = (*reflect.StringHeader)(unsafe.Pointer(&s2))
	fmt.Println("s2:", unsafe.Pointer(ptr.Data))

	s3 := *(*string)(unsafe.Pointer(&s1))
	ptr = (*reflect.StringHeader)(unsafe.Pointer(&s3))
	fmt.Println("s3:", unsafe.Pointer(ptr.Data))

	//字节切片转字符串
	s4 := string(s1)
	ptr = (*reflect.StringHeader)(unsafe.Pointer(&s4))
	fmt.Println("s4:", unsafe.Pointer(ptr.Data))

	s5 := *(*[]byte)(unsafe.Pointer(&s))
	ptr = (*reflect.StringHeader)(unsafe.Pointer(&s5))
	fmt.Println("s5:", unsafe.Pointer(ptr.Data))

	s6 := stringToBytes(s)
	ptr = (*reflect.StringHeader)(unsafe.Pointer(&s6))
	fmt.Println("s6:", unsafe.Pointer(ptr.Data))

}

func stringToBytes(s string) []byte {
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	b := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: strHeader.Data,
		Len:  strHeader.Len,
		Cap:  strHeader.Len,
	}))
	return b
}
