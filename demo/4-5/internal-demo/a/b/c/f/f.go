package f

import (
	"fmt"
	"internal-demo/a/b/c/internal/d"
	"internal-demo/a/b/c/internal/d/e"
)

func f() {
	//internal同级目录下的包可以访问internal目录内部的包
	fmt.Println(d.D)
	fmt.Println(e.E)
}
