package main

import (
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	s := []int{1, 2, 3}
	//s := make([]int, 0, 6)
	//s = append(s, 1, 2, 3)
	//log.Printf("s:%p", s)
	//assignSlice(s)
	//chageSliceItem(s)
	//log.Println("s:", s)
	appendSliceItem(s)
	log.Println("s:", s)

	//通过反射改变切片的长度
	//reflect.ValueOf(&s).Elem().SetLen(4)
	//log.Println("s:", s)

}

func assignSlice(param []int) {
	log.Printf("assignSlice param:%p", param)
	s1 := []int{1, 2, 3}
	//param与函数外的切片s解除引用关系，同时param将指向s1的数据地址
	param = s1
	log.Printf("assignSlice param:%p", param)
	param[0] = 5
}

// 会影响到函数外切片的值
func chageSliceItem(param []int) {
	param[0] = 4
}

func appendSliceItem(param []int) {
	log.Printf("appendSliceItem param addr:%p cap:%d", param, cap(param))

	//发生扩容，指向的数据地址发生了变化（发生了copy），所以不会影响到函数的切片
	param = append(param, 5)
	param[0] = 6
	log.Printf("appendSliceItem param addr:%p cap:%d", param, cap(param))
}
