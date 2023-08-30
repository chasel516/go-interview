package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	// 创建一个文件集
	fset := token.NewFileSet()

	// 创建一个解析器
	f, err := parser.ParseFile(fset, "main_test.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// 打印抽象语法树
	log.Println("Abstract syntax tree:")
	ast.Print(fset, f)
}
