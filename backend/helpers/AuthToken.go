package helpers

import (
	"context"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nguyendt456/software-engineer-project/database"
	"github.com/nguyendt456/software-engineer-project/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Auth struct {
	Name         string
	Username     string
	Usertype     string
	UserIdentity string
	Birthday     string
	jwt.RegisteredClaims
}

var SECRET = "thisissecrett"

func GenerateAuthToken(userToAuth models.User) (string, string, error) {
	var signedClaims = userToAuth
	signedClaims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(2))),
	}

	var refreshClaims = models.User{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(24))),
		},
	}

	signedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, signedClaims).SignedString([]byte(SECRET))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}
	return signedToken, refreshToken, err
}

func UpdateAuthToken(username string, signedToken string, refreshToken string) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	filter := bson.M{
		"username": username,
	}

	update := bson.M{
		"$set": bson.M{
			"signedtoken":  signedToken,
			"refreshtoken": refreshToken,
		},
	}

	var option = options.After

	res := database.UserCollection.FindOneAndUpdate(ctx, filter, update, &options.FindOneAndUpdateOptions{ReturnDocument: &option})
	return res
}

func ValidateAuthToken(clientToken string) (userSignedDetail models.User, valid bool, err error) {
	token, err := jwt.ParseWithClaims(clientToken, &userSignedDetail,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET), nil
		})
	if err != nil {
		return models.User{}, false, err
	}
	return userSignedDetail, token.Valid, err
}
