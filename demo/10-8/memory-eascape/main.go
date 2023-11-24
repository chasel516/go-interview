package main

import (
	"math"
	"strconv"
)

var g1 = 10
var g2 *int
var g3 = &g1
var g4 = "global"
var g5 = []int{1, 2, 3}
var g6 = map[int]int{1: 1} //发生了逃逸
var g7 map[int]int         //只定义则不会发生逃逸
var g8 = make(chan int)
var g9 = [3]int{}

type st struct {
	Name string
	Tag  map[int]int
}

var g10 = st{
	Name: "test",
	Tag:  nil,
}

func main() {
	_ = g3
	_ = g6
}

func f1() (int, *int, string, []int, map[int]int, chan int, [3]int, st) {
	g8 <- g1
	go func() { //func逃逸了，但变量g8没有
		<-g8
	}()
	return g1, g2, g4, g5, g7, g8, g9, g10
}

/*
	func f2() {
		n1 := 1
		fmt.Println(n1)
		n2 := 1
		_ = n2
	}

	func f3() {
		var i interface{}
		var n3 int
		i = n3
		_ = i

}

	func f4() *int {
		var n4 int
		n4 = 42
		return &n4 // x的地址被返回，导致x逃逸到堆上
	}

	func f5() {
		var x [1024]int // x的大小超过栈的容量，导致x逃逸到堆上
		x[0] = 42
	}

	func f6() {
		s1 := make([]byte, 1, 65*1024) // >64k
		_ = s1

		s2 := make([]byte, 1, 64*1024) // 64k
		_ = s2
	}

	func f7() {
		//同理，string占8字节，当其容量大于8*1024时也会发生逃逸
		//golang在1.3版本之后使用了连续栈来取代了分段栈，这样可以提高栈的访问效率，但是也限制了栈的大小。golang在1.4版本中将连续栈的初始大小设为2kb，
		//如果栈的大小超过了64kb，就会触发栈的扩容，这样会影响性能和内存管理。为了避免频繁的栈扩容，golang采用了一种策略，
		//就是当对象的大小超过64kb时，就直接在堆上分配，而不是在栈上。这样可以减少栈的压力，也可以让堆上的对象由GC来回收。
		s3 := make([]string, 1, 8*1024+1)
		_ = s3

		s4 := make([]string, 1, 8*1024)
		_ = s4
	}
*/
func f8() {
	s26 := make([]string, 0)
	//变量在后续使用过程中会发生扩容，扩容后可能会被分配到堆上，但这种情况逃逸分析是没法分析到的
	for i := 0; i < math.MaxInt; i++ {
		s26 = append(s26, strconv.Itoa(i))
	}
}

/*
func f9() func() int {

	z := 7
	return func() int {
		return z // 函数返回后z还在使用，导致z逃逸到堆上
	}
}

func f10() {
	k := 8
	//由于k被goroutine使用，当函数f10返回后，goroutine依然可能在执行，所以goroutine的闭包函数和其中的变量都会发生逃逸
	go func() { //go启动的闭包和使用的变量会发生逃逸
		k += 1
	}()
}

func f11() {

	func() {
		j := 10
		j += 1
	}()
}

func f12() interface{} {
	var str1 string // str1被分配到一个接口中，导致m逃逸到堆上
	return str1
}

func f13() string {
	var str2 string
	return str2
}

func f14() {
	s5 := []interface{}{1, 2}
	s5[0] = 0 //对接口类型的切片元素赋值，切片元素会发生逃逸
}

func f15() {
	m3 := make([]*int, 2)
	i := 1
	m3[0] = &i //切片元素是指针类型时，对切片元素赋值会发生逃逸
}

func f16() {
	s6 := []int{1, 2}
	s6[0] = 0
}

func f17() {
	m1 := make(map[int]interface{}, 2)
	m1[0] = 1 //对接口类型的map元素赋值，元素会发生逃逸
}

func f18() {
	m2 := make(map[int]int, 2)
	m2[0] = 1
}

func f19() {
	m3 := make(map[int][]int, 2)
	m3[0] = []int{1} //map元素为切片时，元素赋值时会发生逃逸
}

// 指针传参
func f21() {
	n20 := 10
	fn1(&n20) //不会发生逃逸
}

func f22() {
	s22 := []int{1}
	fn2(s22) //不会发生逃逸
}

func f23() {
	s23 := []*int{g3}
	fn3(s23) //不会发生逃逸
}

func f24() {
	reflect.TypeOf(g1)
	reflect.ValueOf(g1) //反射获取值的时候会发生逃逸

}


func f25() {
	fn(g2)
	n25 := &g5
	fn(n25)
}

func f26() {
	gr26 := fn(g2)
	_ = gr26
	n26 := &g5
	r26 := fn(n26)
	_ = r26.([]int)
}
*/

func f27() {
	i27 := new(int)
	_ = i27

	str27 := new(string)
	_ = str27

	s27 := new([]int)
	_ = s27

	m27 := new([]map[int]int)
	_ = m27

	st27 := new(st)
	_ = st27

	c27 := new(chan int)
	_ = c27
}

func f28() {
	st28 := new(st)
	st28.Name = "test"
}

// 指针类型的结构体作为返回值导致对象逃逸
func f29() *st {
	st29 := new(st)
	return st29
}

func f30() st {
	st30 := new(st)
	return *st30
}

func f31() *int {
	i31 := new(int)
	return i31
}

func fn1(param1 *int) {
	_ = param1
}

func fn2(param2 []int) {
	_ = param2
}

func fn3(param3 []*int) {
	_ = param3
}

func fn(param2 any) any {
	return param2
}
