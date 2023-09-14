package main

import "fmt"

func main() {
	var m map[string]int // 这是一个nil map
	//m1 := map[string]int{}
	fmt.Println(m)        // 输出: map[]
	fmt.Println(m["key"]) //0
	delete(m, "key")
	//m["key"] = 1 // 运行时错误: panic: assignment to entry in nil map
}
