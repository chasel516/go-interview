package main

import (
	"math"
	"strconv"
)

func main() {}

func f1() interface{} {
	var str1 string // str1被分配到一个接口中，导致m逃逸到堆上
	return str1
}

func f2() interface{} {
	var str2 interface{} // str1被分配到一个接口中，导致m逃逸到堆上
	return str2
}

func fn() {
	sl := make([]string, 0)
	for i := 0; i < math.MaxInt; i++ {
		sl = append(sl, strconv.Itoa(i))
	}
}
