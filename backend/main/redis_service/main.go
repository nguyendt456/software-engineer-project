package main

import (
	"log"
	"net"

	"github.com/nguyendt456/software-engineer-project/main/common"
	pb "github.com/nguyendt456/software-engineer-project/pb"
	"github.com/nguyendt456/software-engineer-project/src/redis_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	CAcert_path = "main/certificate/CAs/ca-cert.pem"
	cert_path   = "main/certificate/server/server.cert"
	key_path    = "main/certificate/server/server.key"
	redis_port  = ":8083"
)

func main() {
	creds := common.LoadCertificateAndKey(CAcert_path, cert_path, key_path)

	grpc_server := grpc.NewServer(grpc.Creds(credentials.NewTLS(creds)))

	pb.RegisterRedisServiceServer(grpc_server, &redis_service.RedisService{})

	lis, err := net.Listen("tcp", redis_port)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	if err := grpc_server.Serve(lis); err != nil {
		log.Fatalf("Error when serve TCP: %s", err.Error())
	}
}
