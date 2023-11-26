package main

import (
	"fmt"
	"unsafe"
)

func main() {
	st := struct{}{}
	fmt.Println(unsafe.Sizeof(st))
	arr := [3]struct{}{{}, {}, {}}
	fmt.Println(unsafe.Sizeof(arr))
}
