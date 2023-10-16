package main

import (
	"log"
	"runtime"
	"runtime/debug"
)

func init() {
	log.SetFlags(log.LstdFlags)
}

func main() {
	//设置 Go 程序可以使用的最大操作系统线程数,超过设置，程序会崩溃。默认值为10000
	//注意，这个设置限制的是操作系统线程数，也就是GMP中M的数量，而不是Goroutine的数量
	//返回值为之前的设置值
	log.Println("SetMaxThreads:", debug.SetMaxThreads(10000))

	//设置单个 goroutine 堆栈可以使用的最大内存量
	//默认值：64位系统上为1 GB，在32位系统上为250 MB
	//返回值为之前的设置值
	log.Println("SetMaxStack", debug.SetMaxStack(1024*1024*1024))

	//Go1.19增加的软内存限制，跟GOMEMLIMIT环境变量等价
	//debug.SetMemoryLimit()函数用于设置Go程序的内存限制。它可以限制程序在运行时使用的内存量，帮助避免程序使用过多的内存导致系统资源耗尽或崩溃。
	//默认情况下，Go程序没有设置内存限制，可以使用任意数量的内存。但是，在某些情况下，我们可能希望限制程序使用的内存，以确保程序在资源受限的环境中运行时不会过度消耗内存。
	log.Println(debug.SetMemoryLimit())

	//设置逻辑处理器P的个数,跟环境变量`GOMAXPROCS` 等价
	//默认值：对于官方标准编译器，在Go 1.5之前，默认初始逻辑处理器的数量为1；Go 1.5之后，默认初始逻辑处理器的数量和逻辑CPU的数量一致；
	//配置建议：在i/o密集型业务中设置P的个数大于CPU核数是有好处的，涉及到 I/O操作频繁的程序，单纯的计算能力可能并不是瓶颈，而是 I/O 操作的延迟。
	//返回值为之前的设置值
	log.Println(runtime.GOMAXPROCS())

	log.Println(runtime.SetMutexProfileFraction())
	log.Println(runtime.SetBlockProfileRate())
}
