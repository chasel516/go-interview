package main

import (
	"fmt"
	"time"
)

type User struct {
	Id   int
	Name string
}

func main() {
	users := make([]User, 0, 10)
	for i := 0; i < 10; i++ {
		users = append(users, User{Id: i, Name: fmt.Sprintf("name_%d", i)})
	}

	for _, user := range users {
		go func(u User) {
			writeDB(u)
		}(user)
	}

	time.Sleep(time.Second)
}

func writeDB(user User) {
	fmt.Println(user)
}

func writeDB1(user *User) {
	fmt.Println(*user)
}
