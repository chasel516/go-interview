package main

import (
	"fmt"
	"log"
	"math"
	"reflect"
)

func main() {
	var a int = 1.0
	fmt.Println(reflect.TypeOf(a))

	b := 1.0
	fmt.Println(reflect.TypeOf(b))

	var c int = 'x'
	fmt.Println(c, reflect.TypeOf(c))

	var x = 0.3
	var y = 0.6
	fmt.Println(x + y)

	var x1 uint64 = math.MaxUint64
	var y1 uint64 = 1

	//整型溢出，导致产生无符号整数反转现象
	addUint(x1, y1)

	//有符号整数运算时出现溢出
	var x2 int32 = math.MaxInt32
	var y2 int32 = 1
	addInt(x2, y2)

	var x3 int32 = math.MaxInt32
	y3 := int16(x3)
	fmt.Println("y3:", y3)
	if (x3 < math.MinInt16) || (x3 > math.MaxInt16) {
		// 错误处理
	}

	//整型转换时出现符号丢失
	var x4 int32 = math.MinInt32
	y4 := uint32(x4)
	fmt.Println("x4:", x4, "y4:", y4)
	if x4 < 0 {
		// 错误处理
	}

	//移位操作的位数不够
	shift(65535, 16)

}

func addUint(x, y uint64) {
	/**
	 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111
	 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0001
	10000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000
	*/

	sum := x + y
	fmt.Println("addUint:", sum)
	fmt.Printf("addUint %b \n", sum)
}

func safeAddUint(x, y uint64) {
	if math.MaxUint64-x < y {
		// 错误处理
		log.Fatal("overflow MaxUint64")
	} else {
		sum := x + y
		fmt.Println(sum)
	}
}

func addInt(x, y int32) {
	sum := x + y
	fmt.Println("addInt: ", sum)
}
func safeAddInt(x, y int32) {
	if (x > 0 && y > (math.MaxInt32-x)) || (y < 0 && x < (math.MinInt32-y)) || (y > 0 && x > (math.MaxInt32-y)) ||
		(x < 0 && y < (math.MinInt32-x)) {
		// 错误处理
		log.Fatal("overflow int32")
	} else {
		sum := x + y
		fmt.Println(sum)
	}
}

func shift(x uint16, bits uint8) {
	if x > (1 << bits) {
		fmt.Println("shift ok")
	} else {
		fmt.Println("shift not ok")
	}
}

func safeShift(x uint16, bits uint8) {
	//uint32(1<<bits)
	if uint32(x) > (uint32(1) << bits) {
		fmt.Println("shift ok")
	} else {
		fmt.Println("shift not ok")
	}
}
