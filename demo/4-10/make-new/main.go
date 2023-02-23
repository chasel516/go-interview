package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var a1 *[]int = new([]int) // allocates slice structure; *p == nil; rarely useful

	var a2 []int = make([]int, 0) // the slice v now refers to a new array of 100 ints
	fmt.Println(*a1 == nil)
	fmt.Println(a2 == nil)
	fmt.Println("a1:", a1, "a2:", a2)
	fmt.Println("a1:", unsafe.Sizeof(a1), "a2:", unsafe.Sizeof(a2))

	//var a1 [100]int
	//var a2 [100]struct{}
	//fmt.Println("a1:", unsafe.Sizeof(a1), "a2:", unsafe.Sizeof(a2))

	//通过make申请
	var s1 = make([]int, 100)

	//初始化第100个元素
	var s2 = []int{99: 0}

	//先创建一个容量为100的数组再截取全部元素作为切片
	var s3 = (&[100]int{})[:]
	var s4 = new([100]int)[:]
	// 100 100 100 100
	println(len(s1), len(s2), len(s3), len(s4))

}

对于使用大量堆栈的应用程序来说，从小的堆栈开始是低效的，因为应用程序必须在每个goroutine堆栈变得可用之前增加它。因此，如果我们知道我们将需要这些空间，那么最好从大的堆栈开始。但是我们要避免在开始时分配太大的堆栈，否则我们会浪费太多的空间，或者花时间不必要地缩小堆栈。

堆栈开始时是最小的要求大小（通常是2KB）。
当堆栈溢出时，我们分配一个两倍大的新堆栈，并将堆栈内容复制到新的分配中。
在GC时间，如果一个堆栈的使用量少于分配的堆栈的1/4，我们就分配一个一半大小的新堆栈，并将堆栈内容复制过去。
