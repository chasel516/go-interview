package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x := 1
	fmt.Println(x)

	y := 1
	_ = y

	s1 := make([]byte, 1, 65*1024) // >64k
	_ = s1

	s2 := make([]byte, 1, 64*1024) // 64k
	_ = s2

	//同理，string占8字节，当其容量大于8*1024时也会发生逃逸
	s3 := make([]string, 1, 8*1024+1)
	_ = s3

	s4 := make([]string, 1, 8*1024)
	_ = s4

	str := ""
	fmt.Println(unsafe.Sizeof(str))
}
