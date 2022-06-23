package controller

import (
	"fmt"
	"misobo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// employees data representation
type employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}

func employeeSanitization(emp employee) bool {
	return emp.ID > 0 && len(emp.Name) > 0 && len(emp.Gender) > 0 && len(emp.Birthday) > 0
}

// GetEmployees responds with the list of all employees as JSON.
func GetEmployees(c *gin.Context) {
	var employees = []employee{}

	rows, err := utils.DB.Query(`select id, name, gender, birthday from "Employees"`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var employeeData employee

		err = rows.Scan(&employeeData.ID, &employeeData.Name, &employeeData.Gender, &employeeData.Birthday)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		employees = append(employees, employeeData)
		fmt.Println(employeeData.Name, employeeData.Gender)
	}
	c.IndentedJSON(http.StatusOK, employees)
}

// AddEmployees adds a new employee in the database.
func AddEmployees(c *gin.Context) {
	var newEmployee employee

	if err := c.BindJSON(&newEmployee); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "body is invalid")
	}

	if !employeeSanitization(newEmployee) {
		c.IndentedJSON(http.StatusBadRequest, "body is invalid")
	}

	_, err := utils.DB.Exec(`insert into "Employees" (id, name, gender, birthday) values ($1, $2, $3, $4)`, newEmployee.ID, newEmployee.Name, newEmployee.Gender, newEmployee.Birthday)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, newEmployee)
}
