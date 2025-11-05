package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/LumiServices/lumi-cli/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration for lumi-cli",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}
		fmt.Println("Loaded config:", cfg)
		return nil
	},
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new configuration file",
	Long:  "Create a new configuration file at ~/.lumi-cli/config.toml",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter endpoint URL: ")
		endpoint, _ := reader.ReadString('\n')
		endpoint = strings.TrimSpace(endpoint)
		fmt.Print("Enter default output format [xml]: ")
		output, _ := reader.ReadString('\n')
		output = strings.TrimSpace(output)
		if output == "" {
			output = "xml"
		}
		fmt.Print("Enter access key: ")
		accessKey, _ := reader.ReadString('\n')
		accessKey = strings.TrimSpace(accessKey)
		fmt.Print("Enter secret access key: ")
		secretKey, _ := reader.ReadString('\n')
		secretKey = strings.TrimSpace(secretKey)
		if endpoint == "" {
			return fmt.Errorf("endpoint cannot be empty")
		}
		if accessKey == "" {
			return fmt.Errorf("access key cannot be empty")
		}
		if secretKey == "" {
			return fmt.Errorf("secret access key cannot be empty")
		}
		err := config.CreateConfig(endpoint, output, accessKey, secretKey)
		if err != nil {
			return fmt.Errorf("failed to create config: %w", err)
		}
		fmt.Println("\nConfiguration initialized successfully!")
		return nil
	},
}
