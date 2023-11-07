package main

import (
	"fmt"
	"log"
	"time"
)

const Num = 10

type User struct {
	Id   int
	Name string
}

var users = make([]User, 0, Num)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	//f8()
	f9()
}
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
		go doSomething(&msg)
	}

	for i := 1; i <= Num; i++ {
		log.Println(<-messages)
	}

	close(messages)
	time.Sleep(time.Second)
}

// 将变量的引用类型作为参数传递给闭包：
func f9() {
	messages := make(chan User, Num)
	for i := 1; i <= Num; i++ {
		msg := User{
			Id:   i,
			Name: fmt.Sprintf("name-%d", i),
		}

		go func(user *User) {
			//time.Sleep(time.Millisecond)
			messages <- *user
		}(&msg)
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
