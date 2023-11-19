package main

import (
	"sync"
	"testing"
)

var sharedResource int
var mu sync.Mutex

func updateSharedResource() {
	for i := 0; i < 1000000; i++ {
		mu.Lock()
		sharedResource++
		mu.Unlock()
	}
}

func BenchmarkUpdateSharedResource(b *testing.B) {
	var wg sync.WaitGroup

	// 启动多个 goroutine，同时访问共享资源
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			updateSharedResource()
		}()
	}

	// 等待所有 goroutine 完成
	wg.Wait()
}
