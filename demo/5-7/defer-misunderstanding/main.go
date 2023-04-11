package main

import "fmt"

func main() {
	//fmt.Println(f(2))
	//f1()
	//f2()
	//f3()
	//f4()
	//f5()
	//f6()
	//f7()
	f8()
}

type number int

func (n number) print()     { fmt.Println(n) }
func (n *number) ptrprint() { fmt.Println(*n) }
func f8() {
	var n number

	defer n.print()
	defer n.ptrprint()
	defer func() { n.print() }()
	defer func() { n.ptrprint() }()

	n = 123
}

func f(x int) (r int) {
	defer func() {
		r += x // 修改返回值
	}()
	return x + x // <=> r = x + x; return
}

// defer中变量的估值时刻
func f1() {
	x := 0
	defer fmt.Println(x)
	x = 1
	fmt.Println("done")
}

// defer中实参的估值时刻
func f2() {
	x := 0
	defer func(paramx int) {
		fmt.Println(paramx)
	}(x)
	x = 1
	fmt.Println("done")
}

// defer中闭包的估值时刻
func f3() {
	x := 0
	defer func() {
		fmt.Println(x)
	}()
	x = 1
	fmt.Println("done")
}

//func f(x int) (r int) {
//	defer func(param int) {
//		r += param // 修改返回值
//	}(x) // 修改返回值
//	return x + x // <=> r = x + x; return
//}

//func f(x int) (r int) {
//	defer func(param int) {
//		r = param + x // 修改返回值
//	}(r) // 修改返回值
//	return x + x // <=> r = x + x; return
//}

func f4() {
	func() {
		for x := 0; x < 3; x++ {
			defer fmt.Println("x:", x)
		}
	}()

	func() {
		for y := 0; y < 3; y++ {
			defer func() {
				fmt.Println("y:", y)
			}()
		}
	}()
}

func f5() int {
	defer func(x int) {
		fmt.Println("defer:", x)
	}(f6())
	fmt.Println("f5")
	return f7()
}

func f6() int {
	fmt.Println("f6")
	return 1
}

func f7() int {
	fmt.Println("f7")
	return 2
}
