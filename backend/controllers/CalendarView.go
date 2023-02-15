package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CalendarView() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
		return
	}
}
