package config

import "github.com/sashabaranov/go-openai"

var OpenAIClient *openai.Client

func InitOpenAIClient(apiKey string) {
	OpenAIClient = openai.NewClient(apiKey)
}