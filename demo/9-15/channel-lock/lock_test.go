package main

import (
	"sync"
	"testing"
)

type Lockers struct {
	mu    sync.Mutex
	muc   chan struct{}
	value int
}

func NewLockers() *Lockers {
	return &Lockers{
		mu:    sync.Mutex{},
		muc:   make(chan struct{}, 1),
		value: 0,
	}
}

func (d *Lockers) MutexWrite() {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.value++
}

func (d *Lockers) MutexChannelWrite() {
	d.muc <- struct{}{}
	d.value++
	<-d.muc
}
func BenchmarkMutexWrite(b *testing.B) {
	l := NewLockers()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.MutexWrite()
	}
}

func BenchmarkMutexChannelWrite(b *testing.B) {
	l := NewLockers()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.MutexChannelWrite()
	}
}
