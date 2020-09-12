package main

import (
	"os"

	"github.com/siddontang/go/log"
	"github.com/spf13/cobra"
	"github.com/virink/virzz/common"
)

var rootCmd = &cobra.Command{
	Use:           "gopher",
	Short:         "Generate Gopher Exp",
	SilenceErrors: true,
	SilenceUsage:  true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		level := log.LevelError
		// Debug Mode
		if common.DebugMode {
			level = log.LevelDebug
		}
		// Logger
		common.InitLogger(level)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.SuggestionsMinimumDistance = 1
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		common.Logger.Error(err)
		os.Exit(1)
	}
}
