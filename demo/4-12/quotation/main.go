package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	s1 := "123\t456"
	s2 := `123\t456`
	fmt.Println(s1)
	fmt.Println(s2)

	// 单引号用来定义一个 byte或者rune,默认是rune
	s3 := 'a'
	fmt.Println(s3, reflect.TypeOf(s3), unsafe.Sizeof(s3))

	var s5 byte = 'a'
	fmt.Println(s5, reflect.TypeOf(s5), unsafe.Sizeof(s5))

	s4 := '中'
	fmt.Println(s4, reflect.TypeOf(s4), unsafe.Sizeof(s4))

}
