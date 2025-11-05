package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
)

func CreateConfig(endpoint, output, accessKey, secretAccessKey string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}
	configDir := filepath.Join(home, ".lumi-cli")
	configPath := filepath.Join(configDir, "config.toml")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}
	config := Profile{
		Endpoint:  endpoint,
		Output:    output,
		VerifySSL: true,
		AccessKey: accessKey,
		SecretKey: secretAccessKey,
	}
	data, err := toml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	if err := os.WriteFile(configPath, data, 0600); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	fmt.Printf("Config file created at: %s\n", configPath)
	return nil
}
