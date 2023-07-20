package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	fmt.Println(debug.SetMaxThreads(1000))
}
