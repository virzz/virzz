package main

import (
	"os"

	"github.com/mozhu1024/virzz/common"
	"github.com/mozhu1024/virzz/web/jwttool"
)

var (
	AppName        = "JWTTool"
	BinName        = "jwttool"
	Version string = ""  // git tag | tail -1
	BuildID string = "0" // head .buildid
)

func init() {
	jwttool.JWTToolCmd.AddCommand(common.VersionCommand(AppName, Version, BuildID))
}

func main() {
	if err := jwttool.JWTToolCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
