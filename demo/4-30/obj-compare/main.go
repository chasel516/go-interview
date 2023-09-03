package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var m map[string]int
	//var n map[string]int
	//
	//fmt.Println(m == nil) //true
	//fmt.Println(n == nil) //true
	//不能通过编译
	//fmt.Println(m ==n)

	//m = make(map[string]int, 10)
	//n = make(map[string]int, 100)
	//for s, i := range m {
	//	fmt.Println(s, i)
	//}
	//m["a"] = 1
	//n["a"] = 1
	//
	//fmt.Println(reflect.DeepEqual(m, n)) //true

	//type person struct-demo {
	//	name string
	//}
	//p1 := person{"test"}
	//p2 := struct-demo {
	//	name string
	//}{"test"}
	//fmt.Println(p1 == p2)                  //true
	//fmt.Println(reflect.DeepEqual(p1, p2)) //false
	//
	//x := new(int)
	//y := new(int)
	//fmt.Println("x==y:", x == y)                             //x==y: false
	//fmt.Println("DeepEqual(x, y):", reflect.DeepEqual(x, y)) //DeepEqual(x, y): true

	type link struct {
		data interface{}
		next *link
	}
	var a, b, c link
	a.next = &b
	b.next = &c
	c.next = &a

	fmt.Println(a == b)                  //false
	fmt.Println(reflect.DeepEqual(a, b)) //true
}
