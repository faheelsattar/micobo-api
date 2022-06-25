package psql

import (
	"database/sql"
	"misobo/entities"
)

// repository represent the repository model
type Repository struct {
	Db *sql.DB
}

// Close attaches the provider and close the connection
func (r *Repository) Close() {
	r.Db.Close()
}

func (repo *Repository) FindEmployees() ([]entities.Employee, error) {
	var employees = []entities.Employee{}

	rows, err := repo.Db.Query(`select id, name, gender, birthday from "Employees"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var newEmployee entities.Employee
		err = rows.Scan(
			&newEmployee.ID,
			&newEmployee.Name,
			&newEmployee.Gender,
			&newEmployee.Birthday,
		)

		if err != nil {
			return nil, err
		}
		employees = append(employees, newEmployee)
	}

	return employees, nil
}

func (repo *Repository) FindEmployeeIds() ([]int, error) {
	var employeeIds = []int{}

	rows, err := repo.Db.Query(`select id from "Employees"`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var employeeId int

		err = rows.Scan(employeeId)
		employeeIds = append(employeeIds, employeeId)
	}

	return employeeIds, err
}

// Checks if an employee exists in database
func (repo *Repository) EmployeeExists(employeeId string) bool {
	i := 0
	rows, err := repo.Db.Query(`select id from "Employees" where id=$1`, employeeId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		i++
	}
	return i == 1
}

// Create attaches the employee repository and creating the data
func (repo *Repository) CreateEmployee(employee *entities.Employee) error {
	_, err := repo.Db.Exec(`insert into "Employees" (id, name, gender, birthday) values ($1, $2, $3, $4)`, employee.ID, employee.Name, employee.Gender, employee.Birthday)
	if err != nil {
		return err
	}
	return err
}

// Update attaches the employee repository and update data based on id
func (repo *Repository) UpdateEmployee(employee *entities.Employee, employeeId string) error {
	_, err := repo.Db.Exec(`update "Employees" set name=$2, gender=$3, birthday=$4 where id=$1`, employeeId, employee.Name, employee.Gender, employee.Birthday)
	if err != nil {
		return err
	}
	return err
}

func (repo *Repository) DeleteEmployee(employeeId string) error {
	_, err := repo.Db.Exec(`delete from "Employees" where id=$1`, employeeId)
	if err != nil {
		return err
	}
	return err
}
