package main

import (
	"log"
	"unsafe"
)

type St struct {
	f1 int8
	f2 int16
	f3 int64
}

type St1 struct {
	f1 int8
	f3 int64
	f2 int16
}

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	st := St{}
	st1 := St1{}
	log.Println(unsafe.Sizeof(st))
	log.Println(unsafe.Sizeof(st1))
	log.Println(unsafe.Sizeof(int8(1)))
	log.Println(unsafe.Sizeof(int16(1)))
	log.Println(unsafe.Sizeof(int64(1)))

	log.Println("st.f1:", unsafe.Offsetof(st.f1))
	log.Println("st.f2:", unsafe.Offsetof(st.f2))
	log.Println("st.f3:", unsafe.Offsetof(st.f3))

	log.Println("st1.f1:", unsafe.Offsetof(st1.f1))
	log.Println("st1.f2:", unsafe.Offsetof(st1.f2))
	log.Println("st1.f3:", unsafe.Offsetof(st1.f3))
}
