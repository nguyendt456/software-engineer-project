package create_user

import (
	"context"
	"fmt"

	"github.com/nguyendt456/software-engineer-project/main/common"
	pb "github.com/nguyendt456/software-engineer-project/pb"
	"github.com/nguyendt456/software-engineer-project/src/validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type CreateUserService struct {
	pb.CreateUserServiceServer
}

const (
	CAcert_path   = "main/certificate/CAs/ca-cert.pem"
	cert_path     = "main/certificate/client/client.cert"
	key_path      = "main/certificate/client/client.key"
	database_port = ":8080"
)

func (service *CreateUserService) CreateUser(client_ctx context.Context, user *pb.User) (*pb.Response, error) {
	md, _ := metadata.FromIncomingContext(client_ctx)
	fmt.Println(md.Get("token")[0])
	err := validator.ValidateUser(user)
	if err != nil {
		return &pb.Response{}, err
	}

	creds, err := common.LoadCertificateAndKey(CAcert_path, cert_path, key_path)
	if err != nil {
		return &pb.Response{}, err
	}
	conn, err := grpc.Dial("0.0.0.0"+database_port, grpc.WithTransportCredentials(creds))
	defer conn.Close()
	if err != nil {
		return &pb.Response{}, err
	}

	client := pb.NewDatabaseClient(conn)
	user.Password = common.HashPassword(user.Password)
	response, err := client.InsertUser(client_ctx, user)
	if err != nil {
		return &pb.Response{}, err
	}
	return response, nil
}
