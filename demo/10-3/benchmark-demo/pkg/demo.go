package pkg

import (
	"time"
)

func f1() {
	time.Sleep(time.Millisecond * 10)
}

func f2() {
	time.Sleep(time.Millisecond * 20)
}
