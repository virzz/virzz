package main

import (
	"github.com/virzz/virzz/cmd/public"
	"github.com/virzz/virzz/modules/misc/hashpow"
)

var (
	Version string = "-"
)

func main() {
	public.RunCliApp(hashpow.Cmd, "hashpow", Version)
}
