package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func main() {
	var s1 []struct{}
	s2 := []struct{}{}
	s3 := make([]struct{}, 0)
	s4 := new([]struct{})
	fmt.Println(s1 == nil, s2 == nil, s3 == nil, s4 == nil)
	s5 := make([]struct{}, 2)
	s6 := make([]struct{}, 0, 2)
	fmt.Println(len(s1), len(s5), len(s6))

	return
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
