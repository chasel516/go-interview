package main

import "fmt"

func main() {
	//s := []int{1, 2, 3}
	//f(s...)

	//s := []interface{}{1, 2, 3}
	//f(s...)
	f1("imooc")
	//
	//a := [5]int{}
	//const A = f2() //编译不通过
	//const B = len(a)
	//
	//f3()
	//f4()
	//var f func()
	//f()
}

func f(params ...interface{}) {
	if len(params) == 0 {
		fmt.Println("")
	}
	for i, param := range params {
		fmt.Println("i:", i, ";param:", param)
	}
}

func f1(name string, params ...interface{}) {
	if len(params) == 0 {
		fmt.Println(name)
		return
	}
	for i, param := range params {
		fmt.Println("i:", i, ";param:", param)
	}
}

func f2() int {
	return 1
}

func f3() int {
a:
	goto a
}

func f4() bool {
	for {
	}
}
