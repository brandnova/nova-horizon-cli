package gemini

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiClient struct {
	client *genai.Client
	model  string
}

func NewGeminiClient(apiKey string, model string) (*GeminiClient, error) {
	cl, err := genai.NewClient(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %w", err)
	}

	return &GeminiClient{
		client: cl,
		model:  model,
	}, nil
}

func (gc *GeminiClient) GenerateContent(ctx context.Context, history []*genai.Content, tools []*genai.Tool) (*genai.GenerateContentResponse, error) {
	model := gc.client.GenerativeModel(gc.model)
	model.Tools = tools
	model.SetTemperature(0) // Default to deterministic

	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text(systemPrompt),
		},
	}

	cs := model.StartChat()

	// Separate history and the last message (which is the new input)
	if len(history) > 0 {
		cs.History = history[:len(history)-1]
		lastMsg := history[len(history)-1]
		return cs.SendMessage(ctx, lastMsg.Parts...)
	}

	// Should not happen if agent loop provides prompt
	return nil, fmt.Errorf("no messages provided")
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
