package main

import "fmt"

const (
	C1 = "msg"
	Cr = "msg"
	C2 = iota
	C3 = iota
)
const (
	C4 = iota
	C5 = iota
)

const (
	A, B = iota + 1, iota + 2 //1,2
	C, D                      //2,3
	E, F                      //3,4
)

const (
	Su = iota
	Mo
	Tu
	We
	Th
	Fr
	Sa
)

const (
	_  = iota             //将0号计数占位，后面从1开始
	KB = 1 << (10 * iota) // <<移位操作，速度比乘除法快
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	//fmt.Printf("C1:%s, C2:%d, C3:%d, C4:%d, C5:%d \n", C1, C2, C3, C4, C5)
	//fmt.Printf("A:%d, B:%d, C:%d, D:%d, E:%d, F:%d \n", A, B, C, D, E, F)
	//fmt.Printf("Su:%d, Mo:%d, Tu:%d, We:%d, Th:%d, Fr:%d, Sa:%d \n", Su, Mo, Tu, We, Th, Fr, Sa)
	fmt.Printf("KB:%d, MB:%d, GB:%d, TB:%d, PB:%d \n", KB, MB, GB, TB, PB)
}
