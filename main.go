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

func employeeSanitization(emp employee) bool {
	return len(emp.ID) > 0 && len(emp.Name) > 0 && len(emp.Post) > 0
}

// getEmployees responds with the list of all employees as JSON.
func getEmployees(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, employees)
}

func addEmployees(c *gin.Context) {
	var newEmployee employee

	if err := c.BindJSON(&newEmployee); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "body is invalid")
	}

	if !employeeSanitization(newEmployee) {
		c.IndentedJSON(http.StatusBadRequest, "body is invalid")
	}

	employees = append(employees, newEmployee)
	c.IndentedJSON(http.StatusCreated, newEmployee)
}

func main() {
	router := gin.Default()
	//api routes
	router.GET("/employees", getEmployees)
	router.POST("/employess", addEmployees)

	//server running
	router.Run("localhost:8080")
}
