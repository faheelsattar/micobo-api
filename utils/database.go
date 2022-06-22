package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func DatabaseConnection() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// getting env variables SITE_TITLE and DB_HOST
	host := os.Getenv("DB_HOST")
	convertedPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := convertedPort
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	dbInstance, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	DB = dbInstance
}
