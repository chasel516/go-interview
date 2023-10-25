package main

import (
	"fmt"
	"sync"
)

func main() {
	//f1()
	//f2()
	//f3()
	//f4()
	//f5()
	f6()

}

//输出的i会出现重复
func f1() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

//解决方式1
func f2() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		x := i
		wg.Add(1)
		go func() {
			fmt.Println(x)
			wg.Done()
		}()
	}
}

//解决方式2
func f3() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
}

//解决方式3
func f4() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go fmt.Println(i)
		//这样能确保上面的go语句执行完成？
		wg.Done()

	}
}

type User struct {
	Id   int
	Name string
}

//闭包捕获的变量在goroutine启动之后被修改，闭包中的值也会被修改
func f5() {
	const Num = 3

	messages := make(chan User, Num)
	for i := 0; i < Num; i++ {
		msg := User{
			Id:   i,
			Name: fmt.Sprintf("name-%d", i),
		}

		//goroutine中的变量可能被外部代码修改
		go func() {
			//time.Sleep(time.Millisecond)
			messages <- msg
		}()
		doSomething(&msg)
	}

	for i := 0; i < Num; i++ {
		fmt.Println(<-messages)
	}

	close(messages)
}

func doSomething(user *User) {
	//do something ...
	user.Id = 100
}

//将变量作为参数传递给闭包：
func f6() {
	const Num = 3

	messages := make(chan User, Num)
	for i := 0; i < Num; i++ {
		msg := User{
			Id:   i,
			Name: fmt.Sprintf("name-%d", i),
		}

		//goroutinene闭包中的msg与外面的msg引用着相同的间接值部
		go func(user User) {
			//time.Sleep(time.Millisecond)
			messages <- msg
		}(msg)
		doSomething(&msg)
	}

	for i := 0; i < Num; i++ {
		fmt.Println(<-messages)
	}

	close(messages)
}
