package main

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/nguyendt456/software-engineer-project/database"
	pb "github.com/nguyendt456/software-engineer-project/proto"
)

func (service_server *UserService) CreateUser(ctx context.Context, user *pb.User) (*pb.UserResponse, error) {
	err := validator(user)

	if err != nil {
		return &pb.UserResponse{
			StatusCode: int32(http.StatusBadRequest),
			Message:    err.Error(),
		}, err
	} else {
		insertation_context, insertion_cancel := context.WithTimeout(ctx, time.Second*2)
		defer insertion_cancel()
		_, err := database.UserCollection.InsertOne(insertation_context, user)
		if err != nil {
			return &pb.UserResponse{
				StatusCode: int32(http.StatusBadRequest),
				Message:    err.Error(),
			}, err
		}
		return &pb.UserResponse{
			StatusCode: int32(http.StatusCreated),
			Message:    "User created",
		}, nil
	}
}

func validator(u *pb.User) (err error) {
	if u.Username == "" {
		return errors.New("Username empty")
	}
	if len(u.Username) < 5 {
		return errors.New("Username (min:5)")
	}
	if u.Password == "" {
		return errors.New("Password empty")
	}
	if len(u.Password) < 8 {
		return errors.New("Password (min:8)")
	}
	if u.Usertype != "janitor" && u.Usertype != "collector" && u.Usertype != "backofficer" {
		return errors.New("Invalid Usertype")
	}
	return nil
}
