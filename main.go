package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/roshan-ican/jacketstore-backend/config"
    "github.com/roshan-ican/jacketstore-backend/routes"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	// Health Check Route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Drip Jacket Store backend is healthy! 🚀",
		})
	})
 // User Routes
 routes.UserRoutes(r)
	log.Println("🚀 App running on http://localhost:8080")
	r.Run() // Default listens on :808
}
