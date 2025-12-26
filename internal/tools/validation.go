package tools

import (
	"fmt"
	"path/filepath"
	"strings"
)

// ValidateFileExtension checks if file extension is allowed for writing
func ValidateFileExtension(filePath string) error {
	allowedExts := []string{".go", ".py", ".sh", ".js", ".ts", ".md", ".txt", ".json", ".yaml", ".yml", ".toml", ".env"}

	ext := filepath.Ext(filePath)
	for _, allowed := range allowedExts {
		if ext == allowed {
			return nil
		}
	}

	return fmt.Errorf("file extension %s not allowed for writing (allowed: %v)", ext, allowedExts)
}

// SanitizePath prevents directory traversal attacks
func SanitizePath(basePath, relativePath string) (string, error) {
	fullPath := filepath.Join(basePath, relativePath)
	fullPath = filepath.Clean(fullPath)

	if !strings.HasPrefix(fullPath, filepath.Clean(basePath)) {
		return "", fmt.Errorf("path traversal detected")
	}

	return fullPath, nil
}
