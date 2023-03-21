package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func main() {
	var b1 []byte
	b2 := []byte{}
	//fmt.Println(reflect.DeepEqual(b1, b2)) //false
	//
	//v1 := reflect.ValueOf(b1)
	//v2 := reflect.ValueOf(b2)
	//fmt.Println(v1.IsNil(), v2.IsNil())
	//
	fmt.Println(bytes.Equal(b1, b2)) //true
	//
	var b3 []byte
	var b4 []byte
	fmt.Println(reflect.DeepEqual(b3, b4))
	fmt.Println(bytes.Equal(b3, b4)) //true
	//
	b5 := []byte{}                         //true
	b6 := []byte{}                         //true
	fmt.Println(reflect.DeepEqual(b5, b6)) //true
	fmt.Println(bytes.Equal(b5, b6))       //true
}
