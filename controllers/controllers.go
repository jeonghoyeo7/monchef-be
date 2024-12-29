package controllers

import (
	"monchef/db"
	"monchef/models"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// GetRecipes retrieves all recipes
func GetRecipes(c *gin.Context) {
	var recipes []models.Recipe

	// 데이터베이스에서 모든 레시피 조회
	result := db.DB.Find(&recipes)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// 결과 반환
	c.JSON(http.StatusOK, recipes)
}

// CreateRecipe adds a new recipe
func CreateRecipe(c *gin.Context) {
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := db.DB.Create(&recipe)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, recipe)
}

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload image"})
		return
	}

	// Save file to a local directory
	savePath := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	// Return the file path
	c.JSON(http.StatusOK, gin.H{"file_path": savePath})
}
