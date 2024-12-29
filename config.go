package main

import (
	"os"
)

// GetDatabaseURL retrieves the database connection string from the environment
func GetDatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}
