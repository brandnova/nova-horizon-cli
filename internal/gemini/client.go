package gemini

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/client"
	"github.com/google/generative-ai-go/types"
)

type GeminiClient struct {
	client *client.Client
	model  string
}

type Message struct {
	Role  string
	Parts []interface{}
}

type FunctionCall struct {
	Name string
	Args map[string]interface{}
}

func NewGeminiClient(apiKey string, model string) (*GeminiClient, error) {
	cl, err := client.NewClient(context.Background(), &client.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %w", err)
	}

	return &GeminiClient{
		client: cl,
		model:  model,
	}, nil
}

func (gc *GeminiClient) GenerateContent(ctx context.Context, messages []interface{}, tools []*types.Tool) (*types.GenerateContentResponse, error) {
	content := &types.Content{
		Role: "user",
		Parts: []types.Part{
			{Text: messages[len(messages)-1].(string)},
		},
	}

	config := &types.GenerateContentConfig{
		Tools: tools,
		SystemInstruction: &types.Content{
			Parts: []types.Part{
				{Text: systemPrompt},
			},
		},
	}

	resp, err := gc.client.Models.GenerateContent(ctx, gc.model, []*types.Content{content}, config)
	if err != nil {
		return nil, fmt.Errorf("API call failed: %w", err)
	}

	return resp, nil
}

const systemPrompt = `You are a helpful AI coding agent.

When a user asks a question or makes a request, make a function call plan. You can perform the following operations:

- List files and directories
- Read file contents
- Write or modify files
- Execute scripts and programs

All paths you provide should be relative to the working directory. You do not need to specify the working directory in your function calls as it is automatically injected for security reasons.

Follow these guidelines:
1. Make function calls to gather information first
2. Plan your approach before making changes
3. Provide clear feedback about what you're doing
4. Show diffs before writing files
5. Stop if you detect infinite loops (same call multiple times)`

func (gc *GeminiClient) Close() error {
	return gc.client.Close()
}
