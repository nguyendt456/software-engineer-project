package main

import (
	"context"
	"log"

	pb "github.com/nguyendt456/software-engineer-project/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	grpc_conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error when connect to gRPC server:", err)
	}
	defer grpc_conn.Close()

	client := pb.NewUserServiceClient(grpc_conn)

	res, err := client.CreateUser(context.Background(), &pb.User{
		Username: "nguyendt456",
		Password: "1234567891",
		Usertype: "backofficer",
	})

	if err != nil {
		log.Fatal("Error when sending gRPC request:", err.Error())
	}
	log.Println(res)
}
