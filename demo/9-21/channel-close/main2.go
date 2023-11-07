package main

import (
	"channel-close/channel"
	"fmt"
)

func main() {
	//c := channel.NewChannel()
	//c.Close()
	//c.Close()
	//_, ok := <-c.C
	//fmt.Println(ok)

	c2 := channel.NewChannel2()
	c2.Close()
	c2.Close()
	_, ok := <-c2.C
	fmt.Println(ok)

}
