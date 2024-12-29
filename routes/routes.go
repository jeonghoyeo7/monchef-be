package routes

import (
	"monchef/controllers"

	"github.com/gin-gonic/gin"
)

// InitializeRoutes sets up all API routes
func InitializeRoutes(router *gin.Engine) {
	// Base endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to MonChef API",
			"status":  "success",
		})
	})

	// Recipe endpoints
	router.GET("/recipes", controllers.GetRecipes)    // Retrieve all recipes
	router.POST("/recipes", controllers.CreateRecipe) // Add a new recipe

	// Image upload endpoint
	router.POST("/upload", controllers.UploadImage) // Upload an image
}
