package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/virink/virzz/common"
)

var (
	showVersion bool = false
)

var rootCmd = &cobra.Command{
	Use:   common.BinName,
	Short: "A tools for terminal",
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			printVersion()
			return
		}
		cmd.Help()
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&showVersion, "version", "V", false, "Show Version")
	rootCmd.SuggestionsMinimumDistance = 1
}

// Execute -
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
