package main

import (
	"monchef/db"
	"monchef/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 데이터베이스 연결
	db.ConnectDatabase()

	// 라우터 초기화
	router := gin.Default()

	// 기본 API
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to MonChef API",
			"status":  "success",
		})
	})

	// /recipes 경로: 데이터베이스에서 모든 레시피 반환
	router.GET("/recipes", func(c *gin.Context) {
		var recipes []models.Recipe

		// 데이터베이스에서 모든 레시피 가져오기
		result := db.DB.Find(&recipes)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, recipes)
	})

	// 서버 실행
	router.Run(":8080")
}
