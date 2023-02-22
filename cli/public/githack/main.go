package main

import (
	"github.com/virzz/virzz/cli/public"
	"github.com/virzz/virzz/modules/web/githack"
)

var (
	Version string = "-"
)

func main() {
	public.RunCliApp(githack.Cmd, "githack", Version)
}
