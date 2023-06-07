package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"auth"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := auth.NewAuthServiceClient(conn)
	res, err := c.Login(context.Background(), &auth.LoginRequest{
		Username: "admin",
		Password: "password",
	})
	if err != nil {
		log.Fatalf("could not login: %v", err)
	}
	if res.Success {
		log.Println("Login successful")
	} else {
		log.Println("Login failed:", res.Message)
	}
}
