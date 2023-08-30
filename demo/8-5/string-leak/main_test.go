package main

import (
	"strings"
	"testing"
)

var s string

func init() {
	s = generateString(1 << 20)
}

func BenchmarkAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s2 := s
		_ = s2
	}

}

func BenchmarkAssignPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s2 := &s
		_ = s2
	}

}

func BenchmarkStringSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s2 := s[:]
		_ = s2
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s2 := strings.Repeat(s, 1)
		_ = s2
	}

}

func BenchmarkStringParam(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f1(s)
	}

}

func BenchmarkStringParamPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f2(&s)
	}
}

func BenchmarkStringSlice1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringSlice1(s)
	}
}

func BenchmarkStringSlice2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringSlice2(s)
	}
}

func BenchmarkStringSliceUseBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringSliceUseBuilder(s)
	}
}

func f1(s string) string {
	return s
}

func f2(s *string) *string {
	return s
}

func generateString(length int) string {
	return strings.Repeat("a", length)

}

func stringSlice1(s string) string {
	s1 := string([]byte(s[:20]))
	return s1
}

func stringSlice2(s string) string {
	s1 := (" " + s[:20])[1:]
	return s1
}

func stringSliceUseBuilder(s string) string {
	var b strings.Builder
	b.Grow(20)
	b.WriteString(s[:20])
	s1 := b.String()
	return s1
}
