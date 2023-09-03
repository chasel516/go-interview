package main

import (
	"log"
	"unsafe"
)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	arr := [3]int{1, 2, 3}
	log.Printf("arr:%p", &arr)
	f1(arr)
	s := []int{1, 2, 3}
	log.Printf("s:%p", s)
	f2(s)
	s1 := s
	log.Printf("s1 ptr: %p", s1)
	f3(s)
	log.Printf("s : %v; ptr: %p", s, s)
	s2 := s[:1]
	log.Printf("s2 : %v;s2-cap:%d ptr: %p", s2, cap(s2), s2)
	s3 := s[1:2]
	log.Printf("s3 : %v;s3-cap:%d ptr: %p", s3, cap(s3), s3)

	for i, v := range s {
		log.Printf("i:%d,v:%d,ptr:%v,uintptr:%v", i, v, unsafe.Pointer(&s[i]), uintptr(unsafe.Pointer(&s[i])))
	}

}

func f1(a [3]int) {
	log.Printf("f1 arr:%p", &a)
}

func f2(param []int) {
	log.Printf("f2 s:%p", param)
}

func f3(param []int) {
	s3 := []int{4, 5, 6}
	param = s3
	log.Printf("f3 param:%p", param)
}
