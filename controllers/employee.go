package controller

import (
	"misobo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// employees data representation
type employee struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}

func employeeSanitization(emp employee) bool {
	return len(emp.ID) > 0 && len(emp.Name) > 0 && len(emp.Gender) > 0 && len(emp.Birthday) > 0
}

func GetEmployeeIds() ([]int, error) {
	var employeeIds = []int{}

	rows, err := utils.DB.Query(`select id from "Employees"`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var employeeID int

		err = rows.Scan(employeeID)
		employeeIds = append(employeeIds, employeeID)
	}

	return employeeIds, err
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
	}

	c.IndentedJSON(http.StatusOK, employees)
}

// AddEmployees adds a new employee in the database.
func AddEmployees(c *gin.Context) {
	var newEmployee employee

	if err := c.BindJSON(&newEmployee); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "body is invalid")
		return
	}

	if !employeeSanitization(newEmployee) {
		c.IndentedJSON(http.StatusBadRequest, "body is invalid")
		return
	}

	_, err := utils.DB.Exec(`insert into "Employees" (id, name, gender, birthday) values ($1, $2, $3, $4)`, newEmployee.ID, newEmployee.Name, newEmployee.Gender, newEmployee.Birthday)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, newEmployee)
}

// UpdateEmployees update an employee in the database.
func UpdateEmployee(c *gin.Context) {
	var newEmployee employee
	employeeId := c.Param("employee_id")

	exists := employeeExists(employeeId)

	if !exists {
		c.IndentedJSON(http.StatusBadRequest, "employee doesnt exist")
		return
	}
	if err := c.BindJSON(&newEmployee); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "body is invalid")
		return
	}

	if !employeeSanitization(newEmployee) {
		c.IndentedJSON(http.StatusBadRequest, "body is invalid")
		return
	}

	_, err := utils.DB.Exec(`update "Employees" set name=$2, gender=$3, birthday=$4 where id=$1`, employeeId, newEmployee.Name, newEmployee.Gender, newEmployee.Birthday)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, "updated employee successfully")
}

// DeleteEmployees deletes an employee from the database.
func DeleteEmployee(c *gin.Context) {
	employeeId := c.Param("employee_id")

	exists := employeeExists(employeeId)

	if !exists {
		c.IndentedJSON(http.StatusBadRequest, "employee doesnt exist")
		return
	}
	_, err := utils.DB.Exec(`delete from "Employees" where id=$1`, employeeId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "employee id wrong")
		return
	}
	c.IndentedJSON(http.StatusCreated, "deleted employee successfully")
}

func employeeExists(id string) bool {
	i := 0
	rows, err := utils.DB.Query(`select id from "Employees" where id=$1`, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		i++
	}
	return i == 1
}
