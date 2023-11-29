package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	fmt.Println(f1())
	f2()
	//f3()
}

func f1() []*int {
	s := []*int{new(int), new(int), new(int), new(int)}
	//索引为0和3的元素在截取后不再使用，只要f1返回后的切片被使用，这两个丢失的元素也不会回收，从而导致内存泄露
	s[0], s[3] = nil, nil
	return s[1:3:3]
}

// 1, 2, 3, 4, 5
// 1, 5, 3, 4, 5
// 1, 5, 3, 4,
func f2() {
	s := []int{1, 2, 3, 4, 5}
	//删除索引为1的元素

	//s[i] = s[len(s)-1]
	s[1] = s[4] //将最后一个元素值存到要删除的位置

	//s = s[:len(s)-1]
	s = s[:4] //截取掉最后一个元素
	fmt.Println(s)
}

// 1, 2, 3, 4, 5
//
//	3, 4, 5, 5
//	3, 4, 5
//
// 1, 3, 4, 5, 5
// 1, 3, 4, 5
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

// 最小代价实现，不关心切片元素的顺序
func f5(s []int, start, end int) []int {
	//如果 len(s)-end 小于 n，说明从 end 到切片末尾的元素数量比从 start 到 end 的数量多。
	//在这种情况下，可以直接用 copy 将剩余元素往前移动，覆盖要删除的元素。
	//否则从切片尾部开始取出跟删除元素个数相同的切片覆盖掉要删除的start到end位置，再截取掉后面的元素
	// 1 2 3 4 5 6 7 8 假设start=1，end=3
	// n=3-1=2      8-3>2 copy(s[start:end], s[len(s)-n:])
	//s[len(s)-n:]  7 8
	//s[start:end] 2 3
	// 1 7 8  4 5 6 7 8
	//s[:len(s)-(end-start)] 1 7 8  4 5 6
	if n := end - start; len(s)-end < n {
		copy(s[start:end], s[end:])
	} else {
		copy(s[start:end], s[len(s)-n:])
	}
	//截取掉切片后面的元素
	return s[:len(s)-(end-start)]
}

func f6(s []int, start, end int) []int {
	//copy(s[start:], s[end:])表示用end到之后的元素覆盖从start开始的元素，再截取start+覆盖的个数
	// 1 2 3 4 5 6 7 8 假设start=1，end=3
	//s[start:]  2 3 4 5 6 7 8
	//s[end:]   4 5 6 7 8
	//copy(s[start:], s[end:])  4 5 6 7 8 7 8 覆盖了5个元素
	//覆盖后的s=[1 4 5 6 7 8 7 8]

	//  s[:start+copy(s[start:], s[end:])] = s[:1+5]=[1 4 5 6 7 8]
	return s[:start+copy(s[start:], s[end:])]
}

func f7(s []int, start, end int) []int {
	return append(s[:start], s[end:]...)
}
