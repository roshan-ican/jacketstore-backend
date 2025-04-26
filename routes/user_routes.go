package routes

import (
    "github.com/gin-gonic/gin"
	"github.com/roshan-ican/jacketstore-backend/controllers"
)

func UserRoutes(router *gin.Engine) {
    userGroup := router.Group("/users")
    {
        userGroup.POST("/register", controllers.RegisterUser)
		userGroup.POST("/login", controllers.LoginUser) 
    }
}