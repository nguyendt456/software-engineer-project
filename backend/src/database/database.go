package database

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "github.com/nguyendt456/software-engineer-project/pb"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseService struct {
	pb.DatabaseServer
}

const (
	mongodb_path    = "mongodb://root:project231@0.0.0.0:27017/"
	database_name   = "Project231"
	collection_name = "user"
)

func ConnectDatabase() (*mongo.Collection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	db_client, err := mongo.NewClient(options.Client().ApplyURI(mongodb_path))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = db_client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println("Connected to database")
	database := db_client.Database(database_name)
	collection := database.Collection(collection_name)
	return collection, nil
}

var Collection, _ = ConnectDatabase()

func (service *DatabaseService) InsertUser(client_ctx context.Context, client_user *pb.User) (*pb.Response, error) {
	var search_user pb.User
	find_result := Collection.FindOne(client_ctx, bson.D{
		{
			Key:   "username",
			Value: client_user.Username,
		},
	}).Decode(&search_user)
	if find_result == nil {
		return &pb.Response{
			StatusCode: int32(http.StatusInternalServerError),
			Message:    "Duplicate username",
		}, errors.New("Duplicate username")
	}

	result, err := Collection.InsertOne(client_ctx, client_user)

	result_string := fmt.Sprintf("%v", result)
	if err != nil {
		return &pb.Response{
			StatusCode: int32(http.StatusInternalServerError),
			Message:    result_string,
		}, err
	}
	return &pb.Response{
		StatusCode: int32(http.StatusCreated),
		Message:    "User Created",
	}, nil
}

func (service *DatabaseService) UpdateUserToken(ctx context.Context, user_token *pb.UserToken) (*pb.Response, error) {
	var filter = bson.M{
		"username": user_token.Username,
	}

	var update = bson.M{
		"$set": bson.M{
			"signedtoken":  user_token.Token,
			"refreshtoken": user_token.RefreshToken,
		},
	}

	var option = options.After
	res := Collection.FindOneAndUpdate(ctx, filter, update, &options.FindOneAndUpdateOptions{ReturnDocument: &option})

	if res.Err() != nil {
		return &pb.Response{
			StatusCode: int32(http.StatusInternalServerError),
			Message:    "",
		}, res.Err()
	}
	var response = pb.UserToken{}
	res.Decode(&response)
	return &pb.Response{
		StatusCode: int32(http.StatusAccepted),
		Message:    response.Token,
	}, nil
}

func (service *DatabaseService) GetUserByUsername(client_ctx context.Context, username *pb.Username) (*pb.User, error) {
	var user = pb.User{}
	var filter = bson.D{
		{
			Key:   "username",
			Value: username.Username,
		},
	}
	res := Collection.FindOne(client_ctx, filter)
	if res.Err() != nil {
		return &pb.User{}, res.Err()
	}
	res.Decode(&user)
	return &user, nil
}
