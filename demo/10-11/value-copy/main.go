package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s []int
	s1 := s
	fmt.Println(unsafe.Sizeof(s), unsafe.Sizeof(s1))
}
