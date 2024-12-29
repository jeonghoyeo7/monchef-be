package routes

import (
	"monchef/controllers"

	"github.com/gin-gonic/gin"
)

// InitializeRoutes sets up all API routes
func InitializeRoutes(router *gin.Engine) {
	router.GET("/recipes", controllers.GetRecipes)
	router.POST("/recipes", controllers.CreateRecipe)
}
