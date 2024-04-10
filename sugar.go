package main

import (
	_ "embed"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

//go:embed .env
var embeddedEnvContent string

func loadDotenv() {
	// Try to load dotenv from file in directory
	err := godotenv.Load()

	// If dotenv file is not found in directory, use embedded content
	if err != nil {
		lines := strings.Split(embeddedEnvContent, "\n")
		for _, line := range lines {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := parts[0]
				value := parts[1]
				os.Setenv(key, value)
			}
		}
	}
}

func makeSentenceMoreFriendly(sentence string, apiKey string) (string, error) {
	// Create client with API key
	client := openai.NewClient(apiKey)

	// Call the function that returns the chat completion
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Make the user-provided sentence less blunt, and a bit more friendly, but without exaggeration",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: sentence,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return strings.TrimSpace(response.Choices[0].Message.Content), nil
}

func main() {

	// Parse sentence to sweeten
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: ./sugar <blunt_sentence>")
		return
	}
	bluntSentence := strings.Join(args, " ")

	// Load OpenAI API key
	loadDotenv()
	apiKey := os.Getenv("OPENAI_API_KEY")

	// Sweeten the sentence with GPT
	friendlySentence, err := makeSentenceMoreFriendly(bluntSentence, apiKey)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// The only program output is the rewritten sentence
	fmt.Println(friendlySentence)
}
