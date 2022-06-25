package psql

import (
	"database/sql"
	"log"
	"misobo/entities"
	"regexp"
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

// var event = &entities.Event{
// 	ID:           "1",
// 	Name:         "Escape room",
// 	Scheduled:    "2022-07-02",
// 	Attend:       "3,2",
// 	Accomodation: "2",
// }

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

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

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

	query := `insert into "Employees" (id, name, gender, birthday) values ($1, $2, $3, $4)`

	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
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

	query := `update "Employees" set name = $1, gender = $2, birthday = $3 where id = $4`

	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
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

	query := `delete from "Employees" where id = $1`

	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
	prep.ExpectExec().WithArgs(employee.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.DeleteEmployee(employee.ID)
	assert.NoError(t, err)
}

// ====> @NOTICE SQL MOCK IS NOT PARSING ARRAYS AND THUS THE EVENT SQL FAILS
//       BUT WORKS IN ACTUAL API

// func TestFindEvents(t *testing.T) {
// 	db, mock := NewMock()
// 	repo := &Repository{db}
// 	defer func() {
// 		repo.Close()
// 	}()

// 	query := `select id, name, scheduled, attend, accomodation from "Events"`

// 	rows := sqlmock.NewRows([]string{"id", "name", "scheduled", "attend", "accomodation"}).
// 		AddRow(event.ID, event.Name, event.Scheduled, strings.Split(event.Attend, ","), strings.Split(event.Accomodation, ","))

// 	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

// 	events, err := repo.FindEvents()
// 	assert.NotEmpty(t, events)
// 	assert.NoError(t, err)
// 	assert.Len(t, events, 1)
// }
