package main

import (
	"log"
	"runtime"
	"strings"
)

func main() {
	traceMemStats()
	s1 := creteBigString()
	traceMemStats()
	s2 := s1
	traceMemStats()
	s3 := s2[:]
	s3 = strings.Repeat(s2[:], 1)
	traceMemStats()
	_ = s3

}

func creteBigString() string {
	b := []byte{}
	for i := 0; i < 1024000; i++ {
		b = append(b, '1')
	}
	return string(b)
}

func traceMemStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	log.Printf("Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}
