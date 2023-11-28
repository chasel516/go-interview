package main

import "testing"

type Arr [1024]int

var s1 = make([]Arr, 1000)
var s2 = make([]Arr, 1000)
var s3 = make([]Arr, 1000)

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s1); j++ {
			s1[j][0] += 1
		}
	}
}

func BenchmarkRangeWithIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := range s2 {
			s2[j][0] += 1
		}
	}
}

func BenchmarkRangeWithValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range s3 {
			v[0] += 1
		}
	}
}
