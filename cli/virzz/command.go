package main

import (
	"github.com/urfave/cli/v3"
	"github.com/virzz/virzz/modules/crypto/basex"
	"github.com/virzz/virzz/modules/crypto/basic"
	"github.com/virzz/virzz/modules/crypto/hash"
	"github.com/virzz/virzz/modules/crypto/hashpow"
	"github.com/virzz/virzz/modules/parser"
	"github.com/virzz/virzz/modules/tools/domain"
	"github.com/virzz/virzz/modules/web/gopher"
	"github.com/virzz/virzz/modules/web/jwttool"
	"github.com/virzz/virzz/modules/web/leakcode/githack"
)

var commands = []*cli.Command{}

func init() {
	// Add SubCommands
	commands = append(commands, aliasCmd,
		githack.Cmd,
		gopher.Cmd,
		hashpow.Cmd,
		jwttool.Cmd,
		parser.Cmd,
		domain.Cmd,
		basic.Cmd,
		basex.Cmd,

		hash.BcryptCmd,
	)
}
