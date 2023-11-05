package main

import (
	"fmt"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	//lock := sync.Mutex{}
	//lock.Lock()
	//lock.Lock()
	panic(1)
}
