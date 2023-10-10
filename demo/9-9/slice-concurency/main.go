package main

import (
	"log"
	"sync"
)

func main() {
	const Num = 10000
	s := make([]int, 0, Num)
	wg := sync.WaitGroup{}
	for i := 0; i < Num; i++ {
		x := i
		wg.Add(1)
		go func() {
			s = append(s, x)
			wg.Done()
		}()
	}
	wg.Wait()
	log.Println("s.len:", len(s))
}
