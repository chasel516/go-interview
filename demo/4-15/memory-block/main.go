package main

func main() {
	//debug.SetMaxStack(1 << 16)
	var x int
	println(&x)
	getArrValue(1)
	println(&x)
}

// getArrValue下面go:noinline这行注释防止函数发生内联

//go:noinline
func getArrValue(i int) int {
	var a [1 << 20]int // 使栈增长
	return a[i]
}
