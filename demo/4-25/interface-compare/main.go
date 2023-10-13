package main

import (
	"fmt"
	"reflect"
)

type ReadErr struct {
	error
}

type Iperson interface {
	say()
}

type person struct {
}

func (p person) say() {

}

func main() {

	var p person
	var inter Iperson
	inter = p
	fmt.Println(p, inter)

	var x int
	var ii interface{}
	ii = x
	fmt.Println(ii)

	err := read()
	fmt.Println(err) //<nil>
	//var e error
	//println(err) //(0xe4298,0x0)
	//println(e)            //(0x0,0x0)
	//fmt.Println(e == err) //false
	if err != nil {
		fmt.Println("has error") //has error
	} else {
		fmt.Println("no error")
	}

	//判断指向底层数据的值是否为nil（实际开发中常用）
	fmt.Println(IsNil(err))

	//var a int
	//var b interface{} = a
	//var c int32
	//var d interface{} = c
	//println("b=", b, "d=", d) //b= (0xe116c0,0xc000043f68) d= (0xe11740,0xc000043f64)
	//println("b==d?", b == d)  //b==d? false

	var e error
	var emptyI interface{}                  // 空接口类型
	println("e = nil:", e == nil)           //e = nil: true
	println("emptyI = nil:", emptyI == nil) //emptyI = nil: true
	println("e :", e, "emptyI", emptyI)     //e : (0x0,0x0) emptyI (0x0,0x0)
	println("e==emptyI", e == emptyI)       //e==emptyI true

	var s []string
	fmt.Println("s==nil", s == nil)
	NilParamToInterface(s)

}

func read() error {
	var err *ReadErr
	return err
}

//type Interface1 interface {
//	f1(string)
//}
//
//type Interface2 interface {
//	f1()
//	f2()
//}
//
//type Interface3 interface {
//	Interface1
//	Interface2
//	f3()
//}

func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	vi := reflect.ValueOf(i)
	return vi.IsNil()
}

func NilParamToInterface(i interface{}) {
	fmt.Println("NilParamToInterface", i == nil)
	fmt.Println("NilParamToInterface IsNil", IsNil(i))
}
