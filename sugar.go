package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func makeSentenceMoreFriendly(sentence string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

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
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: ./friendly_app <blunt_sentence>")
		return
	}

	bluntSentence := strings.Join(args, " ")
	
	friendlySentence, err := makeSentenceMoreFriendly(bluntSentence)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(friendlySentence)
}
