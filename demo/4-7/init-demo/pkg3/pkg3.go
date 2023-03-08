package pkg3

import (
	"fmt"
	"init-demo/pkg2"
)

var Pkg3 = "pkg3"

func init() {
	fmt.Println("pkg3-pkg3")
}

func F() {
	//pkg2中导入了pkg3，pkg3中又导入了pkg2，导致了循环引用，而无法通过编译
	fmt.Println(pkg2.Pkg2)
}
