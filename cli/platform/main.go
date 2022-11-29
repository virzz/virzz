package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/mozhu1024/virzz/logger"

	"github.com/mozhu1024/virzz/common"
	"github.com/spf13/cobra"
)

var (
	showVersion  bool = false
	setDebugMode bool

	AppName        = "Unknown Platform"
	BinName        = "virzz-platform"
	Version string = "dev" // git tag | tail -1
	BuildID string = "0"   // head .buildid
)

var versionCmd = common.VersionCommand(AppName, Version, BuildID)

var rootCmd = &cobra.Command{
	Use:           BinName,
	Short:         "The Cyber Swiss Army Knife for platform",
	SilenceErrors: true,
	SilenceUsage:  true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Args
		if len(args) > 0 {
			// FIX ~ as HOME
			if strings.HasPrefix(args[0], "~/") {
				if home := os.Getenv("HOME"); home != "" {
					args[0] = filepath.Join(home, args[0][2:len(args[0])])
				}
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			versionCmd.Run(cmd, args)
			return
		}
		cmd.Help()
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&showVersion, "version", "V", false, "Show Version")
	rootCmd.PersistentFlags().BoolVarP(&setDebugMode, "debug", "D", false, "Set Debug Mode")
	rootCmd.SuggestionsMinimumDistance = 1
	rootCmd.AddCommand(versionCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
