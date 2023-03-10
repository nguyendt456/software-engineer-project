package main

import (
	"log"
	"net"

	"github.com/nguyendt456/software-engineer-project/main/common"
	pb "github.com/nguyendt456/software-engineer-project/pb"
	"github.com/nguyendt456/software-engineer-project/src/database"
	"google.golang.org/grpc"
)

const (
	CAcert_path   = "main/certificate/CAs/ca-cert.pem"
	cert_path     = "main/certificate/server/server.cert"
	key_path      = "main/certificate/server/server.key"
	database_port = ":8080"
)

func main() {
	creds, err := common.LoadCertificateAndKey(CAcert_path, cert_path, key_path)
	if err != nil {
		log.Fatal(err)
		return
	}
	grpc_server := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterDatabaseServer(grpc_server, &database.DatabaseService{})

	lis, err := net.Listen("tcp", database_port)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	if err := grpc_server.Serve(lis); err != nil {
		log.Fatalf("Error when serve TCP: %s", err.Error())
	}
}
