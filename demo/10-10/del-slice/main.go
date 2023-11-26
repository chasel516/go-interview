package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	//f1()
	//f2()
	f3()
}

func f1() []*int {
	s := []*int{new(int), new(int), new(int), new(int)}
	//索引为0和3的元素在截取后不再使用，只要f1返回后的切片被使用，这两个丢失的元素也不会回收，从而导致内存泄露
	s[0], s[3] = nil, nil
	return s[1:3:3]
}

func f2() {
	s := []int{1, 2, 3, 4, 5}
	//删除索引为1的元素

	//s[i] = s[len(s)-1]
	s[1] = s[4] //将最后一个元素值存到要删除的位置

	//s = s[:len(s)-1]
	s = s[:4] //截取掉最后一个元素
	fmt.Println(s)
}

func f3() {
	s := []int{1, 2, 3, 4, 5}
	// s = s[:i + copy(s[i:], s[i+1:])]
	//假设需要删除索引为i的元素
	//先使用copy将i下标开始的元素使用i+1下标开始的元素进行覆盖，这样下标i的值将被i+1位置的值所取代，相当于将整个切片元素从i开始向后移动了一位
	//i后面的元素被覆盖后再从头开始截取到i的位置即可
	//1 2 3 4 5
	//1 3 4 5 5  执行copy(s[i:], s[i+1:]) 覆盖元素个数为3
	//1 3 4 5  再截取前1+3个元素
	s = s[:1+copy(s[1:], s[2:])]
	fmt.Println(s)
}

func f4() {
	s := []int{1, 2, 3, 4, 5}
	//s = append(s[:i], s[i+1:]...)
	//先截取前i个元素，再追加上i+1开始的元素
	s = append(s[:1], s[2:]...)
}
