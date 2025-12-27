package tools

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// GenerateDiff creates a unified diff between old and new content
func (tm *ToolManager) GenerateDiff(oldContent, newContent string) string {
	oldLines := strings.Split(oldContent, "\n")
	newLines := strings.Split(newContent, "\n")

	var diff strings.Builder
	diff.WriteString("--- original\n")
	diff.WriteString("+++ modified\n")

	// Simple line-by-line diff (can be enhanced with proper diff algorithm)
	maxLen := len(oldLines)
	if len(newLines) > maxLen {
		maxLen = len(newLines)
	}

	for i := 0; i < maxLen; i++ {
		oldLine := ""
		newLine := ""

		if i < len(oldLines) {
			oldLine = oldLines[i]
		}
		if i < len(newLines) {
			newLine = newLines[i]
		}

		if oldLine != newLine {
			if oldLine != "" {
				diff.WriteString(fmt.Sprintf("-%s\n", oldLine))
			}
			if newLine != "" {
				diff.WriteString(fmt.Sprintf("+%s\n", newLine))
			}
		} else if oldLine != "" {
			diff.WriteString(fmt.Sprintf(" %s\n", oldLine))
		}
	}

	return diff.String()
}

// PrintColoredDiff prints a diff with color coding
func PrintColoredDiff(diff string) {
	lines := strings.Split(diff, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++") {
			color.Green(line)
		} else if strings.HasPrefix(line, "-") && !strings.HasPrefix(line, "---") {
			color.Red(line)
		} else if strings.HasPrefix(line, " ") {
			fmt.Println(line)
		} else {
			fmt.Println(line)
		}
	}
}
