package main

import (
	"sync"
)

func main() {
	lock := sync.Mutex{}
	lock.Lock()
	//lock.Unlock()
	test(&lock)
}
func test(l *sync.Mutex) {
	l.Lock()
}
