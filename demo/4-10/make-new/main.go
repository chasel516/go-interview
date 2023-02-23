package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a1 [100]int
	var a2 [100]struct{}
	fmt.Println("a1:", unsafe.Sizeof(a1), "a2:", unsafe.Sizeof(a2))
}
