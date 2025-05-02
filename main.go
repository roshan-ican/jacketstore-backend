package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/roshan-ican/jacketstore-backend/config"
	"github.com/roshan-ican/jacketstore-backend/routes"
)

func main() {
	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// Initialize MongoDB
	config.ConnectDB()

	// Initialize OpenAI Client with API key from env
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("❌ OPENAI_API_KEY not found in environment")
	}
	config.InitOpenAIClient(apiKey)

	// Initialize router
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback
	}

	log.Printf("🚀 App running on http://localhost:%s", port)
	r.Run(":" + port)
}