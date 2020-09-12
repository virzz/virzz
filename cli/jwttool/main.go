package main

import (
	"os"

	"github.com/siddontang/go/log"
	"github.com/spf13/cobra"
	"github.com/virink/virzz/common"
)

var rootCmd = &cobra.Command{
	Use:   "jwttool",
	Short: "A jwt tool with Print/Crack/Modify",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		level := log.LevelError
		// Debug Mode
		if common.DebugMode {
			level = log.LevelDebug
		}
		// Logger
		common.InitLogger(level)
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.SuggestionsMinimumDistance = 1
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
