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
	//os.Args[0] 是当前程序名。如果我们在项目根目录执行程序 bin/cwd，以上程序返回的 binary 结果是 bin/cwd
	fmt.Println("os.Args:", os.Args[0])
	filePath, _ := exec.LookPath(os.Args[0])
	fmt.Println("filePath:", filePath)
	absFilePath, _ := filepath.Abs(filePath)
	fmt.Println("absFilePath:", absFilePath)
	rootDir := path.Dir(absFilePath)
	fmt.Println("rootDir:", rootDir)

}

func GetWorkPathByCaller() {
	_, callPath, _, _ := runtime.Caller(0)
	fmt.Println("callPath:", callPath)
	rootPath := path.Dir(path.Dir(callPath))
	fmt.Println("rootPath:", rootPath)
}

func GetWorkPathByExec() {
	//类似的需求很常见，Go 在 1.8 专门为这样的需求增加了一个函数
	execPath, _ := os.Executable()
	rootDir := filepath.Dir(filepath.Dir(execPath))
	fmt.Println("rootDir:", rootDir)
	rootPath, _ := filepath.EvalSymlinks(rootDir)
	fmt.Println("rootPath:", rootPath)
}
