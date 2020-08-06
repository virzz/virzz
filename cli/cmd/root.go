package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/virink/virzz/common"
)

var rootCmd = &cobra.Command{
	Use:   common.AppName,
	Short: "A tools for terminal",
}

// Execute -
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
