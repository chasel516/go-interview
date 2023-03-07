package main

import (
	"fmt"
)

type ReadErr struct {
	error
}

func main() {
	err := read()
	println(err)
	if err != nil {
		fmt.Println("has error")
	} else {
		fmt.Println("no error")
	}
}

func read() error {
	var err *ReadErr
	return err
}
