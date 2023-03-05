package main

import "root-dir/configs"

func main() {
	//configs.GetWorkPath()
	//configs.GetWorkPathByArg()
	//configs.GetWorkPathByCaller()
	//configs.GetWorkPathByExec()
	configs.GetWorkPathByEnv()
}
