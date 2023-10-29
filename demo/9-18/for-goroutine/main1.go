package main

import (
	"log"
	"sync"
)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	assignPointerData()
}

type S struct {
	F1 string
	F2 int
	F3 int
}

func assignPointerData() {
	wg := sync.WaitGroup{}
	s1 := &S{
		F1: "test",
		F2: 1,
		F3: 0,
	}
	for i := 0; i < 10; i++ {
		s2 := s1
		s2.F3 = i
		readData(s2, &wg)
	}
	wg.Wait()
}

func readData(s *S, wg *sync.WaitGroup) {
	//s是指针类型，在协程外有修改，内部访问的值不确定
	wg.Add(1)
	go func() {
		log.Println("F3:", s.F3)
		wg.Done()
	}()
}
