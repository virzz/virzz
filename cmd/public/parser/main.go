package main

import (
	"github.com/virzz/virzz/cmd/public"
	"github.com/virzz/virzz/modules/misc/parser"
)

var (
	Version string = "-"
)

func main() {
	public.RunCliApp(parser.Cmd, "parser", Version)
}
