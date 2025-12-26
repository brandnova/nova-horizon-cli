package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	MaxFileSize = 100000 // 100KB
)

type ToolManager struct {
	workDir string
	verbose bool
}

func NewToolManager(workDir string, verbose bool) *ToolManager {
	return &ToolManager{
		workDir: workDir,
		verbose: verbose,
	}
}

func (tm *ToolManager) validatePath(filePath string) (string, error) {
	absWorkDir, err := filepath.Abs(tm.workDir)
	if err != nil {
		return "", fmt.Errorf("invalid working directory")
	}

	absPath, err := filepath.Abs(filepath.Join(tm.workDir, filePath))
	if err != nil {
		return "", fmt.Errorf("invalid file path")
	}

	// Ensure path is within working directory
	if !strings.HasPrefix(absPath, absWorkDir) {
		return "", fmt.Errorf("path traversal not allowed: %s is outside working directory", filePath)
	}

	return absPath, nil
}

// GetFilesInfo lists files in a directory
func (tm *ToolManager) GetFilesInfo(directory string) (string, error) {
	if directory == "" {
		directory = "."
	}

	absPath, err := tm.validatePath(directory)
	if err != nil {
		return "", err
	}

	entries, err := os.ReadDir(absPath)
	if err != nil {
		return "", fmt.Errorf("failed to read directory: %w", err)
	}

	var result strings.Builder
	for _, entry := range entries {
		info, _ := entry.Info()
		size := info.Size()
		isDir := entry.IsDir()
		result.WriteString(fmt.Sprintf("- %s: file_size=%d bytes, is_dir=%v\n", entry.Name(), size, isDir))
	}

	return result.String(), nil
}

// GetFileContent reads file contents
func (tm *ToolManager) GetFileContent(filePath string) (string, error) {
	absPath, err := tm.validatePath(filePath)
	if err != nil {
		return "", err
	}

	fileInfo, err := os.Stat(absPath)
	if err != nil {
		return "", fmt.Errorf("file not found: %w", err)
	}

	if fileInfo.IsDir() {
		return "", fmt.Errorf("cannot read directory as file")
	}

	if fileInfo.Size() > MaxFileSize {
		return "", fmt.Errorf("file too large (%d bytes, max %d)", fileInfo.Size(), MaxFileSize)
	}

	content, err := os.ReadFile(absPath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(content), nil
}

// WriteFile writes content to a file
func (tm *ToolManager) WriteFile(filePath string, content string) (string, error) {
	absPath, err := tm.validatePath(filePath)
	if err != nil {
		return "", err
	}

	// Check content size
	if len(content) > MaxFileSize {
		return "", fmt.Errorf("content too large (%d bytes, max %d)", len(content), MaxFileSize)
	}

	// Create parent directories if needed
	parentDir := filepath.Dir(absPath)
	if err := os.MkdirAll(parentDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directories: %w", err)
	}

	// Write file
	if err := os.WriteFile(absPath, []byte(content), 0644); err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	return fmt.Sprintf("File %s written successfully with %d characters", filePath, len(content)), nil
}
