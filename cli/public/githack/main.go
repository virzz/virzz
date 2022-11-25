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

func init() {
	githack.GithackCmd.AddCommand(common.VersionCommand(AppName, Version, BuildID))
}

func main() {
	if err := githack.GithackCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
