package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nguyendt456/software-engineer-project/database"
	"github.com/nguyendt456/software-engineer-project/helpers"
	"github.com/nguyendt456/software-engineer-project/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Validator = validator.New()

func AddUser() gin.HandlerFunc {
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

		if err := Validator.Struct(user); err != nil {
			validationErrArray := err.(validator.ValidationErrors)
			var validationErrFields []string
			for _, validationErr := range validationErrArray {
				validationErrFields = append(validationErrFields, validationErr.Field())
			}

			c.JSON(
				http.StatusBadRequest,
				models.Response{
					Status:  http.StatusBadRequest,
					Message: "Validation failed",
					Detail: map[string]interface{}{
						"error_field": validationErrFields,
						"detail":      validationErrArray.Error(),
					},
				},
			)
			return
		}

		err := database.UserCollection.FindOne(ctx, bson.D{
			{
				Key:   "username",
				Value: user.UserName,
			},
		}).Decode(&user)

		if err == nil {
			c.JSON(
				http.StatusBadRequest,
				models.Response{
					Status:  http.StatusBadRequest,
					Message: "Duplicate username",
					Detail: map[string]interface{}{
						"detail": "Duplicate Username: " + user.UserName,
					},
				},
			)
			return
		}

		user.Id = primitive.NewObjectID()
		user.Password = helpers.HashPassword(user.Password)

		res, err := database.UserCollection.InsertOne(ctx, user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Response{Status: http.StatusBadRequest, Message: "Error when insert", Detail: map[string]interface{}{"detail": err.Error()}})
		}

		c.JSON(http.StatusCreated, models.Response{Status: http.StatusCreated, Message: "Success", Detail: map[string]interface{}{"detail": res}})
	}
}
