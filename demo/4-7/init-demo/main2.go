package main

import (
	"fmt"
	"init-demo/pkg2"
	"init-demo/pkg3"
)

func main() {
	/**
	//pkg3被pkg2导入一次，又被main包导入了一次，最终只执行了被pkg2导入时的init
	pkg3-pkg3
	pkg2-pkg2

	pkg2
	pkg3
	*/
	fmt.Println(pkg2.Pkg2)
	fmt.Println(pkg3.Pkg3)
	//fmt.Println("main.init:", config.GetConfig())
}
