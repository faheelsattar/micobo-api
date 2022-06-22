package controller

import (
	"fmt"
	"misobo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// employees data representation
type employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Post string `json:"post"`
}

func employeeSanitization(emp employee) bool {
	return len(emp.ID) > 0 && len(emp.Name) > 0 && len(emp.Post) > 0
}

// getEmployees responds with the list of all employees as JSON.
func GetEmployees(c *gin.Context) {
	rows, err := utils.DB.Query(`select name, post from "Employees"`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var post string

		err = rows.Scan(&name, &post)
		if err != nil {
			panic(err)
		}
		fmt.Println(name, post)
	}
	c.IndentedJSON(http.StatusOK, "ad")
}

func AddEmployees(c *gin.Context) {
	var newEmployee employee

	if err := c.BindJSON(&newEmployee); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "body is invalid")
	}

	if !employeeSanitization(newEmployee) {
		c.IndentedJSON(http.StatusBadRequest, "body is invalid")
	}

	// employees = append(employees, newEmployee)
	c.IndentedJSON(http.StatusCreated, newEmployee)
}
