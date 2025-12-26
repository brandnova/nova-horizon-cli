package gemini

import (
	"github.com/google/generative-ai-go/genai"
)

// BuildTools creates tool definitions for the Gemini API
func BuildTools() []*genai.Tool {
	return []*genai.Tool{
		{
			FunctionDeclarations: []*genai.FunctionDeclaration{
				getFilesInfoSchema(),
				getFileContentSchema(),
				writeFileSchema(),
				runFileSchema(),
			},
		},
	}
}

func getFilesInfoSchema() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "get_files_info",
		Description: "Lists files in a specified directory relative to the working directory, providing file size and directory status",
		Parameters: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"directory": {
					Type:        genai.TypeString,
					Description: "Directory path to list files from, relative to the working directory (default is the working directory itself)",
				},
			},
		},
	}
}

func getFileContentSchema() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "get_file_content",
		Description: "Retrieves the content of a specified file relative to the working directory",
		Parameters: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"file_path": {
					Type:        genai.TypeString,
					Description: "Path of the file to read, relative to the working directory",
				},
			},
			Required: []string{"file_path"},
		},
	}
}

func writeFileSchema() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "write_file",
		Description: "Writes content to a specified file or creates a new file relative to the working directory. Creates directories if they do not exist.",
		Parameters: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"file_path": {
					Type:        genai.TypeString,
					Description: "Path of the file to write, relative to the working directory",
				},
				"content": {
					Type:        genai.TypeString,
					Description: "Content to write to the file as a string",
				},
			},
			Required: []string{"file_path", "content"},
		},
	}
}

func runFileSchema() *genai.FunctionDeclaration {
	return &genai.FunctionDeclaration{
		Name:        "run_file",
		Description: "Executes a specified file relative to the working directory (.go, .py, .sh, .js, .ts supported), with optional CLI args",
		Parameters: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"file_path": {
					Type:        genai.TypeString,
					Description: "Path of the file to execute, relative to the working directory",
				},
				"args": {
					Type: genai.TypeArray,
					Items: &genai.Schema{
						Type: genai.TypeString,
					},
					Description: "Optional array of string arguments to pass to the file",
				},
			},
			Required: []string{"file_path"},
		},
	}
}
