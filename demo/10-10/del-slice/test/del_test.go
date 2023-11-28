package test

import (
	"testing"
)

const N = 10000
const DelIndex = 1000

func BenchmarkDelUseCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, N)
		s = s[:DelIndex+copy(s[DelIndex:], s[DelIndex+1:])]
	}
}

func BenchmarkDelUseAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, N)
		s = append(s[:DelIndex], s[DelIndex+1:]...)
	}
}
