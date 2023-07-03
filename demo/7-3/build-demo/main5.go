package main

func main() {
	n := 1
	for i := 0; i < 10; i++ {
		n = incr(n)
	}
	println(n)
}

// //go:noinline
func incr(n int) int {
	return n + 1
}
