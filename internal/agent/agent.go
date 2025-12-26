package agent

import (
	"context"
	"fmt"

	"github.com/brandnova/nova-horizon-cli/internal/gemini"
	"github.com/brandnova/nova-horizon-cli/internal/tools"
	"github.com/fatih/color"
	"github.com/google/generative-ai-go/genai"
)

type Config struct {
	APIKey    string
	Model     string
	WorkDir   string
	Verbose   bool
	DryRun    bool
	MaxSteps  int
	AllowRun  bool
	ApplyDiff bool
}

type Agent struct {
	config    *Config
	client    *gemini.GeminiClient
	toolMgr   *tools.ToolManager
	seenCalls map[string]bool
}

func NewAgent(cfg *Config) *Agent {
	return &Agent{
		config:    cfg,
		toolMgr:   tools.NewToolManager(cfg.WorkDir, cfg.Verbose),
		seenCalls: make(map[string]bool),
	}
}

func (a *Agent) Run(prompt string) error {
	var err error
	a.client, err = gemini.NewGeminiClient(a.config.APIKey, a.config.Model)
	if err != nil {
		return err
	}
	defer a.client.Close()

	ctx := context.Background()

	// Initialize messages with user prompt
	messages := []*genai.Content{
		{
			Role: "user",
			Parts: []genai.Part{
				genai.Text(prompt),
			},
		},
	}

	for step := 0; step < a.config.MaxSteps; step++ {
		if a.config.Verbose {
			fmt.Printf("[Step %d/%d]\n", step+1, a.config.MaxSteps)
		}

		// Build tools
		toolDefs := gemini.BuildTools()

		// Call Gemini API
		resp, err := a.client.GenerateContent(ctx, messages, toolDefs)
		if err != nil {
			return fmt.Errorf("API call failed: %w", err)
		}

		if resp == nil || len(resp.Candidates) == 0 {
			return fmt.Errorf("empty response from API")
		}

		candidate := resp.Candidates[0]
		if candidate.Content == nil {
			return fmt.Errorf("malformed response")
		}

		// Add response to messages
		messages = append(messages, candidate.Content)

		// Check for function calls
		hasFunctionCall := false
		var functionResponses []genai.Part

		for _, part := range candidate.Content.Parts {
			if fc, ok := part.(genai.FunctionCall); ok {
				hasFunctionCall = true

				// Check for loops
				callSignature := fmt.Sprintf("%s:%v", fc.Name, fc.Args)
				if a.seenCalls[callSignature] {
					color.Yellow("Model is looping on the same function call. Aborting.")
					return nil
				}
				a.seenCalls[callSignature] = true

				// Execute function
				result, err := a.executeFunction(fc)
				if err != nil {
					color.Red("Error executing %s: %v", fc.Name, err)
					result = fmt.Sprintf("Error: %v", err)
				}

				if a.config.Verbose {
					fmt.Printf(" - Called: %s\n", fc.Name)
				} else {
					fmt.Printf(" - Calling function: %s\n", fc.Name)
				}

				// Add response part
				functionResponses = append(functionResponses, genai.FunctionResponse{
					Name:     fc.Name,
					Response: map[string]interface{}{"result": result},
				})
			} else if text, ok := part.(genai.Text); ok {
				// Final response
				if string(text) != "" {
					fmt.Println(string(text))
				}
			}
		}

		// If we have function responses, add them to history
		if len(functionResponses) > 0 {
			messages = append(messages, &genai.Content{
				Role: "function", // genai uses 'function' role for tool outputs? Or 'tool'?
				// genai.Client usually expects 'function' or 'user' depending on API version but docs say 'function' for responses
				// Wait, the SDK constants say RoleFunction = "function"
				Parts: functionResponses,
			})
		}

		// If no function calls, we're done
		if !hasFunctionCall {
			return nil
		}
	}

	color.Yellow("Reached maximum steps (%d)", a.config.MaxSteps)
	return nil
}

func (a *Agent) executeFunction(fc genai.FunctionCall) (string, error) {
	switch fc.Name {
	case "get_files_info":
		dir, _ := fc.Args["directory"].(string)
		if dir == "" {
			dir = "."
		}
		return a.toolMgr.GetFilesInfo(dir)

	case "get_file_content":
		filePath, ok := fc.Args["file_path"].(string)
		if !ok {
			return "", fmt.Errorf("missing file_path argument")
		}
		return a.toolMgr.GetFileContent(filePath)

	case "write_file":
		filePath, ok := fc.Args["file_path"].(string)
		if !ok {
			return "", fmt.Errorf("missing file_path argument")
		}
		content, ok := fc.Args["content"].(string)
		if !ok {
			return "", fmt.Errorf("missing content argument")
		}

		if a.config.DryRun {
			return fmt.Sprintf("[DRY RUN] Would write %d bytes to %s", len(content), filePath), nil
		}

		return a.toolMgr.WriteFile(filePath, content)

	case "run_file":
		filePath, ok := fc.Args["file_path"].(string)
		if !ok {
			return "", fmt.Errorf("missing file_path argument")
		}

		if !a.config.AllowRun {
			return "", fmt.Errorf("program execution not allowed (use --allow-run flag)")
		}

		args := []string{}
		if argsVal, ok := fc.Args["args"].([]interface{}); ok {
			for _, arg := range argsVal {
				if s, ok := arg.(string); ok {
					args = append(args, s)
				}
			}
		}

		return a.toolMgr.RunFile(filePath, args)

	default:
		return "", fmt.Errorf("unknown function: %s", fc.Name)
	}
}
