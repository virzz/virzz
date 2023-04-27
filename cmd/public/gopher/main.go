package main

import (
	"github.com/virzz/virzz/cmd/public"
	"github.com/virzz/virzz/modules/web/gopher"
)

var (
	Version string = "-"
)

func main() {
	public.RunCliApp(gopher.Cmd, "gopher", Version)
}
