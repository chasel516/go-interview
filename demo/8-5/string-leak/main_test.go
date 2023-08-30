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

func f1(s string) string {
	return s
}

func f2(s *string) *string {
	return s
}

func generateString(length int) string {
	return strings.Repeat("a", length)

}
