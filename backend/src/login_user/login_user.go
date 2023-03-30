package login_user

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"sync"

	"github.com/nguyendt456/software-engineer-project/main/common"
	pb "github.com/nguyendt456/software-engineer-project/pb"
	"github.com/nguyendt456/software-engineer-project/src/setup_env"
	"github.com/nguyendt456/software-engineer-project/src/validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type LoginUserService struct {
	pb.LoginUserServiceServer
}

var (
	mu    sync.Mutex
	creds *tls.Config
)

func TLSLoad() {
	if creds == nil {
		mu.Lock()
		defer mu.Unlock()
		creds = common.LoadCertificateAndKey(
			setup_env.CAcert_path,
			setup_env.Cert_path,
			setup_env.Key_path,
		)
	}
}

func (login_user *LoginUserService) LoginUser(client_ctx context.Context, login_form *pb.LoginForm) (*pb.LoginResponse, error) {
	err := validator.ValidateLoginForm(login_form)
	if err != nil {
		return &pb.LoginResponse{
			StatusCode: int32(http.StatusUnauthorized),
		}, err
	}

	TLSLoad()

	user_info, err := GetUserInfo(client_ctx, login_form)
	if err != nil {
		return &pb.LoginResponse{
			StatusCode: int32(http.StatusInternalServerError),
		}, err
	}

	if ok := common.CompareHashPassword(login_form.Password, user_info.Password); ok == false {
		return &pb.LoginResponse{
			StatusCode: int32(http.StatusUnauthorized),
		}, fmt.Errorf("Wrong username or password")
	}

	signed_token, refresh_token, err := RedisCheckUserSession(client_ctx, &pb.UserID{Key: user_info.Username})
	if err == nil {
		return &pb.LoginResponse{
			StatusCode:   int32(http.StatusAccepted),
			Token:        signed_token,
			RefreshToken: refresh_token,
		}, nil
	}

	signed_token, refresh_token, err = common.GenerateAuthToken(
		common.AuthClaims{
			Username: user_info.Username,
			Usertype: user_info.Usertype,
			Name:     user_info.Name,
		},
	)
	if err != nil {
		return &pb.LoginResponse{
			StatusCode: int32(http.StatusUnauthorized),
		}, err
	}

	err = RedisSetUserSession(client_ctx, user_info, signed_token, refresh_token)
	if err != nil {
		fmt.Println(err)
		return &pb.LoginResponse{
			StatusCode: int32(http.StatusInternalServerError),
		}, err
	}

	return &pb.LoginResponse{
		StatusCode:   int32(http.StatusAccepted),
		Token:        signed_token,
		RefreshToken: refresh_token,
	}, nil
}

func GetUserInfo(ctx context.Context, login_info *pb.LoginForm) (*pb.User, error) {
	database_conn, err := grpc.Dial(setup_env.MongoDB_service_addr, grpc.WithTransportCredentials(credentials.NewTLS(creds)))
	defer database_conn.Close()
	if err != nil {
		return nil, err
	}
	client := pb.NewDatabaseClient(database_conn)
	response, err := client.GetUserByUsername(ctx, &pb.Username{Username: login_info.Username})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func RedisCheckUserSession(ctx context.Context, user *pb.UserID) (string, string, error) {
	redis_conn, err := grpc.Dial(setup_env.Redis_service_addr, grpc.WithTransportCredentials(credentials.NewTLS(creds)))
	defer redis_conn.Close()
	if err != nil {
		return "", "", err
	}
	redis_client := pb.NewRedisServiceClient(redis_conn)
	login_form, err := redis_client.CheckExistedUserSession(ctx, user)
	if err != nil {
		return "", "", err
	}
	return login_form.Token, login_form.RefreshToken, nil
}

func RedisSetUserSession(ctx context.Context, user *pb.User, signed_token string, refresh_token string) error {
	redis_conn, err := grpc.Dial(setup_env.Redis_service_addr, grpc.WithTransportCredentials(credentials.NewTLS(creds)))
	defer redis_conn.Close()
	if err != nil {
		return err
	}
	redis_client := pb.NewRedisServiceClient(redis_conn)
	_, err = redis_client.SetUUID(ctx, &pb.UserID{
		Key: user.Username,
		Value: []string{
			signed_token,
			refresh_token,
		},
	})
	if err != nil {
		return err
	}
	return nil
}
