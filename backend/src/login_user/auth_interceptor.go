package login_user

import (
	"context"
	"fmt"

	"github.com/nguyendt456/software-engineer-project/main/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func JWTAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	token, ok := md["token"]
	if !ok {
		resp, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}

		return resp, nil
	}
	AuthInfo, valid, err := common.ValidateToken(token[0])
	if !valid {
		fmt.Println("invalid token")
	}
	ctx = context.WithValue(ctx, "User", AuthInfo)

	resp, err := handler(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
