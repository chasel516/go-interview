package main

import (
	"sync/atomic"
)

func main() {
	var v atomic.Value
	var m = map[any]any{1: 1}
	v.Store(m)
	v.Load()
}
