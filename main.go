package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// employees data representation
type employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Post string `json:"post"`
}

// employee in memory data
var employees = []employee{
	{ID: "1", Name: "Frank", Post: "Product manager"},
	{ID: "2", Name: "James", Post: "Solidity Developer"},
}

// getEmployees responds with the list of all employees as JSON.
func getEmployees(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, employees)
}

func main() {
	router := gin.Default()
	router.GET("/employees", getEmployees)

	router.Run("localhost:8080")
}
