package main

import (
	"log"
	"sync"
)

func main() {
	m := SafeMap{
		mu: sync.RWMutex{},
		m:  map[any]any{},
	}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(index int) {
			m.Set(index, index)
			wg.Done()
		}(i)
	}
	wg.Wait()
	m.Range(func(key, value any) bool {
		log.Println(key, value)
		return true
	})

}

type SafeMap struct {
	mu sync.RWMutex
	m  map[any]any
}

func (sm *SafeMap) Get(key any) any {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.m[key]
}

func (sm *SafeMap) Set(key, value any) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func (sm *SafeMap) Del(key any) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.m, key)
}

func (sm *SafeMap) Range(f func(key, value any) bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	for k, v := range sm.m {
		if !f(k, v) {
			break
		}
	}
}
