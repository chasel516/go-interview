package main

import (
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获了一个panic：", err)
			fmt.Println("防止了程序崩溃")
		}
	}()

	println("call main")
	fn()
	println("exit main")

	//f()
}

func fn() {
	println("call fn")
	fn1()
	println("exit fn")
}
func fn1() {
	println("call fn1")
	panic("panic in fn1")
	fn2()
	println("exit fn1")
}
func fn2() {
	println("call fn2")
	println("exit fn2")
}

func f() {
	go func() {
		defer func() {
			if err := recover(); err == nil {
				fmt.Println("recover", err)
			}
		}()
	}()
	panic("未知错误") // 演示目的产生的一个panic

}
