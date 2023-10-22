package main

import (
	"sync"
	"testing"
)

type data struct {
	mu    sync.Mutex
	rwmu  sync.RWMutex
	value int
}

func (d *data) mutexWrite() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.value++
}

func (d *data) rwMutexWrite() {
	d.rwmu.Lock()
	defer d.rwmu.Unlock()
	d.value++
}
func BenchmarkMutexWrite(b *testing.B) {
	d := &data{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.mutexWrite()
	}
}

func BenchmarkRWMutexWrite(b *testing.B) {
	d := &data{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.rwMutexWrite()
	}
}
