package main

import (
	"runtime"
	"runtime/debug"
)

func main() {
	debug.SetMaxThreads()
	debug.SetMaxStack()
	debug.SetMemoryLimit()
	debug.SetPanicOnFault()
	runtime.GOMAXPROCS()
	runtime.SetCPUProfileRate()
	runtime.SetMutexProfileFraction()
	runtime.SetBlockProfileRate()
}
