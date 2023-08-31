package test

import (
	"strings"
	"testing"
)

var s = strings.Repeat("1", 1<<20)

func BenchmarkStringSlice1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringSlice1(s)
	}
}

func BenchmarkStringSlice2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringSlice2(s)
	}
}

func BenchmarkStringSliceUseBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringSliceUseBuilder(s)
	}
}

func StringSlice1(s string) string {
	s1 := string([]byte(s[:20]))
	return s1
}

func StringSlice2(s string) string {
	s1 := (" " + s[:20])[1:]
	return s1
}

func StringSliceUseBuilder(s string) string {
	var b strings.Builder
	b.Grow(20)
	b.WriteString(s[:20])
	s1 := b.String()
	return s1
}
