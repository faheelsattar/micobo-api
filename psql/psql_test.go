package psql

import (
	"database/sql"
	"log"
	"misobo/entities"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var employee = &entities.Employee{
	ID:       "2",
	Name:     "Devin",
	Gender:   "male",
	Birthday: "1992-12-10",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestFindEmployees(t *testing.T) {
	db, mock := NewMock()
	repo := &Repository{db}
	defer func() {
		repo.Close()
	}()

	query := `select id, name, gender, birthday from "Employees"`

	rows := sqlmock.NewRows([]string{"id", "name", "gender", "birthday"}).
		AddRow(employee.ID, employee.Name, employee.Gender, employee.Birthday)

	mock.ExpectQuery(query).WillReturnRows(rows)

	employee, err := repo.FindEmployees()
	assert.NotEmpty(t, employee)
	assert.NoError(t, err)
	assert.Len(t, employee, 1)
}

func TestCreate(t *testing.T) {
	db, mock := NewMock()
	repo := &Repository{db}
	defer func() {
		repo.Close()
	}()

	query := "insert into Employees \\(id, name, gender, birthday\\) values \\(\\$1, \\$2, \\$3, \\$4\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(employee.ID, employee.Name, employee.Gender, employee.Birthday).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.CreateEmployee(employee)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	db, mock := NewMock()
	repo := &Repository{db}
	defer func() {
		repo.Close()
	}()

	query := "update Employees set name = \\$1, gender = \\$2, birthday = \\$3 where id = \\$4"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(employee.Name, employee.Gender, employee.Birthday, employee.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.UpdateEmployee(employee, employee.ID)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	db, mock := NewMock()
	repo := &Repository{db}
	defer func() {
		repo.Close()
	}()

	query := "delete from Employees where id = \\$1"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(employee.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.DeleteEmployee(employee.ID)
	assert.NoError(t, err)
}
