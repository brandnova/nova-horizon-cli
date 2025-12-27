package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/brandnova/nova-horizon-cli/internal/agent"
	"github.com/brandnova/nova-horizon-cli/internal/config"
	"github.com/spf13/cobra"
)

var (
	workDir   string
	verbose   bool
	dryRun    bool
	model     string
	maxSteps  int
	allowRun  bool
	applyDiff bool
	showInfo  bool
)

var rootCmd = &cobra.Command{
	Use:   "nova-hrzn [prompt]",
	Short: "Local AI coding agent powered by Gemini",
	Long: `Nova Horizon is a local coding agent that helps you with file operations and code execution.
It uses the Google Gemini API to understand and execute your requests.

Examples:
  nova-hrzn "Create a hello world program in Python"
  nova-hrzn --dir ./myproject "List all files in this directory"
  nova-hrzn --verbose --allow-run "Execute my test script"
  nova-hrzn --info`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Just show info if requested
		if showInfo {
			printBanner()
			printInfo()
			return nil
		}

		// Interactive shell mode if no args
		if len(args) == 0 {
			printBanner()
			printInfo()
			return runShell()
		}

		return runAgent(args[0])
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&workDir, "dir", "d", "", "Working directory (default: current directory)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Show what would be done without making changes")
	rootCmd.PersistentFlags().StringVar(&model, "model", "gemini-2.5-flash", "Model to use")
	rootCmd.PersistentFlags().IntVar(&maxSteps, "max-steps", 10, "Maximum agent loop iterations")
	rootCmd.PersistentFlags().BoolVar(&allowRun, "allow-run", false, "Allow execution of programs")
	rootCmd.PersistentFlags().BoolVar(&applyDiff, "apply", false, "Automatically apply file changes without confirmation")
	rootCmd.PersistentFlags().BoolVar(&showInfo, "info", false, "Show information about Nova Horizon")
}

func runShell() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Entering interactive mode. Type 'exit' to quit.")

	for {
		fmt.Print("\nnova-hrzn> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		input = strings.TrimSpace(input)
		if input == "exit" || input == "quit" {
			fmt.Println("Goodbye!")
			break
		}

		if input == "" {
			continue
		}

		if err := runAgent(input); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
	return nil
}

func runAgent(prompt string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Resolve working directory (logic same as before, ensures absolute path)
	resolvedWorkDir := workDir
	if resolvedWorkDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get working directory: %w", err)
		}
		resolvedWorkDir = wd
	} else {
		absPath, err := filepath.Abs(resolvedWorkDir)
		if err != nil {
			return fmt.Errorf("invalid working directory: %w", err)
		}
		resolvedWorkDir = absPath
	}

	// Create and run agent
	agentConfig := &agent.Config{
		APIKey:    cfg.APIKey,
		Model:     model,
		WorkDir:   resolvedWorkDir,
		Verbose:   verbose,
		DryRun:    dryRun,
		MaxSteps:  maxSteps,
		AllowRun:  allowRun,
		ApplyDiff: applyDiff,
	}

	ag := agent.NewAgent(agentConfig)
	return ag.Run(prompt)
}

func printBanner() {
	banner := `
  _   _                  _   _            _              
 | \ | |                | | | |          (_)             
 |  \| | ___ __   __ _  | |_| | ___  _ __ _ _______  _ __ 
 | .   |/ _ \ \ / / _ \ |  _  |/ _ \| '__| |_  / _ \| '_ \ 
 | |\  | (_) \ V / (_| || | | | (_) | |  | |/ / (_) | | | |
 \_| \_/\___/ \_/ \__,_|\_| |_/\___/|_|  |_/___\___/|_| |_|
                                                           
 Local AI Coding Agent`
	fmt.Println(banner)
}

func printInfo() {
	fmt.Println(`
Creator: Brand Nova
Description: A powerful CLI tool that uses Google's Gemini models to act as your intelligent coding assistant.

Quick Start Commands:
  nova-hrzn --info                    # Show this information
  nova-hrzn "your prompt"             # Run a single command
  nova-hrzn                           # Enter interactive shell mode`)
}

func Execute() error {
	return rootCmd.Execute()
}
