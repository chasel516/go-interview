package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i *int = nil
	fmt.Println("i.size:", unsafe.Sizeof(i)) //8

	var i8 *int8 = nil
	fmt.Println("i8.size:", unsafe.Sizeof(i8)) //8

	var s *string = nil
	fmt.Println("s.size:", unsafe.Sizeof(s)) //8

	var ps *struct{} = nil
	fmt.Println("ps.size:", unsafe.Sizeof(ps)) //8

	var f func() = nil
	fmt.Println("f.size:", unsafe.Sizeof(f)) //8
	var si []int = nil
	var si1 []int = nil
	fmt.Println("si.size:", unsafe.Sizeof(si)) //24

	var m map[string]int = nil
	fmt.Println("m.size:", unsafe.Sizeof(m)) //8
	var ii interface{} = nil
	//var ii1 interface{} = nil
	fmt.Println("ii.size:", unsafe.Sizeof(ii)) //16

	fmt.Println(i == si)   //编译不通过
	fmt.Println(i == ii)   //fasle
	fmt.Println(si == si1) //编译不通过
}
