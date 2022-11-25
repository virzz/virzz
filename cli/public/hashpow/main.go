package main

import (
	"os"

	"github.com/mozhu1024/virzz/common"
	"github.com/mozhu1024/virzz/crypto/hash"
)

var (
	AppName        = "Hashpow"
	BinName        = "hashpow"
	Version string = ""  // git tag | tail -1
	BuildID string = "0" // head .buildid
)

func init() {
	hash.HashPowCmd.AddCommand(common.VersionCommand(AppName, Version, BuildID))
}

func main() {
	if err := hash.HashPowCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
