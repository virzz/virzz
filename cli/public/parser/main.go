package main

import (
	"os"

	"github.com/virzz/virzz/common"
	"github.com/virzz/virzz/modules/parser"
)

var (
	AppName        = "Parser"
	BinName        = "parser"
	Version string = ""  // git tag | tail -1
	BuildID string = "0" // head .buildid
)

func main() {
	rootCmd := parser.ExportCommand()[0]
	rootCmd.AddCommand(common.VersionCommand(AppName, Version, BuildID))
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
