package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:           "lumi-cli",
	Short:         "A CLI for lumi & lumina services",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
