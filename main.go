package main

import (
	"gin-mongo-api/configs"
	"gin-mongo-api/controllers"
	
	"gin-mongo-api/routes" //add this

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()

	controllers.ServiceAccount("client_secret.json")

	//routes
	routes.UserRoute(router) //add this

	router.Run("localhost:6000")
}
