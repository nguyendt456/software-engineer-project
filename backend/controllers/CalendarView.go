package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CalendarView() gin.HandlerFunc {
	return func(c *gin.Context) {
		// switch userType := c
		c.JSON(http.StatusOK, "")
		return
	}
}
