package pkg

import "testing"

func BenchmarkF1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f1()
	}
}
func BenchmarkF2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f2()
	}
}
