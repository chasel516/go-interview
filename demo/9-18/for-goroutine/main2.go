package main

import (
	"fmt"
	"log"
	"sync"
)

const Num = 10

type User struct {
	Id   int
	Name string
}

var users = make([]User, 0, Num)

func init() {
	log.SetFlags(log.Lshortfile)
	for i := 1; i <= Num; i++ {
		user := User{
			Id:   i,
			Name: fmt.Sprintf("name%d", i),
		}
		users = append(users, user)
	}
	log.Println("init.users:", users)
}

func (u *User) PrintName() {
	fmt.Println(u.Name)
}

func main() {
	wg := sync.WaitGroup{}
	for _, u := range users {
		wg.Add(1)
		go func() {
			u.PrintName()
			wg.Done()
		}()
	}
	wg.Wait()
}
