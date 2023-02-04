package main

import (
	"os"

	"github.com/virzz/logger"
	"github.com/virzz/virzz/common"
	"github.com/virzz/virzz/modules/crypto/hashpow"
)

var (
	AppName        = "Hashpow"
	BinName        = "hashpow"
	Version string = "latest"
	BuildID string = "0"
)

func main() {
	rootCmd := hashpow.ExportCommand()[0]
	rootCmd.SilenceErrors = true
	rootCmd.AddCommand(common.VersionCommand(AppName, Version, BuildID))
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
