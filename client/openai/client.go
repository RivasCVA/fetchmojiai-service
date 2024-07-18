package openai

import (
	"context"
	"fmt"
	"os"

	openaiExternal "github.com/sashabaranov/go-openai"
)

type OpenAIClientInterface interface {
	// Sends a chat to GPT.
	Chat(prompt string) (string, error)
}

type OpenAIClient struct {
	openai *openaiExternal.Client
}

func NewClient() OpenAIClientInterface {
	// create a new external openai client
	token := os.Getenv("OPENAI_TOKEN")
	openai := openaiExternal.NewClient(token)

	return &OpenAIClient{
		openai: openai,
	}
}

func (c *OpenAIClient) Chat(prompt string) (string, error) {
	// send the prompt to gpt
	resp, err := c.openai.CreateChatCompletion(context.Background(), openaiExternal.ChatCompletionRequest{
		Model: openaiExternal.GPT3Dot5Turbo,
		Messages: []openaiExternal.ChatCompletionMessage{
			{
				Role:    openaiExternal.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("Chat: unable to create a chat completion: %w", err)
	}

	// return the message response
	content := resp.Choices[0].Message.Content
	return content, nil
}
