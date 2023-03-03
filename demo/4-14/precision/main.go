package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func main() {
	fmt.Println(strconv.IntSize)
	fmt.Println(math.MaxInt16) //32767
	fmt.Println(math.MinInt16) //-32768
	//如果使用常量，这种情况就不会发生精度损失。
	const c1 = 0.3
	const c2 = 0.6
	fmt.Println("c1+c2=", c1+c2, "c1 type:", reflect.TypeOf(c1))

	//精度的概念：在一个指定范围内，将d位十进制数（按照科学计数法表达）转换为二进制数，再将二进制数转换为d位十进制数，如果数据转换不发生损失，则意味着在此范围内有d位精度
	//Go语言中默认浮点数打印出的值为8位
	var f1 float32 = 0.111111119
	var f2 float32 = 0.111111112
	fmt.Println(f1, f2)

	//浮点可以输出-0
	fmt.Println(math.Float32frombits(0))
	fmt.Println(math.Float32frombits(1 << 31))

	//浮点数不满足结合率

	var x = 0.3
	var y = 0.6
	var z float64 = 0.1
	fmt.Println("(x+y)+z:", (x+y)+z, "x+(y+z):", x+(y+z))

	//四舍但不一定五入
	fmt.Printf("%.2f\n", 9.824)
	fmt.Printf("%.2f\n", 9.8250)
	fmt.Printf("%.2f\n", 9.8251)
}
