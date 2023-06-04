package main

import (
	"github.com/virzz/virzz/cmd/public"
	"github.com/virzz/virzz/modules/gormgen"
)

var (
	Version string = "-"
)

func main() {
	public.RunCliApp(gormgen.Cmd, "gormgen", Version)
}
