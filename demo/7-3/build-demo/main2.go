package main

import (
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// 创建一个文件集
	fset := token.NewFileSet()

	// 创建一个扫描器
	var s scanner.Scanner

	// 打开源文件
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fileInfo, _ := file.Stat()
	// 初始化扫描器
	s.Init(fset.AddFile("main.go", fset.Base(), int(fileInfo.Size())), content, nil, scanner.ScanComments)

	// 遍历每个token
	log.Println("Token sequence:")
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		log.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}

	// 创建一个解析器
	f, err := parser.ParseFile(fset, "main.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// 打印抽象语法树
	log.Println("Abstract syntax tree:")
	ast.Print(fset, f)
}
