package main

import (
	"log"
)

func main() {
	s := []int{255: 1}
	log.Printf("s-ptr:%p; s-cap:%d", s, cap(s))
	s = append(s, 1)
	log.Printf("s-ptr:%p; s-cap:%d", s, cap(s))
	s = append(s, []int{256: 1}...)
	log.Println("newcap:=", 512+(512+(3*256))/4)
	log.Printf("s-ptr:%p; s-cap:%d", s, cap(s))
}
