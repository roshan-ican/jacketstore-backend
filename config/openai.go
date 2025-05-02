package config

import (
	"os"

	openai "github.com/sashabaranov/go-openai"
)

var OpenAIClient *openai.Client

func InitOpenAI() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	OpenAIClient = openai.NewClient(apiKey)
}