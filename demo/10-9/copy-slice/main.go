package main

import (
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	//f1()
	//f2()
	//f3()
	//f4()
	f5()
}

func f1() {
	s1 := []int{3: 100}
	s2 := s1[:4]
	log.Println(s2)
	s1[0] = 1
	log.Println(s2)
	s2[0] = 2
	log.Println(s1[0])

	s := []int{1, 2, 3, 4, 5}
	s3 := s[0:2]
	s3 = append(s3, 0)
	log.Println(s[2], s3[2]) //0 0 由于截取后的切片被追加，原切片的元素被修改

	//使用三元索引解决这种追加导致原切片被修改的问题
	s4 := []int{1, 2, 3, 4, 5}
	s5 := s[0:2:2]
	s5 = append(s5, 0)
	log.Println(s4[2], s5[2]) //3,0 由于截取时限制了共享原切片的结束位置，所以追加时一旦超过原索引的结束位置就会触发截取后切片的扩容，使得截取后的切片不再与原切片共享同一块内存空间
}

func f2() {
	var s1 []int
	s1 = []int{1, 2}
	s2 := []int{3, 4, 5}
	affectItems := copy(s2, s1)
	log.Println(s2, affectItems) //[1 2 5] 2
}

func f3() {
	var s1 []int
	s2 := []int{3, 4, 5}
	affectItems := copy(s2, s1)
	log.Println(s2, affectItems)
}

func f4() {
	s1 := []int{3: 100}
	s2 := []int{3, 4, 5}
	affectItems := copy(s2, s1)
	log.Println(s2, affectItems)
}

func f5() {
	s1 := []int{1, 2, 3}
	s2 := make([]int, len(s1))
	copy(s2, s1)
	log.Println(s2)
}
