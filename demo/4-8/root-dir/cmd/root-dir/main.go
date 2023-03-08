package main

import "root-dir/configs"

func main() {
	//configs.GetWorkPath()
	//configs.GetWorkPathByArg()
	//configs.GetWorkPathByExec()
	//configs.GetWorkPathByCaller()
	configs.GetWorkPathByEnv()
}
