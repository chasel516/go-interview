package pkg

import (
	"testing"
	"time"
)

func BenchmarkF1(b *testing.B) {
	//b.N = 100
	time.Sleep(time.Second)
	//b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f1()
	}
}
func BenchmarkF2(b *testing.B) {
	//b.N = 100
	b.StopTimer()
	time.Sleep(time.Second)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		f2()
	}

}
