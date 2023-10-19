package main

import (
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	m := make(map[int]int, 6)
	m[0] = 0
	m[1] = 1
	m[2] = 2
	log.Printf("s:%p", m)
	//assignMap(m)
	//log.Println("s:%v", m)
	chageMapItem(m)
	log.Println("m:", m)
	appendMapItem(m)
	log.Println("m:", m)

}

func assignMap(param map[int]int) {
	log.Printf("assignMap param:%p", param)
	m1 := make(map[int]int, 6)
	//param与函数外的切片s解除引用关系，同时param将指向s1的数据地址
	param = m1
	log.Printf("assignMap param:%p", param)
	param[0] = 4
}

// 会影响到函数外切片的值
func chageMapItem(param map[int]int) {
	param[1] = 5
}

func appendMapItem(param map[int]int) {
	log.Printf("appendSliceItem param addr:%p len:%d", param, len(param))

	//发生扩容，指向的数据地址发生了变化（发生了copy），所以不会影响到函数的切片
	param[3] = 6
	log.Printf("appendSliceItem param addr:%p len:%d", param, len(param))
}
