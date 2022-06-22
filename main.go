package main

import (
	controller "misobo/controllers"
	"misobo/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	utils.DatabaseConnection()
	router := gin.Default()

	//api routes
	router.GET("/employees", controller.GetEmployees)
	router.POST("/employees", controller.AddEmployees)

	//server running
	router.Run("localhost:8080")
}
