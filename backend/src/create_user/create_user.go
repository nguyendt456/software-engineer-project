package create_user

import (
	"context"

	"github.com/nguyendt456/software-engineer-project/main/common"
	pb "github.com/nguyendt456/software-engineer-project/pb"
	"github.com/nguyendt456/software-engineer-project/src/setup_env"
	"github.com/nguyendt456/software-engineer-project/src/validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type CreateUserService struct {
	pb.CreateUserServiceServer
}

const (
	CAcert_path = "main/certificate/CAs/ca-cert.pem"
	cert_path   = "main/certificate/client/client.cert"
	key_path    = "main/certificate/client/client.key"
)

func (service *CreateUserService) CreateUser(client_ctx context.Context, user *pb.User) (*pb.Response, error) {
	err := validator.ValidateUser(user)
	if err != nil {
		return &pb.Response{}, err
	}
	var creds = common.LoadCertificateAndKey(
		CAcert_path,
		setup_env.Cert_path,
		setup_env.Key_path,
	)

	conn, err := grpc.Dial(setup_env.MongoDB_service_addr, grpc.WithTransportCredentials(credentials.NewTLS(creds)))
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
