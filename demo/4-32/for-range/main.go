package main

import "fmt"

type St struct {
	name string
}

func (s *St) Point() {
	fmt.Println(s.name)
}

//func (s St) Point() {
//	fmt.Println(s.name)
//}

func main() {

	st := []St{
		{"a"},
		{"b"},
		{"c"},
	}
	for _, t := range st {
		//fmt.Println(reflect.TypeOf(t))
		defer t.Point()
	}
}
