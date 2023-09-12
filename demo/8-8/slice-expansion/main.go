package main

import (
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	s := []int{1, 2, 3}
	log.Printf("s-ptr:%p; s-cap:%d", s, cap(s))
	s = append(s, 4)
	log.Printf("s-ptr:%p; s-cap:%d", s, cap(s))
	s = append(s, 5)
	log.Printf("s-ptr:%p; s-cap:%d", s, cap(s))
}
