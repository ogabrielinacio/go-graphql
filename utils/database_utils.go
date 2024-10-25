package utils

import (
	"fmt"
	"os"
)

func GetConnectionString() string {
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s  sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
}
