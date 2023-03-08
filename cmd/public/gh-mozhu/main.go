package main

import (
	"github.com/virzz/virzz/cmd/public"
	"github.com/virzz/virzz/modules/exts/ghext"
)

var (
	Version string = "-"
)

func main() {
	public.RunCliApp(ghext.Cmd, "gh-mozhu", Version)
}
