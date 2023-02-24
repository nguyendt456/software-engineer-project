package main

import (
	"log"
	"net"

	"github.com/nguyendt456/software-engineer-project/database"
	pb "github.com/nguyendt456/software-engineer-project/proto"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type UserService struct {
	pb.UserServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Error when starting server:", err)
	}

	database.ConnectToDatabase()

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, &UserService{})

	log.Println("Server started at:", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Error when starting gRPC:", err)
	}
}
