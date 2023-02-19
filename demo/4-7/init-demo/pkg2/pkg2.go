package pkg2

import "fmt"
import "init-demo/pkg3"

var Pkg2 = "pkg2"

func init() {
	fmt.Println("pkg2-pkg2")
}

func F() {
	Pkg2 = pkg3.Pkg3
}
