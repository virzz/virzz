package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/virzz/logger"
	"github.com/virzz/virzz/common"
)

var (
	AppName        = "VirzzPlatform"
	BinName        = "virzz-platform"
	Version string = "dev"
	BuildID string = "0"
)

var versionCmd = common.VersionCommand(AppName, Version, BuildID)

var rootCmd = &cobra.Command{
	Use:   BinName,
	Short: "The Cyber Swiss Army Knife for platform",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func main() {
	rootCmd.AddCommand(versionCmd, configCmd, daemonCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
