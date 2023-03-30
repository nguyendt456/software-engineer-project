package common

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nguyendt456/software-engineer-project/src/setup_env"
	"golang.org/x/crypto/bcrypt"
)

const (
	secretKey = "testing"
)

func HashPassword(password string) string {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Panic(err)
	}
	return string(hashed_password)
}

func CompareHashPassword(user_password string, hashed_password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(user_password))
	if err == nil {
		return true
	}
	return false
}

type AuthClaims struct {
	Username string
	Usertype string
	Name     string
	jwt.RegisteredClaims
}

func ValidateToken(client_token string) (userSignedDetail AuthClaims, valid bool, err error) {
	token, err := jwt.ParseWithClaims(client_token, &AuthClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
	fmt.Println(token.Claims)
	if err != nil {
		return AuthClaims{}, false, err
	}
	return userSignedDetail, token.Valid, err
}

func GenerateAuthToken(userToAuth AuthClaims) (string, string, error) {
	var signedClaims = userToAuth
	signedClaims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(setup_env.TokenDuration)),
	}

	signedToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, signedClaims).SignedString([]byte(secretKey))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}

	signedClaims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(setup_env.RefreshTokenDuration)),
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, signedClaims).SignedString([]byte(secretKey))
	if err != nil {
		log.Panic(err)
		return "", "", err
	}
	return signedToken, refreshToken, err
}
