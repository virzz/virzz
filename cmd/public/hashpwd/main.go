package main

import (
	"github.com/virzz/virzz/cmd/public"
	"github.com/virzz/virzz/modules/crypto/hashpwd"
)

var (
	Version string = "-"
)

func main() {
	public.RunCliApp(hashpwd.Cmd, "hashpwd", Version)
}
