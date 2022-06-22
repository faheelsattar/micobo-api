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
var (
	host     = ""
	port     = 5432
	user     = ""
	password = ""
	dbname   = ""
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
	db, ok := c.MustGet("db").(*sql.DB)
	if !ok {
		c.IndentedJSON(http.StatusBadGateway, "db connection not found")
	}
	rows, err := db.Query(`SELECT "Name", "Roll" FROM "Students"`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var roll int

		err = rows.Scan(&name, &roll)
		if err != nil {
			panic(err)
		}
		fmt.Println(name, roll)
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

// ApiMiddleware will add the db connection to the context
func DatabaseMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	host = goDotEnvVariable("HOST")
	i, err := strconv.Atoi((goDotEnvVariable("PORT")))
	if err != nil {
		fmt.Println(err)
	}
	port = i
	user = goDotEnvVariable("USER")
	password = goDotEnvVariable("PASSWORD")
	dbname = goDotEnvVariable("DBNAME")

	db := databaseConnection()
	router := gin.Default()
	//api routes
	router.Use(DatabaseMiddleware(db))
	router.GET("/employees", getEmployees)
	router.POST("/employess", addEmployees)

	//server running
	router.Run("localhost:8080")
}
