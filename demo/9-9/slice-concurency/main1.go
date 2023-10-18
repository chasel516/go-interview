package main

import (
	"fmt"
	"sync"
)

type SafeSlice struct {
	data []interface{}
	lock sync.RWMutex
}

func (s *SafeSlice) Append(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data = append(s.data, item)
}

func (s *SafeSlice) Get(index int) interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if index < 0 || index >= len(s.data) {
		return nil
	}

	return s.data[index]
}

func (s *SafeSlice) len() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.data)
}

func main() {
	s := SafeSlice{}
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			s.Append(n)
		}(i)
	}
	wg.Wait()
	fmt.Println("Length of SafeSlice:", s.len())
}
