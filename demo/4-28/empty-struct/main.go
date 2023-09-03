package main

import (
	"empty-struct/set"
	"fmt"
	"unsafe"
)

func main() {
	//st := struct-demo{}{} //unsafe.Sizeof.st: 0
	//fmt.Println("unsafe.Sizeof.st:", unsafe.Sizeof(st))

	set1 := map[int]struct{}{1: {}}
	set2 := map[int]bool{1: false}
	fmt.Println("set1.size:", unsafe.Sizeof(set1[1]), "set2.size:", unsafe.Sizeof(set2[1])) //set1.size: 0 set2.size: 1

	////exists
	if _, ok := set1[1]; ok {
		fmt.Println("exists") //exists
	}
	////add item
	set1[2] = struct{}{}

	////delete item
	delete(set1, 1)
	//fmt.Println("set1", set1) //set1 map[2:{}]
	//
	s := make(set.Set)
	s.Put(1)
	s.Put(2)
	fmt.Println(s.Has(1)) //true
	s.Remove(1)
	fmt.Println(s.Has(1)) //false
}
