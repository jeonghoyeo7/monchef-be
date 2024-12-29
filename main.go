package main

import (
	"monchef/db"
	"monchef/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	db.ConnectDatabase()

	// Initialize the Gin router
	router := gin.Default()

	// Register routes
	routes.InitializeRoutes(router)

	// Start the server
	router.Run(":8080")
}
