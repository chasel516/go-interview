package main

import (
	"fmt"
	"log"
	"sync"
)

const Num = 3

type User struct {
	Id   int
	Name string
}

var users = make([]User, 0, Num)
var pusers = make([]*User, 0, Num)

func init() {
	log.SetFlags(log.Lshortfile)
	for i := 1; i <= Num; i++ {
		user := User{
			Id:   i,
			Name: fmt.Sprintf("name%d", i),
		}
		users = append(users, user)
		pusers = append(pusers, &user)
	}
	log.Println("init.users:", users)
}

func main() {
	//fn1()
	//fn2()
	//f1()
	//f2()
	//f3()
	//f4()
	//f5()
	//f6()
	f7()
	//f8()

}

// for range取址问题
func fn1() {
	s := make([]*User, 0, Num)
	for _, user := range users {
		log.Printf("user.p:%p", &user)
		//使用下标进行追加
		s = append(s, &user)
	}
	log.Println(*s[0], *s[1])
}

func fn2() {
	s := make([]*User, 0, Num)
	for _, user := range pusers {
		log.Printf("user.p:%p", &user)
		s = append(s, user)
	}
	//log.Println(s)
	log.Println(*s[0], *s[1])
}

// 闭包延迟绑定问题
func f1() {
	var fns []func()
	for _, user := range users {
		fns = append(fns, func() {
			log.Println(user)
		})
	}
	for _, fn := range fns {
		fn()
	}
}

func f2() {
	for _, user := range users {
		func() {
			log.Println(user)
		}()
	}
}

// goroutine延迟绑定问题
// 输出的i会出现重复
func f3() {
	//runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			log.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}

// 解决方式1
func f4() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		x := i
		log.Printf("i.ptr:%p;x.ptr:%p", &i, &x)
		wg.Add(1)
		go func() {
			log.Println(x)
			wg.Done()
		}()
	}
	wg.Wait()
}

// 解决方式2
func f5() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(param int) {
			log.Printf("param.ptr:%p;param.val:%d", &param, param)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// 解决方式3
func f6() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go print(i, &wg)
		//这样能确保上面的go语句执行完成？

	}
	wg.Wait()
}

func print(v int, wg *sync.WaitGroup) {
	log.Println(v)
	wg.Done()
}

// 闭包捕获的变量在goroutine启动之后被修改，闭包中的值也会被修改
func f7() {
	messages := make(chan User, Num)
	for i := 1; i <= Num; i++ {
		msg := User{
			Id:   i,
			Name: fmt.Sprintf("name-%d", i),
		}

		//goroutine中的变量可能被外部代码修改
		go func() {
			messages <- msg
		}()
		//time.Sleep(time.Millisecond * 100)
		doSomething(&msg)
	}

	for i := 1; i <= Num; i++ {
		log.Println(<-messages)
	}

	close(messages)
}

func doSomething(user *User) {
	//do something ...
	user.Id = 100
}

// 将变量作为参数传递给闭包：
func f8() {
	messages := make(chan User, Num)
	for i := 1; i <= Num; i++ {
		msg := User{
			Id:   i,
			Name: fmt.Sprintf("name-%d", i),
		}

		go func(user User) {
			//time.Sleep(time.Millisecond)
			messages <- user
		}(msg)
		doSomething(&msg)
	}

	for i := 1; i <= Num; i++ {
		log.Println(<-messages)
	}

	close(messages)
}
