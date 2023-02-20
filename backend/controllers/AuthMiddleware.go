package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyendt456/software-engineer-project/helpers"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")

		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "No authenticate token in headers",
			})
			c.Abort()
			return
		}
		userSignedDetail, res, err := helpers.ValidateAuthToken(clientToken)
		fmt.Println(userSignedDetail)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
				"token": clientToken,
			})
			c.Abort()
			return
		}
		if !res {
			log.Print("Invalid Token")
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid Token",
			})
			c.Abort()
			return
		}
		c.Set("user", userSignedDetail)
		return
	}
}
