package main

import (
	"log"
	"unsafe"
)

type St2 struct {
	f1 int32
	f2 struct{}
}

type St3 struct {
	f2 struct{}
	f1 int32
}

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	log.Println(unsafe.Sizeof(struct{}{}))
	log.Println(unsafe.Sizeof(int32(1)))
	log.Println(unsafe.Sizeof(St2{}))
	log.Println(unsafe.Sizeof(St3{}))
}
