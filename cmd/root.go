package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/brandnova/nova-horizon-cli/internal/agent"
	"github.com/brandnova/nova-horizon-cli/internal/config"
)

var (
	workDir   string
	verbose   bool
	dryRun    bool
	model     string
	maxSteps  int
	allowRun  bool
	applyDiff bool
)

var rootCmd = &cobra.Command{
	Use:   "nova-horizon [prompt]",
	Short: "Local AI coding agent powered by Gemini",
	Long: `Nova Horizon is a local coding agent that helps you with file operations and code execution.
It uses the Google Gemini API to understand and execute your requests.

Examples:
  nova-horizon "Create a hello world program in Python"
  nova-horizon --dir ./myproject "List all files in this directory"
  nova-horizon --verbose --allow-run "Execute my test script"`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
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
}

func runAgent(prompt string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Resolve working directory
	if workDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get working directory: %w", err)
		}
		workDir = wd
	} else {
		absPath, err := filepath.Abs(workDir)
		if err != nil {
			return fmt.Errorf("invalid working directory: %w", err)
		}
		workDir = absPath
	}

	// Print startup banner
	printBanner()

	// Create and run agent
	agentConfig := &agent.Config{
		APIKey:    cfg.APIKey,
		Model:     model,
		WorkDir:   workDir,
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
	fmt.Println(`Nova Horizon
────────────────
Local Coding Agent`)
	fmt.Printf("Working directory: %s\n\n", workDir)
}

func Execute() error {
	return rootCmd.Execute()
}
