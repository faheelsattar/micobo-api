package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//database config
var db *sql.DB

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "faheel"
	dbname   = "misobo"
)

func databaseConnection() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return db
}

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
	rows, err := db.Query(`select name, post from "Employees"`)
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

func goDotEnvVariable() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// getting env variables SITE_TITLE and DB_HOST
	host = os.Getenv("DB_HOST")
	convertedPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port = convertedPort
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname = os.Getenv("DB_NAME")
}

func main() {
	goDotEnvVariable()
	db = databaseConnection()
	router := gin.Default()

	//api routes
	router.GET("/employees", getEmployees)
	router.POST("/employess", addEmployees)

	//server running
	router.Run("localhost:8080")
}
