package main

import "sync"

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
