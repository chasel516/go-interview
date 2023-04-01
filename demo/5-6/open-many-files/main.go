package main

import (
	"os"
	"strconv"
	"sync"
)

func main() {

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

func processFiles1() error {
	for i := 0; i < 100000; i++ {
		processFile("file" + strconv.Itoa(i) + ".txt")
	}
	return nil
}

func processFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 处理文件内容
	// ...
	return nil
}

func processFiles2() error {
	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			processFile("file" + strconv.Itoa(i) + ".txt")
		}(i)
	}
	wg.Wait()
	return nil
}
