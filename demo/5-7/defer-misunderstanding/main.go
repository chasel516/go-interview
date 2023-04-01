package main

import "fmt"

type number int

func (n number) print()     { fmt.Println(n) }
func (n *number) ptrprint() { fmt.Println(*n) }

func main() {
	//f1()
	//fmt.Println(f2(2))
	//f3()
	//f4()
	//f5()
	//f6()
	//fn()
	f()
}

func f() {
	var n number

	defer n.print()
	defer n.ptrprint()
	defer func() { n.print() }()
	defer func() { n.ptrprint() }()

	n = 123
}

func f1() {
	x := 0

	defer fmt.Println(x)                    // 作为外部变量传递
	defer func(j int) { fmt.Println(j) }(x) // 作为参数传递
	defer func() { fmt.Println(x) }()       // 作为闭包（closure）进行引用

	x = 123

}

func f2(x int) (r int) {
	defer func() {
		r += x // 修改返回值
	}()

	return x + x // <=> r = x + x; return
}

// defer中变量的估值时刻
func f3() {
	x := 0
	defer fmt.Println(x)
	x = 1
	fmt.Println("done")
}

// defer中实参的估值时刻
func f4() {
	x := 0
	defer func(paramx int) {
		fmt.Println(paramx)
	}(x)
	x = 1
	fmt.Println("done")
}

// defer中闭包的估值时刻
func f5() {
	x := 0
	defer func() {
		fmt.Println(x)
	}()
	x = 1
	fmt.Println("done")
}

func f6() {
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

func fn() int {
	defer func(x int) {
		fmt.Println("defer:", x)
	}(fn1())
	fmt.Println("fn")
	return fn2()
}

func fn1() int {
	fmt.Println("fn1")
	return 1
}

func fn2() int {
	fmt.Println("fn2")
	return 2
}
