package main

import "log"

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	m := make(map[string]string)
	m1 := m
	m1["k1"] = "v1"
	log.Printf("m:%v;ptr:%p", m, m)
	log.Printf("m1:%v;ptr:%p", m1, m1)

	s := make([]string, 0, 10)
	s1 := s
	log.Printf("s:%v;ptr:%p", s, s)
	log.Printf("s1:%v;ptr:%p", s1, s1)
	s1 = append(s1, "v1")
	log.Printf("s:%v;ptr:%p", s, s)
	log.Printf("s1:%v;ptr:%p", s1, s1)

}
