package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
)

type Config struct {
	APIKey string `toml:"api_key"`
}

func LoadConfig() (*Config, error) {
	// First, try environment variable
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey != "" {
		return &Config{APIKey: apiKey}, nil
	}

	// Then try config file
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("GEMINI_API_KEY not set and could not read home directory")
	}

	configPath := filepath.Join(homeDir, ".config", "nova-horizon", "config.toml")
	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("GEMINI_API_KEY environment variable not set and no config file found at %s", configPath)
		}
		return nil, fmt.Errorf("failed to read config file at %s: %w", configPath, err)
	}

	var cfg Config
	if err := toml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file at %s: %w", configPath, err)
	}

	if cfg.APIKey == "" {
		return nil, fmt.Errorf("config file at %s parsed successfully but 'api_key' is missing or empty", configPath)
	}

	return &cfg, nil
}
