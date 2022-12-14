package main

import (
	"gin-mongo-boilerplate/configs"

	"gin-mongo-boilerplate/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	routes.UserRoute(router)

	router.Run("localhost:6000")
}
