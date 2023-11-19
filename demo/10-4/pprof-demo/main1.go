package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		for {
		}
	}()
	go func() {
		test()
	}()
	server := http.Server{
		Addr: ":8080",
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Println(err.Error())
	}
}

func test() {
	s := [10240000]byte{}
	_ = string(s[:])
	for {
	}
}
