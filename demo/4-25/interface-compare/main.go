package main

import (
	"fmt"
)

type ReadErr struct {
	error
}

func main() {
	err := read()
	fmt.Println(err)
	if err != nil {
		fmt.Println("has error")
	} else {
		fmt.Println("no error")
	}

	var a int
	var b interface{} = a
	var c int32
	var d interface{} = c
	println("b==d", b == d)

	var e error
	println("e:", e, "err:", err)

	var emptyI interface{} // 空接口类型
	// 非空接口类型
	println("e = nil:", e == nil)
	println("emptyI = nil:", emptyI == nil)
	println("e :", e, "emptyI", emptyI)
	println("e==emptyI", e == emptyI)
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
