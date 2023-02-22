package main

import (
	"github.com/virzz/virzz/cli/public"
	"github.com/virzz/virzz/modules/crypto/hashpow"
)

var (
	Version string = "-"
)

func main() {
	public.RunCliApp(hashpow.Cmd, "jwttool", Version)
}
