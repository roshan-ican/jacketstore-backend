package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/roshan-ican/jacketstore-backend/controllers"
)

func UserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/dalle", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "GET dalle working"})
		})
		userGroup.POST("/register", controllers.RegisterUser)
		userGroup.POST("/login", controllers.LoginUser)
		userGroup.POST("/dalle", controllers.PostDallePrompt)
	}
}