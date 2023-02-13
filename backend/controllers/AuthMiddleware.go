package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyendt456/software-engineer-project/helpers"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")

		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "No authenticate token in headers",
			})
			c.Abort()
			return
		}
		res, err := helpers.ValidateAuthToken(clientToken)
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
		}
	}
}
