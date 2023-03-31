package main

import "fmt"

type number int

func (n number) print()     { fmt.Println(n) }
func (n *number) ptrprint() { fmt.Println(*n) }

func main() {
	//f1()
	//f2()
	fmt.Println(f3(2))
}

func f1() {
	var n number

	defer n.print()
	defer n.ptrprint()
	defer func() { n.print() }()
	defer func() { n.ptrprint() }()

	n = 3
}

func f2() {
	x := 0

	defer fmt.Println(x)                    // 作为外部变量传递
	defer func(j int) { fmt.Println(j) }(x) // 作为参数传递
	defer func() { fmt.Println(x) }()       // 作为闭包（closure）进行引用

	x = 123

}

func f3(x int) (r int) {
	defer func() {
		r += x // 修改返回值
	}()

	return x + x // <=> r = x + x; return
}
