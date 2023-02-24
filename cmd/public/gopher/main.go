package main

import (
	"github.com/virzz/virzz/cmd/public"
	"github.com/virzz/virzz/modules/web/jwttool"
)

var (
	Version string = "-"
)

func main() {
	public.RunCliApp(jwttool.Cmd, "gopher", Version)
}
