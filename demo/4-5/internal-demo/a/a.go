package a

import (
	"fmt"
	//internal父级的父级的父级目录下的包不可以访问internal目录内部的包
	"internal-demo/a/b/c/internal/d"
	"internal-demo/a/b/c/internal/d/e"
)

func f() {
	fmt.Println(d.D)
	fmt.Println(e.E)
}
