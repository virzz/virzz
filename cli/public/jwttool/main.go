package main

import (
	"github.com/virzz/virzz/cli/public"
	"github.com/virzz/virzz/modules/web/jwttool"
)

var (
	Version string = "-"
)

func main() {
	public.RunCliApp(jwttool.Cmd, "jwttool", Version)
}
