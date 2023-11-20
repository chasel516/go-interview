package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		c := make(chan int)
		for i := 0; i < 100000000; i++ {
			go func(index int) {
				if index%2 == 0 {
					<-c
				} else {
					c <- index
				}
			}(i)
		}
	}()

	server := http.Server{
		Addr: ":8080",
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err.Error())
	}
}
