package main

import (
	"github.com/virzz/virzz/cmd/public"
	"github.com/virzz/virzz/internal/gostrip"
)

var (
	Version string = "-"
)

func main() {
	public.RunCliApp(gostrip.Cmd, "gostrip", Version)
}
