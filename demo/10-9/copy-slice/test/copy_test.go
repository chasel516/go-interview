package test

import (
	"testing"
)

const N = 10000

var s1ForCopy = make([]int, N)
var s1ForAppend = make([]int, N)
var s2ForCopy []int
var s2ForAppend []int

func BenchmarkCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s2ForCopy = make([]int, N)
		copy(s2ForCopy, s1ForCopy)
	}
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s2ForAppend = append(s1ForAppend[:0:0], s1ForAppend...)
	}
}
