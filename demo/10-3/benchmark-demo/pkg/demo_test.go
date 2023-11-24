package pkg

import (
	"testing"
)

func BenchmarkF1(b *testing.B) {
<<<<<<< HEAD
	//b.N = 100
	time.Sleep(time.Second)
	b.ResetTimer()
=======
	b.N = 100
	//time.Sleep(time.Second)
	//b.ResetTimer()
>>>>>>> 1dd392215a15efd147a3c74d8a5568afe6eb8488
	for i := 0; i < b.N; i++ {
		f1()
	}
}
func BenchmarkF2(b *testing.B) {
	b.N = 100
	//b.StopTimer()
	//time.Sleep(time.Second)
	//b.StartTimer()
	for i := 0; i < b.N; i++ {
		f2()
	}

}

func BenchmarkA2(b *testing.B) {
	//b.N = 100
	b.StopTimer()
	time.Sleep(time.Second)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		f2()
	}

}
