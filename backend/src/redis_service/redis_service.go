package redis_service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/nguyendt456/software-engineer-project/main/common"
	pb "github.com/nguyendt456/software-engineer-project/pb"
	"github.com/nguyendt456/software-engineer-project/src/setup_env"
	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	pb.RedisServiceServer
}

const (
	CAcert_path = "main/certificate/CAs/ca-cert.pem"
	cert_path   = "main/certificate/server/server.cert"
	key_path    = "main/certificate/server/server.key"
)

var creds = common.LoadCertificateAndKey(
	CAcert_path,
	cert_path,
	key_path,
)

var redis_client = redis.NewClient(&redis.Options{
	Addr:        setup_env.Redis_addr,
	Password:    "",
	DB:          0,
	TLSConfig:   creds,
	PoolSize:    100,
	PoolTimeout: time.Second * 5,
})

func (redis_service *RedisService) SetUUID(client_ctx context.Context, user_id *pb.UserID) (*pb.Response, error) {
	for _, value := range user_id.Value {
		fmt.Println(value)
		err := redis_client.LPush(client_ctx, user_id.Key, value).Err()
		if err != nil {
			return &pb.Response{
				StatusCode: int32(http.StatusInternalServerError),
				Message:    "Error when operate Redis",
			}, err
		}
		err = redis_client.Expire(client_ctx, user_id.Key, setup_env.TokenDuration).Err()
		if err != nil {
			return &pb.Response{
				StatusCode: int32(http.StatusInternalServerError),
				Message:    "Error when operate Redis",
			}, err
		}
	}
	return &pb.Response{
		StatusCode: int32(http.StatusAccepted),
	}, nil
}

func (redis_service *RedisService) RemoveUUID(client_ctx context.Context, user_id *pb.UserID) (*pb.Response, error) {
	err := redis_client.Del(client_ctx, user_id.Key).Err()
	if err != nil {
		return &pb.Response{
			StatusCode: int32(http.StatusInternalServerError),
		}, err
	}
	return &pb.Response{
		StatusCode: int32(http.StatusAccepted),
	}, nil
}

func (redis_service *RedisService) CheckExistedUserSession(client_ctx context.Context, user *pb.UserID) (*pb.LoginResponse, error) {
	signed_token, err := redis_client.LIndex(client_ctx, user.Key, 0).Result()
	if err != nil {
		return &pb.LoginResponse{
			StatusCode: int32(http.StatusNotFound),
		}, err
	}
	refresh_token, err := redis_client.LIndex(client_ctx, user.Key, 1).Result()
	if err != nil {
		return &pb.LoginResponse{
			StatusCode: int32(http.StatusNotFound),
		}, err
	}
	return &pb.LoginResponse{
		StatusCode:   int32(http.StatusAccepted),
		Token:        signed_token,
		RefreshToken: refresh_token,
	}, nil
}
