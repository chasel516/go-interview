package main

import "github.com/phper95/pkg/routine"

const PollNameDefault = "default"

func main() {
	routine.InitPoolWithName(PollNameDefault, 1, 1, 0)
	for i := 0; i < 100; i++ {
		routine.
	}
}
