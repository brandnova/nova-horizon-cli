package tools

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"time"
)

var allowedExtensions = []string{".go", ".py", ".sh", ".js", ".ts"}

// RunFile executes a file
func (tm *ToolManager) RunFile(filePath string, args []string) (string, error) {
	absPath, err := tm.validatePath(filePath)
	if err != nil {
		return "", err
	}

	// Check file extension
	ext := filepath.Ext(absPath)
	allowed := false
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			allowed = true
			break
		}
	}
	if !allowed {
		return "", fmt.Errorf("file type not allowed: %s (allowed: %v)", ext, allowedExtensions)
	}

	// Determine command based on extension
	var cmd *exec.Cmd
	switch ext {
	case ".go":
		cmd = exec.Command("go", append([]string{"run", absPath}, args...)...)
	case ".py":
		cmd = exec.Command("python3", append([]string{absPath}, args...)...)
	case ".sh":
		cmd = exec.Command("bash", append([]string{absPath}, args...)...)
	case ".js":
		cmd = exec.Command("node", append([]string{absPath}, args...)...)
	case ".ts":
		cmd = exec.Command("ts-node", append([]string{absPath}, args...)...)
	}

	// Set working directory
	cmd.Dir = tm.workDir

	// Capture output with timeout
	done := make(chan error, 1)
	output := make(chan string, 1)

	go func() {
		stdOut, _ := cmd.Output()
		output <- string(stdOut)
		done <- cmd.Run()
	}()

	select {
	case <-time.After(30 * time.Second):
		cmd.Process.Kill()
		return "", fmt.Errorf("execution timeout (30 seconds)")
	case err := <-done:
		if err != nil {
			return "", fmt.Errorf("execution failed: %w", err)
		}
		return <-output, nil
	}
}
