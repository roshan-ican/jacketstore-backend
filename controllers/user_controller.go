package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/roshan-ican/jacketstore-backend/config"
	"github.com/roshan-ican/jacketstore-backend/models"
	openai "github.com/sashabaranov/go-openai"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func RegisterUser(c *gin.Context) {
	var registerData struct {
		FullName string `json:"full_name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Phone    string `json:"phone"`
		Address  string `json:"address"`
	}

	if err := c.BindJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userCollection := config.GetCollection("jacketstore", "users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var existingUser models.User
	err := userCollection.FindOne(ctx, bson.M{"email": registerData.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with this email already exists"})
		return
	}

	hashedPassword, _ := HashPassword(registerData.Password)

	newUser := models.User{
		ID:           primitive.NewObjectID(),
		FullName:     registerData.FullName,
		Email:        registerData.Email,
		PasswordHash: hashedPassword,
		Phone:        registerData.Phone,
		Address:      registerData.Address,
		CreatedAt:    primitive.NewDateTimeFromTime(time.Now()),
	}

	_, err = userCollection.InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully!",
		"user_id": newUser.ID.Hex(),
	})
}

func LoginUser(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Received Login Request:", loginData)

	userCollection := config.GetCollection("jacketstore", "users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"email": loginData.Email}).Decode(&user)
	if err != nil {
		fmt.Println("User not found for email:", loginData.Email)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	fmt.Println("Fetched User from DB:", user)

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginData.Password))
	if err != nil {
		// Passwords do not match
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	fmt.Println("Login successful for user:", user.Email)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful!",
		"user_id": user.ID.Hex(),
	})
}

func PostDallePrompt(c *gin.Context) {
	fmt.Println("🔥 POST /dalle hit!")
	var requestBody struct {
		Prompt string `json:"prompt"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	resp, err := config.OpenAIClient.CreateImage(ctx, openai.ImageRequest{
		Prompt:         requestBody.Prompt,
		N:              1,
		Size:           "512x512",
		ResponseFormat: "url",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "OpenAI API error", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"photo": resp.Data[0].URL,
	})
}