package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyendt456/software-engineer-project/models"
)

func CalendarView() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.MustGet("user").(models.User)
		if !ok {
			c.JSON(http.StatusInternalServerError, nil)
		}
		switch {
		case user.UserType == "backofficer":
			c.JSON(http.StatusAccepted, nil)
			break
		case user.UserType == "janitor" || user.UserType == "collector":
			c.JSON(http.StatusAccepted, nil)
			break
		}
	}
}
