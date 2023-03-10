package login_user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nguyendt456/software-engineer-project/main/common"
	pb "github.com/nguyendt456/software-engineer-project/pb"
	"github.com/nguyendt456/software-engineer-project/src/validator"
	"google.golang.org/grpc"
)

const (
	CAcert_path   = "main/certificate/CAs/ca-cert.pem"
	cert_path     = "main/certificate/client/client.cert"
	key_path      = "main/certificate/client/client.key"
	database_port = ":8080"
)

type LoginUserService struct {
	pb.LoginUserServiceServer
}

func (login_user *LoginUserService) LoginUser(client_ctx context.Context, login_form *pb.LoginForm) (*pb.Response, error) {
	err := validator.ValidateLoginForm(login_form)
	if err != nil {
		return &pb.Response{
			StatusCode: int32(http.StatusUnauthorized),
			Message:    err.Error(),
		}, err
	}

	creds, err := common.LoadCertificateAndKey(CAcert_path, cert_path, key_path)
	if err != nil {
		return &pb.Response{
			StatusCode: int32(http.StatusInternalServerError),
		}, err
	}
	conn, err := grpc.Dial("0.0.0.0"+database_port, grpc.WithTransportCredentials(creds))
	defer conn.Close()
	if err != nil {
		return &pb.Response{
			StatusCode: int32(http.StatusInternalServerError),
		}, err
	}

	client := pb.NewDatabaseClient(conn)
	response, err := client.GetUserByUsername(client_ctx, &pb.Username{Username: login_form.Username})
	if err != nil {
		return &pb.Response{
			StatusCode: int32(http.StatusInternalServerError),
		}, err
	}

	if ok := common.CompareHashPassword(login_form.Password, response.Password); ok == false {
		return &pb.Response{
			StatusCode: int32(http.StatusUnauthorized),
			Message:    "Wrong username or password",
		}, fmt.Errorf("Wrong username or password")
	}

	signed_token, refresh_token, err := common.GenerateAuthToken(
		common.AuthClaims{
			Username: response.Username,
			Usertype: response.Usertype,
			Name:     response.Name,
		},
	)

	_, err = client.UpdateUserToken(client_ctx,
		&pb.UserToken{
			Username:     response.Username,
			Token:        signed_token,
			RefreshToken: refresh_token,
		},
	)
	if err != nil {
		return &pb.Response{
			StatusCode: int32(http.StatusUnauthorized),
			Message:    err.Error(),
		}, err
	}
	return &pb.Response{
		StatusCode: int32(http.StatusAccepted),
		Message:    signed_token,
	}, nil
}
