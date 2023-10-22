package main

import (
	"testing"
)

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
