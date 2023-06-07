package main

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	"google.golang.org/grpc"

	"grpc-demo/auth"
)

func main() {
	//连接grpc服务
	conn, err := grpc.Dial("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//退出关闭连接
	defer conn.Close()
	c := auth.NewAuthServiceClient(conn)
	//设置调用超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	//rpc 调用远程服务的login方法
	res, err := c.Login(ctx, &auth.LoginRequest{
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
