package main

import (
	"os"

	"github.com/mozhu1024/virzz/common"
	"github.com/mozhu1024/virzz/web/gopher"
)

var (
	AppName        = "Gopher"
	BinName        = "gopher"
	Version string = ""  // git tag | tail -1
	BuildID string = "0" // head .buildid
)

func init() {
	gopher.GopherCmd.AddCommand(common.VersionCommand(AppName, Version, BuildID))
}

func main() {
	if err := gopher.GopherCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
