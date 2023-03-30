package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/nguyendt456/software-engineer-project/pb"
	"github.com/nguyendt456/software-engineer-project/src/login_user"
	"github.com/nguyendt456/software-engineer-project/src/setup_env"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
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
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
		runtime.WithIncomingHeaderMatcher(CustomMatcher),
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := pb.RegisterLoginUserServiceHandlerFromEndpoint(ctx, grpc_mux, setup_env.Login_user_addr, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal("cannot register handler server")
	}

	http_mux := http.NewServeMux()

	http_mux.Handle("/", grpc_mux)

	lis, err := net.Listen("tcp", setup_env.Login_user_gw)
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

	lis, err := net.Listen("tcp", setup_env.Login_user_addr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err.Error())
		return
	}
}
