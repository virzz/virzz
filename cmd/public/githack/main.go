package main

import (
	"github.com/virzz/virzz/cmd/public"
	"github.com/virzz/virzz/modules/web/githack"
)

var (
	Version string = "-"
)

func main() {
	public.RunCliApp(githack.Cmd, "githack", Version)
}
