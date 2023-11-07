package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	lock := sync.Mutex{}
	lock.Lock()
	test1(&lock)

}
func test1(l *sync.Mutex) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		test2(l)
		wg.Done()
	}()
	wg.Wait()
}
func test2(l *sync.Mutex) {
	defer l.Unlock()
	r := rand.Intn(10)
	if r%2 == 0 {
		return
	}
	doSth(l)
}

func doSth(l *sync.Mutex) {
	l.Lock()
	fmt.Println("doSth")
}
