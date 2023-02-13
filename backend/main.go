package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nguyendt456/software-engineer-project/routes"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	routes.Routes(router)

	return router
}

func main() {
	router := SetupRouter()

	router.Run()
}
