package main

import (
	"fmt"
)

type MyErr struct {
	code int
	Msg  string
}

func (e *MyErr) Error() string {
	return fmt.Sprintf("code:%d;msg:%s", e.code, e.Msg)
}

func main() {
	var e error
	e = &MyErr{
		code: 1,
		Msg:  "server error",
	}
	fmt.Println(e)

}
