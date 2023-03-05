package configs

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
)

func GetWorkPath() {
	wd, err := os.Getwd()
	fmt.Println("wd", wd, err)
}

func GetWorkPathByArg() {
	//os.Args[0] 是当前程序名。如果我们在项目根目录执行程序 bin/cwd，以上程序返回的 binary 结果是bin/cwd
	fmt.Println("os.Args:", os.Args[0])
	filePath, _ := exec.LookPath(os.Args[0])
	fmt.Println("filePath:", filePath)
	absFilePath, _ := filepath.Abs(filePath)
	fmt.Println("absFilePath:", absFilePath)
	rootDir, _ := filepath.Abs(path.Dir(absFilePath))
	fmt.Println("rootDir:", rootDir)
}

func GetWorkPathByCaller() {
	_, callPath, _, _ := runtime.Caller(0)
	fmt.Println("callPath:", callPath)
	rootPath := path.Dir(path.Dir(callPath))
	fmt.Println("rootPath:", rootPath)
}

func GetWorkPathByExec() {
	//Go1.8 增加的函数
	execPath, _ := os.Executable()
	fmt.Println("execPath:", execPath)
	rootDir := filepath.Dir(execPath)
	fmt.Println("rootDir:", rootDir)
	rootPath, _ := filepath.EvalSymlinks(rootDir)
	fmt.Println("rootPath:", rootPath)
}

func GetWorkPathByEnv() {
	os.Getenv("GOPATH")
	os.Getenv("APPPATH")
}
