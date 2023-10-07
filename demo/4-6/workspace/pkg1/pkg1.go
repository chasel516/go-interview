package pkg1

import "github.com/astaxie/beego/logs"

var Pkg1 = "pkg1"

func init() {
	logs.Debug("pkg1")
}
