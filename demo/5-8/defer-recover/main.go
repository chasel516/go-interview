package main

import (
	"fmt"
)

func main() {
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("捕获了一个panic：", err)
	//		fmt.Println("防止了程序崩溃")
	//	}
	//}()
	//
	//println("call main")
	//fn()
	//println("exit main")

	//f()
	//paincInCovered()
	//panicInDefer()
	CoveredByCurrentGorutine()
}

func fn() {
	println("call fn")
	fn1()
	println("exit fn")
}
func fn1() {
	println("call fn1")
	defer func() {
		fmt.Println("defer before panic in fn1")
	}()
	panic("panic in fn1")
	defer func() {
		fmt.Println("defer after panic in fn1")
	}()
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
			if err := recover(); err != nil {
				fmt.Println("recover", err)
			}
		}()
	}()
	panic("未知错误") // 演示目的产生的一个panic

}

func paincInCovered() {
	defer func() {
		if err := recover(); err != nil {
			// main panic is override by defer2 panic
			fmt.Println(err)
		} else {
			fmt.Println("defer1 recover nil")
		}
	}()

	defer func() {
		panic("defer2 panic")
	}()

	panic("main panic")
}

func panicInDefer() {

	defer func() {
		fmt.Println("defer1")
		panic("defer1 panic")
	}()

	defer func() {
		fmt.Println("defer2")
		panic("defer2 panic")
	}()

	panic("main panic")

}

func CoveredByCurrentGorutine() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	f()
}
