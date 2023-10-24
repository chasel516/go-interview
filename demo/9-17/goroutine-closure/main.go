package main

import "fmt"

func main()  {
	f1()
	f2()
	f3()
	f4()

}

//输出的i会出现重复
func f1()  {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}


//解决方式1
func f2()  {
	for i := 0; i < 10; i++ {
		x:=i
		go func() {
			fmt.Println(x)
		}()
	}
}

//解决方式2
func f3()  {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}

//解决方式3
func f4()  {
	for i := 0; i < 10; i++ {
		go fmt.Println(i)
	}
}








1. 闭包中变量的生命周期：闭包可以捕获其外部作用域的变量，但是需要注意变量的生命周期。如果闭包捕获的变量在goroutine启动之后被修改，那么闭包中的值也会被修改。

```
func main() {
    messages := make(chan string)

    for i := 0; i < 3; i++ {
        msg := fmt.Sprintf("message-%d", i)

        go func() {
            messages <- msg
        }()
    }

    for i := 0; i < 3; i++ {
        fmt.Println(<-messages)
    }
}
```

这个例子中，虽然看似为每一个goroutine创建了一个新的`msg`变量，但实际上它们共享了同一个变量。为了解决这个问题，可以将变量作为参数传递给闭包：

```
func main() {
    messages := make(chan string)

    for i := 0; i < 3; i++ {
        msg := fmt.Sprintf("message-%d", i)

        go func(msg string) {
            messages <- msg
        }(msg)
    }

    for i := 0; i < 3; i++ {
        fmt.Println(<-messages)
    }
}
```

1. 并发访问共享资源：在使用goroutine和闭包时，需要注意并发访问共享资源，例如map、切片等。在并发场景下，需要使用互斥锁（sync.Mutex）或者原子操作（sync/atomic）来保护共享资源。
2. 同步与顺序问题：在使用goroutine和闭包时，需要注意同步问题。如果需要等待所有的goroutine执行完毕，可以使用sync.WaitGroup来实现同步。同时，如果需要顺序执行goroutine，可以使用通道（channel）来实现。

总之，在使用goroutine和闭包时，需要注意循环变量捕获、变量生命周期、并发访问共享资源以及同步与顺序问题。遵循这些原则，可以避免很多潜在的错误。