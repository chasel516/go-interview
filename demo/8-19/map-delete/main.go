package main

import (
	"fmt"
	"log"
	"runtime"
)

const Cnt = 100000

var m = make(map[int]int, Cnt)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {

	for i := 0; i < Cnt; i++ {
		m[i] = i
	}
	runtime.GC()
	log.Println(getMemStats())
	for k, _ := range m {
		delete(m, k)
	}
	runtime.GC()
	log.Println(getMemStats())
	m = nil
	runtime.GC()
	log.Println(getMemStats())
}

func getMemStats() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return fmt.Sprintf("分配的内存 = %vKB, GC的次数 = %v\n", m.Alloc/1024, m.NumGC)
}
