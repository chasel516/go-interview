package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS(1))
	//物理CPU个数
	fmt.Println(runtime.NumCPU())

	//var wg sync.WaitGroup
	//wg.Add(1)
	//wg.Wait()

}
