package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	//var a int = 1.0
	//fmt.Println(reflect.TypeOf(a))
	//
	//b := 1.0
	//fmt.Println(reflect.TypeOf(b))
	//
	//var c int = 'x'
	//fmt.Println(c, reflect.TypeOf(c))

	//var x = 0.3
	//var y = 0.6
	//fmt.Println(x + y)
	//
	//var x1 uint64 = math.MaxUint64
	//var y1 uint64 = 1
	////无符号整型溢出
	//addUint(x1, y1)
	//
	////有符号整数运算时出现溢出
	//var x2 int64 = math.MaxInt64
	//var y2 int64 = 1
	//addInt(x2, y2)
	//
	//x3 := math.MinInt64
	//y3 := 1
	//fmt.Println("math.MinInt64", math.MinInt64, " x3-y3:", x3-y3)

	//对于常量值，不允许溢出，在运行时阶段就会被阻止
	//fmt.Println(math.MinInt64 - 1)

	////数值类型转换时需要考虑溢出问题
	//var x4 int32 = math.MaxInt32
	//y4 := int16(x4)
	//fmt.Println("y4:", y4)
	//if (x4 < math.MinInt16) || (x4 > math.MaxInt16) {
	//	// 错误处理
	//}

	//整型转换时出现符号丢失
	//var x5 int32 = math.MinInt32
	//y5 := uint32(x5)
	//fmt.Println("x5:", x5, "y5:", y5)
	//if x5 < 0 {
	//	// 错误处理
	//}

	//移位操作的位数不够
	//shift(65535, 16)
	safeShift(65535, 16)

	//浮点型溢出
	//x := math.MaxFloat64
	//y := float64(1000)
	//z := x + y
	//fmt.Println("z:", z, "math.MaxFloat64:", x, x == z)

	//浮点数不能被移位
	//var n uint = 32
	//var _ = 1.23 << n
}

func addUint(x, y uint64) {
	/**
	 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111
	 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0001
	00000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000
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

func addInt(x, y int64) {
	sum := x + y
	fmt.Println("addInt64: ", sum, "max int64", math.MaxInt64)
}
func safeAddInt(x, y int64) {
	if (x > 0 && y > (math.MaxInt64-x)) || (y < 0 && x < (math.MinInt64-y)) || (y > 0 && x > (math.MaxInt64-y)) ||
		(x < 0 && y < (math.MinInt64-x)) {
		// 错误处理
		log.Fatal("overflow int64")
	} else {
		sum := x + y
		fmt.Println(sum)
	}
}

func shift(x uint16, bits uint8) {
	fmt.Println("MaxUint16:", math.MaxUint16)
	fmt.Println("x:", x, "1 << bits:", 1<<bits, " uint16(1<<bits):", uint16(1<<bits))
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
