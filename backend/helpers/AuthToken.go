package helpers

import (
	"context"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nguyendt456/software-engineer-project/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Auth struct {
	Username string
	Usertype string
	jwt.RegisteredClaims
}

var SECRET = "thisissecrett"

func GenerateAuthToken(username string, usertype string) (string, string, error) {
	var signedClaims = &Auth{
		Username: username,
		Usertype: usertype,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(2))),
		},
	}

	var refreshClaims = &Auth{
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

	res := database.UserCollection.FindOneAndUpdate(ctx, filter, update)
	return res
}

func ValidateAuthToken(clientToken string) (valid bool, err error) {
	token, err := jwt.ParseWithClaims(clientToken, &Auth{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET), nil
		})
	return token.Valid, err
}
