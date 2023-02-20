package configs

import (
	"fmt"
	"os"
)

func GetWorkPath() {
	wd, err := os.Getwd()
	fmt.Println(wd, err)
}

func GetWorkPathByArg() {
	wd, err := os.Getwd()
	fmt.Println(wd, err)
}
