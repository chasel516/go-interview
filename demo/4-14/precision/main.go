package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 1.0
	fmt.Println(reflect.TypeOf(a))

	var b int = 'x'
	fmt.Println(b, reflect.TypeOf(b))

	c := 1.0
	fmt.Println(reflect.TypeOf(c))
}
