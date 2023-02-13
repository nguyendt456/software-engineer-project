package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyendt456/software-engineer-project/controllers"
)

func Routes(router *gin.Engine) {
	// router.Use(controllers.AuthMiddleware())
	router.POST("/registry", controllers.AddUser())
	router.POST("/login", controllers.LoginUser())
}
