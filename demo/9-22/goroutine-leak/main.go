package main

import (
	"os"
	"strconv"
)

func main() {
	processFiles()
}
func processFiles() error {
	for i := 0; i < 100000; i++ {
		file, err := os.Open("file" + strconv.Itoa(i) + ".txt")
		if err != nil {
			return err
		}
		defer file.Close()

		// 处理文件内容
		// ...
	}
	return nil
}
