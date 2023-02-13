package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nguyendt456/software-engineer-project/database"
	"github.com/nguyendt456/software-engineer-project/helpers"
	"github.com/nguyendt456/software-engineer-project/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(
				http.StatusBadRequest,
				models.Response{
					Status:  http.StatusBadRequest,
					Message: "Bad request",
					Detail:  map[string]interface{}{"detail": err.Error()},
				},
			)
			return
		}

		var userToAuth models.User
		err := database.UserCollection.FindOne(ctx, bson.D{
			{
				Key:   "username",
				Value: user.UserName,
			},
		}).Decode(&userToAuth)

		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				models.Response{
					Status:  http.StatusInternalServerError,
					Message: "Internal Server Error",
					Detail:  map[string]interface{}{"detail": err.Error()},
				},
			)
		}

		res := bcrypt.CompareHashAndPassword([]byte(userToAuth.Password), []byte(user.Password))

		if res != nil {
			c.JSON(
				http.StatusBadRequest,
				models.Response{
					Status:  http.StatusBadRequest,
					Message: "Wrong username or password",
					Detail: map[string]interface{}{
						"Username": user.UserName,
					},
				},
			)
			return
		}
		signedToken, refreshToken, _ := helpers.GenerateAuthToken(userToAuth.UserName, userToAuth.UserType)

		result := helpers.UpdateAuthToken(userToAuth.UserName, signedToken, refreshToken)
		if result.Err() != nil {
			fmt.Println(result.Err())
			return
		}
		result.Decode(&user)
		c.JSON(http.StatusOK, user)
	}
}
