package main

import "github.com/astaxie/beego/logs"
import "gitee.com/phper95/pkg/errors"

func main() {
	logs.Debug(123)
	errors.New("err")
}
