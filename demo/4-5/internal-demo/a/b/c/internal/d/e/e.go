package e

import (
	"fmt"
	//internal内部的包也是可以互相调用的
	"internal-demo/a/b/c/internal/d"
)

var E string

func f() {
	fmt.Println(d.D)
}
