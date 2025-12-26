package gemini

import (
	"github.com/google/generative-ai-go/types"
)

// BuildTools creates tool definitions for the Gemini API
func BuildTools() []*types.Tool {
	return []*types.Tool{
		{
			FunctionDeclarations: []*types.FunctionDeclaration{
				getFilesInfoSchema(),
				getFileContentSchema(),
				writeFileSchema(),
				runFileSchema(),
			},
		},
	}
}

func getFilesInfoSchema() *types.FunctionDeclaration {
	return &types.FunctionDeclaration{
		Name:        "get_files_info",
		Description: "Lists files in a specified directory relative to the working directory, providing file size and directory status",
		Parameters: &types.Schema{
			Type: types.TypeObject,
			Properties: map[string]*types.Schema{
				"directory": {
					Type:        types.TypeString,
					Description: "Directory path to list files from, relative to the working directory (default is the working directory itself)",
				},
			},
		},
	}
}

func getFileContentSchema() *types.FunctionDeclaration {
	return &types.FunctionDeclaration{
		Name:        "get_file_content",
		Description: "Retrieves the content of a specified file relative to the working directory",
		Parameters: &types.Schema{
			Type: types.TypeObject,
			Properties: map[string]*types.Schema{
				"file_path": {
					Type:        types.TypeString,
					Description: "Path of the file to read, relative to the working directory",
				},
			},
			Required: []string{"file_path"},
		},
	}
}

func writeFileSchema() *types.FunctionDeclaration {
	return &types.FunctionDeclaration{
		Name:        "write_file",
		Description: "Writes content to a specified file or creates a new file relative to the working directory. Creates directories if they do not exist.",
		Parameters: &types.Schema{
			Type: types.TypeObject,
			Properties: map[string]*types.Schema{
				"file_path": {
					Type:        types.TypeString,
					Description: "Path of the file to write, relative to the working directory",
				},
				"content": {
					Type:        types.TypeString,
					Description: "Content to write to the file as a string",
				},
			},
			Required: []string{"file_path", "content"},
		},
	}
}

func runFileSchema() *types.FunctionDeclaration {
	return &types.FunctionDeclaration{
		Name:        "run_file",
		Description: "Executes a specified file relative to the working directory (.go, .py, .sh, .js, .ts supported), with optional CLI args",
		Parameters: &types.Schema{
			Type: types.TypeObject,
			Properties: map[string]*types.Schema{
				"file_path": {
					Type:        types.TypeString,
					Description: "Path of the file to execute, relative to the working directory",
				},
				"args": {
					Type: types.TypeArray,
					Items: &types.Schema{
						Type: types.TypeString,
					},
					Description: "Optional array of string arguments to pass to the file",
				},
			},
			Required: []string{"file_path"},
		},
	}
}
