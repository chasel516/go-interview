package main

import (
	"context"
	"log"
	"net"
	// 导入grpc包
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	// 导入刚才我们生成的代码所在的auth包
	"grpc-demo/auth"
)

type authServer struct {
	auth.UnimplementedAuthServiceServer
}

func (s *authServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	log.Printf("Received: %v", req.GetUsername())
	if req.Username == "admin" && req.Password == "password" {
		return &auth.LoginResponse{
			Success: true,
			Message: "Login successful",
		}, nil
	} else {
		return &auth.LoginResponse{
			Success: false,
			Message: "Invalid username or password",
		}, nil
	}
}

func main() {
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//初始化grpc server
	s := grpc.NewServer()

	//注册auth server
	auth.RegisterAuthServiceServer(s, &authServer{})
	reflection.Register(s)
	//
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("grpc server started")
}
