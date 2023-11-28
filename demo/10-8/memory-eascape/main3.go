package main

import (
	"math"
	"strconv"
)

func main() {}

func f5() {
	var arr [1024 * 65]int
	_ = arr
}

func f6() {
	s6 := make([]byte, 1, 65*1024) // >64k
	_ = s6

	s61 := make([]byte, 1, 64*1024) // 64k
	_ = s61
}

func f7() {
	//同理，string占8字节，当其容量大于8*1024时也会发生逃逸
	//golang在1.3版本之后使用了连续栈来取代了分段栈，这样可以提高栈的访问效率，但是也限制了栈的大小。golang在1.4版本中将连续栈的初始大小设为2kb，
	//如果栈的大小超过了64kb，就会触发栈的扩容，这样会影响性能和内存管理。为了避免频繁的栈扩容，golang采用了一种策略，
	//就是当对象的大小超过64kb时，就直接在堆上分配，而不是在栈上。这样可以减少栈的压力，也可以让堆上的对象由GC来回收。
	s7 := make([]string, 1, 8*1024+1)
	_ = s7

	s71 := make([]string, 1, 8*1024)
	_ = s71
}

func f8() {
	s8 := make([]string, 0)
	//变量s8在后续使用过程中会发生扩容，扩容后的对象会申请新的内存空间而被分配到堆上
	for i := 0; i < math.MaxInt; i++ {
		s8 = append(s8, strconv.Itoa(i))
	}
}

func f9() {
	var length int
	s9 := make([]string, 0, length) //切片容量不确定时会发生逃逸
	_ = s9

	m9 := make(map[int]int, length) //map容量不确定时没有发生逃逸
	_ = m9

	c9 := make(chan int, length) //channel是否逃逸跟容量无关
	_ = c9

}
