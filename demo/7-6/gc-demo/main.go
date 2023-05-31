package main

import (
	"log"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)

	s := make([]int, 10000000)
	for i := range s {
		s[i] = i
	}

	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)

	s = nil
	//设置为nil后并不会立马GC
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)

	runtime.GC()

	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)
}
