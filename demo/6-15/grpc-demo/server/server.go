package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"auth"
)

type authServer struct{}

func (s *authServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
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
	s := grpc.NewServer()
	auth.RegisterAuthServiceServer(s, &authServer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
