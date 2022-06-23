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
	router.PUT("/employees/:employee_id", controller.UpdateEmployee)
	router.DELETE("/employees/:employee_id", controller.DeleteEmployee)

	router.GET("/events", controller.GetEvents)
	router.GET("/events/event_id", controller.GetEvent)

	//server running
	router.Run("localhost:8080")
}
