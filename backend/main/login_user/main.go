package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/nguyendt456/software-engineer-project/pb"
	"github.com/nguyendt456/software-engineer-project/src/login_user"
	"google.golang.org/grpc"
)

const (
	Login_user_port    = ":8084"
	Login_user_gw_port = ":8085"
)

func CustomMatcher(key string) (string, bool) {
	switch key {
	case "Token":
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

func startLoginUserGateway() {
	grpc_mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(CustomMatcher),
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := pb.RegisterLoginUserServiceHandlerFromEndpoint(ctx, grpc_mux, "0.0.0.0"+Login_user_port, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal("cannot register handler server")
	}

	http_mux := http.NewServeMux()

	http_mux.Handle("/", grpc_mux)

	lis, err := net.Listen("tcp", Login_user_gw_port)
	if err != nil {
		log.Fatal("failed to listen on tcp port:", err.Error())
	}
	log.Println("Starting gRPC gateway endpoint")

	if err := http.Serve(lis, http_mux); err != nil {
		log.Fatal("cannot start HTTP gateway server:", err)
	}
}

func main() {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(login_user.JWTAuthInterceptor),
	)
	pb.RegisterLoginUserServiceServer(grpcServer, &login_user.LoginUserService{})

	go startLoginUserGateway()

	lis, err := net.Listen("tcp", Login_user_port)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err.Error())
		return
	}
}
