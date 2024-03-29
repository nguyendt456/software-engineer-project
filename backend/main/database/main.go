package main

import (
	"log"
	"net"

	"github.com/nguyendt456/software-engineer-project/main/common"
	pb "github.com/nguyendt456/software-engineer-project/pb"
	"github.com/nguyendt456/software-engineer-project/src/database"
	"github.com/nguyendt456/software-engineer-project/src/setup_env"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	CAcert_path = "main/certificate/CAs/ca-cert.pem"
	cert_path   = "main/certificate/server/server.cert"
	key_path    = "main/certificate/server/server.key"
)

func main() {
	creds := common.LoadCertificateAndKey(CAcert_path, cert_path, key_path)

	grpc_server := grpc.NewServer(grpc.Creds(credentials.NewTLS(creds)))

	pb.RegisterDatabaseServer(grpc_server, &database.DatabaseService{})

	lis, err := net.Listen("tcp", setup_env.MongoDB_service_addr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	if err := grpc_server.Serve(lis); err != nil {
		log.Fatalf("Error when serve TCP: %s", err.Error())
	}
}
