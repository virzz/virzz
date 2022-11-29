package main

import (
	"os"

	"github.com/mozhu1024/virzz/common"
	"github.com/mozhu1024/virzz/web/leakcode/githack"
)

var (
	AppName        = "GitHack"
	BinName        = "githack"
	Version string = ""  // git tag | tail -1
	BuildID string = "0" // head .buildid
)

func main() {
	rootCmd := githack.ExportCommand()[0]
	rootCmd.AddCommand(common.VersionCommand(AppName, Version, BuildID))
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
