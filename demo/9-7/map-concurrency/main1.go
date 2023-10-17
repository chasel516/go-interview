package main

import (
	"log"
	"sync"
)

type operation int

const (
	opGet operation = iota
	opSet
	opDelete
	opRange
)

// 将所有对map的操作丢到channel中，利用channel的并发安全来避免map的并发操作
type request struct {
	op    operation
	key   any
	value any
	resp  chan any
}

type ConcurrentMap struct {
	m        map[any]any
	requests chan request
}

func NewConcurrentMap() *ConcurrentMap {
	cm := &ConcurrentMap{
		m:        make(map[any]any),
		requests: make(chan request),
	}
	go cm.run()
	return cm
}

func (cm *ConcurrentMap) run() {
	for req := range cm.requests {
		switch req.op {
		case opGet:
			value, ok := cm.m[req.key]
			if ok {
				req.resp <- value
			} else {
				req.resp <- nil
			}
		case opSet:
			cm.m[req.key] = req.value
			req.resp <- nil
		case opDelete:
			delete(cm.m, req.key)
			req.resp <- nil

		case opRange:
			f := (req.value).(func(key, value any) bool)
			for k, v := range cm.m {
				if !f(k, v) {
					break
				}
			}
			req.resp <- nil
		}
	}
}

func (cm *ConcurrentMap) Get(key any) any {
	resp := make(chan interface{})
	cm.requests <- request{
		op:   opGet,
		key:  key,
		resp: resp,
	}
	return <-resp
}

func (cm *ConcurrentMap) Set(key, value any) {
	resp := make(chan interface{})
	cm.requests <- request{
		op:    opSet,
		key:   key,
		value: value,
		resp:  resp,
	}
	<-resp
}

func (cm *ConcurrentMap) Delete(key any) {
	resp := make(chan interface{})
	cm.requests <- request{
		op:   opDelete,
		key:  key,
		resp: resp,
	}
	<-resp
}

// 使用匿名函数的入参接收迭代后的值
func (cm *ConcurrentMap) Range(value func(key, value any) bool) {
	resp := make(chan interface{})
	cm.requests <- request{
		op:    opRange,
		value: value,
		resp:  resp,
	}
	<-resp
}

func main() {
	cm := NewConcurrentMap()
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(index int) {
			cm.Set(index, index)
			wg.Done()
		}(i)
	}
	wg.Wait()
	cm.Range(func(key, value any) bool {
		log.Println(key, value)
		return true
	})
}
