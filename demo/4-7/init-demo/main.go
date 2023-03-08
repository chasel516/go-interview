package main

import (
	. "fmt"
	"init-demo/pkg1"
	"init-demo/pkg2"
)

// 同一个go文件中的多个init方法，按照定义的顺序执行
func init() {
	Println(1)
}

func init() {
	Println(2)
}

func main() {
	/**
	//先按照包的导入顺序执行，pkg先导入所以先执行pkg文件中的init
	//pkg包中又是按照文件名先后顺序执行的a.go>b.go>pkg.go
	pkg-a
	pkg-b
	pkg-pkg
	//按照导入的先后顺序执行pkg1
	pkg1-pkg1
	//pkg2中依赖了pkg3，先执行了pkg3的init
	pkg3-pkg3
	//再执行pkg2本身的init
	pkg2-pkg2
	//执行main包中的init，且按照代码声明的顺序执行
	1
	2

	//最后执行main函数中的打印
	pkg
	pkg1
	pkg2


	*/
	//println(pkg.Pkg)
	println(pkg1.Pkg1)
	println(pkg2.Pkg2)

}
