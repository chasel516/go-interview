package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type persons struct {
	names []string
}

func main() {
	//s1 := make([]int, 4, 10)
	//for i, _ := range s1 {
	//	s1[i] = i
	//}
	//s2 := s1
	//fmt.Println("s1=", s1, "s2=", s2)                   //	s1= [0 1 2 3] s2= [0 1 2 3]
	//fmt.Println("s1.cap=", cap(s1), "s2.cap=", cap(s2)) //s1.cap= 10 s2.cap= 10
	//reflect.ValueOf(&s1).Elem().SetCap(5)
	//fmt.Println("s1.cap=", cap(s1), "s2.cap=", cap(s2)) //	s1.cap= 5 s2.cap= 10
	//fmt.Printf("s1.pointer=%p;s2.pointer=%p\n", s1, s2) //	s1.pointer=0xc00001a280;s2.pointer=0xc00001a280
	//s := []string{"a", "b", "c"}
	//fmt.Println("s1:origin", s)
	//fmt.Printf("s.pointer:%p\n", s)
	//changes1(s)
	//fmt.Println("s:f1", s)

	//changes2(s)
	//fmt.Println("s:f2", s)
	////
	//changes3(s)
	//fmt.Println("s:f3", s)

	//str := "abc"
	//fmt.Printf("str.pointer=%p\n", &str) //str.pointer=0xc000048050
	//changeStr(str)
	//fmt.Println("str:", str) //str: abc
	//
	//m := map[string]int{"a": 1}
	//fmt.Printf("m.pointer=%p\n", m) //m.pointer=0xc0000261b0
	//changeMap(m)
	//fmt.Println("m:", m) //m: map[a:2]

	//var c = make(chan int)
	//fmt.Printf("c.pointer=%p\n", c) //c.pointer=0xc000022180
	//go func() {
	//	c <- 1
	//	addChannel(c)
	//	close(c)
	//}()
	//
	//for item := range c {
	//	//item: 1
	//	//item: 2
	//	fmt.Println("item:", item)
	//}

	//var f func() = func() {}
	//fmt.Printf("f.pointer=%p\n", f) //f.pointer=0x142af60
	//changeFun(f)

	p1 := person{
		name: "go",
		age:  1,
	}
	changeStruct1(p1)
	fmt.Println("p1:=", p1) //p1:= {go 1}
	//
	s := []string{"a", "b", "c"}
	p2 := persons{names: s}
	fmt.Println("p2:=", p2)                        //p2:= {[a b c]}
	fmt.Printf("p2.names.pointer:%p \n", p2.names) //p2.names.pointer:0xc0000261b0
	//
	changeStruct2(p2)
	fmt.Println("p2:=", p2) //p2:= {[a b c]}
	//
	//changeStruct3(p2)
	//p3 := p2
	//fmt.Printf("p2.names.pointer:%p;p3.names.pointer:%p \n", p2.names, p3.names) //p2.names.pointer:0xc0000261b0;p3.names.pointer:0xc0000261b0

}

func changes1(s []string) {
	fmt.Printf("s.pointer:%p\n", s)
	var tmp = []string{"x", "y", "z"}
	s = tmp
}

func changes2(s []string) {
	// item只是一个副本，不能改变s中元素的值
	for i, item := range s {
		item = "d"
		fmt.Printf("item=%s;s[%d]=%s", item, i, s[i])
	}
}

func changes3(s []string) {
	for i := range s {
		s[i] = "d"
	}
}

func changeStr(str string) {
	fmt.Printf("str.pointer=%p\n", &str) //str.pointer=0xc000048060
	//for i, _ := range str {
	//	str[i] = 'a'
	//}
}

func changeMap(mp map[string]int) {
	mp["a"] = 2
	fmt.Printf("mp.pointer=%p\n", mp) //mp.pointer=0xc0000261b0
}

func addChannel(done chan int) {
	done <- 2
	fmt.Printf("done.pointer=%p\n", done) //done.pointer=0xc000022180
}

func changeFun(fun func()) {
	fmt.Printf("fun.pointer=%p\n", fun) //fun.pointer=0x142af60
}

func changeStruct1(p person) {
	p.age = 2
	p.name = "面试"
}

func changeStruct2(ps persons) {
	fmt.Printf("ps.names.pointer:%p \n", ps.names) //ps.names.pointer:0xc0000261b0
	ps.names[2] = "d"
}

func changeStruct3(ps persons) {
	fmt.Printf("ps.names.pointer1:%p \n", ps.names) //ps.names.pointer1:0xc0000261b0

	//赋值后指向了不同的间接值部
	ps.names = []string{"e", "f", "g"}
	fmt.Printf("ps.names.pointer2:%p \n", ps.names) //ps.names.pointer2:0xc0000261e0
}
