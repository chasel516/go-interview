package main

import (
	"fmt"
)

type S struct {
	Name *string
	Id   int
}

func (s *S) String() string {
	return fmt.Sprintf("${Id:%v,Name:%v}", s.Id, *s.Name)
}

type M map[string]*S

func main() {
	name1 := "name1"
	name2 := "name2"
	s1 := S{
		Name: &name1,
		Id:   1,
	}

	s2 := &S{
		Name: &name2,
		Id:   2,
	}
	m := make(map[string]*S)
	m["m1"] = &s1
	m["m2"] = s2

	//S2声明的时候使用了地址，所以会调用String方法
	fmt.Printf("s1 : %v ; s2 : %v ; m: %v \n", s1, s2, m)
	fmt.Printf("s1 : %+v ;s2 : %v ; m: %+v \n", s1, s2, m)
	fmt.Printf("s1 : %#v ;s2 : %v ; m: %#v \n", s1, s2, m)

	//输出百分数且保留2位小数 (%在)
	fmt.Printf("%.2f%% \n", 99.99)

}
