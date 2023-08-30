package main

import (
	"fmt"
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
	file, err := os.Open("main_test.go")
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
	s.Init(fset.AddFile("main_test.go", fset.Base(), int(fileInfo.Size())), content, nil, scanner.ScanComments)

	// 遍历每个token
	log.Println("Token sequence:")
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}
