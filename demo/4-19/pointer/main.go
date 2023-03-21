package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//x := new(int)
	//y := *x
	//y := x
	//fmt.Println("x=", x, "y=", y)
	//tmp := 5
	//*x = tmp
	//fmt.Println("x=", x, *x, "y=", y, *y)

	//var a int
	//ptr := &a
	//pptr := &ptr
	//ppptr := &pptr
	//
	//fmt.Println("a的地址:", &a)
	//fmt.Println("ptr存的地址:", ptr)
	//fmt.Println("pptr存的地址：", pptr)
	//fmt.Println("ppptr存的地址：", ppptr)
	//fmt.Println("ppptr存的地址的指向的指向：", **ppptr)

	//a := new([3]int)
	//fmt.Println(a)
	//b := 3
	//fmt.Println(&b)

	//c := 1
	//f1(c)
	//fmt.Println("f1:c=", c) //c=1
	//f2(&c)
	//fmt.Println("f2:c=", c) //c=10

	//i := int64(5)
	//p := &i
	//p++
	//atomic.AddInt64(p, 1)

	type bl = *bool
	type m = map[int]int

	//var s string  // 一个具名非指针类型
	//var err error // 一个具名接口类型
	//var m1 m      //一个无名类型的别名
	//var b bl      // 一个无名指针类型的别名
	//var iptr *int // 一个无名指针类型

	type Myint int32
	type P1 *int32
	type P2 *Myint

	var x1 P1
	var x2 *int32
	x1 = x2 //隐式转换
	//fmt.Println(x1, x2)
	x2 = x1 //隐式转换
	var x3 *Myint
	//
	//x4 := (*int32)(x3) //显示转换
	//x5 := (*Myint)(x2) //*Myint显示转换到*int32
	//fmt.Println(x4, x5)

	//var x6 P2
	//x6 = x1                         //P1不能直接隐式转换到P2,编译不通过
	//x6 = (*Myint)(x1)               //也不能直接显示转换，编译不通过
	//x6 = P2((*Myint)((*int32)(x1))) //P1经过3层显示转换到P2
	//fmt.Println(x6)
	//
	//fmt.Println(x1 == x2) // 可以比较
	//fmt.Println(x2 == x3) //不能比较
	//fmt.Println(x1 == x6) //不能比较
	//fmt.Println(x2 == x6) //不能比较

	fmt.Println("unsafe.Pointer:", unsafe.Pointer(x2) == unsafe.Pointer(x3)) //可以比较
}

func f1(x int) {
	x = 10
}

func f2(x *int) {
	//实参x的副本跟c指向同一个地址，所以修改副本指向的值也会影响到函数外c的值
	*x = 10
	//这里直接修改了实参x的副本地址，相当于让实参x指向了一个新的地址，所以这里改变x的地址并不会影响函数外c的值
	x = nil
}
