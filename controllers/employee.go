package controller

import (
	"fmt"
	"misobo/entities"
	"misobo/psql"
	"misobo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func employeeSanitization(employee *entities.Employee) bool {
	return len(employee.ID) > 0 && len(employee.Name) > 0 && len(employee.Gender) > 0 && len(employee.Birthday) > 0
}

// GetEmployees responds with the list of all employees as JSON.
func GetEmployees(c *gin.Context) {
	repo := &psql.Repository{Db: utils.DB}

	employees, err := repo.FindEmployees()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, employees)
}

// AddEmployees adds a new employee in the database.
func AddEmployees(c *gin.Context) {
	var newEmployee entities.Employee

	if err := c.BindJSON(&newEmployee); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "body is invalid")
		return
	}

	if !employeeSanitization(&newEmployee) {
		c.IndentedJSON(http.StatusBadRequest, "body is invalid")
		return
	}

	repo := &psql.Repository{Db: utils.DB}

	err := repo.CreateEmployee(&newEmployee)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, newEmployee)
}

// UpdateEmployees update an employee in the database.
func UpdateEmployee(c *gin.Context) {
	var newEmployee entities.Employee
	employeeId := c.Param("employee_id")
	repo := &psql.Repository{Db: utils.DB}

	exists := repo.EmployeeExists(employeeId)
	fmt.Println("exists!!!", exists)
	if !exists {
		c.IndentedJSON(http.StatusBadRequest, "employee doesnt exist")
		return
	}
	if err := c.BindJSON(&newEmployee); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "body is invalid")
		return
	}

	if !employeeSanitization(&newEmployee) {
		c.IndentedJSON(http.StatusBadRequest, "body is invalid")
		return
	}

	err := repo.UpdateEmployee(&newEmployee, employeeId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, "updated employee successfully")
}

// DeleteEmployees deletes an employee from the database.
func DeleteEmployee(c *gin.Context) {
	employeeId := c.Param("employee_id")
	repo := &psql.Repository{Db: utils.DB}

	exists := repo.EmployeeExists(employeeId)

	if !exists {
		c.IndentedJSON(http.StatusBadRequest, "employee doesnt exist")
		return
	}

	err := repo.DeleteEmployee(employeeId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "employee id wrong")
		return
	}
	c.IndentedJSON(http.StatusCreated, "deleted employee successfully")
}
