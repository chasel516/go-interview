package configs

import (
	"fmt"
	"path/filepath"
)

func GetPath() {
	path, err := filepath.Abs("./")
	fmt.Println(path, err)
}
