package psql

import (
	"database/sql"
	"log"
	"misobo/utils"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestFind(t *testing.T) {
	utils.DatabaseConnection()

	_, mock := NewMock()

	query := `select id, name, gender, birthday from "Employees"`

	rows := sqlmock.NewRows([]string{"id", "name", "gender", "birthday"}).
		AddRow("2", "Devin", "male", "1992-12-10").AddRow("3", "Levin", "female", "1998-01-01")

	mock.ExpectQuery(query).WillReturnRows(rows)

	users, err := Find()
	assert.NotEmpty(t, users)
	assert.NoError(t, err)
	assert.Len(t, users, 2)
}
