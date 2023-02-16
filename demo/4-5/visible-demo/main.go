package main

import (
	"fmt"
	"visible-demo/pkg1"
)

func main() {
	fmt.Println(pkg1.TestConst1)
	fmt.Println(pkg1.TestVer1)
	ts := pkg1.TestStruct1{Field1: "test"}
	ts.Test1()
}
