package main

import (
	"github.com/urfave/cli/v3"
	"github.com/virzz/virzz/modules/crypto/basex"
	"github.com/virzz/virzz/modules/crypto/basic"
	"github.com/virzz/virzz/modules/crypto/classical"
	"github.com/virzz/virzz/modules/crypto/hash"
	"github.com/virzz/virzz/modules/crypto/hashpwd"
	"github.com/virzz/virzz/modules/exts/ghext"
	"github.com/virzz/virzz/modules/exts/gostrip"
	"github.com/virzz/virzz/modules/misc/domain"
	"github.com/virzz/virzz/modules/misc/hashpow"
	"github.com/virzz/virzz/modules/misc/nettool"
	"github.com/virzz/virzz/modules/misc/parser"
	"github.com/virzz/virzz/modules/misc/qrcode"
	"github.com/virzz/virzz/modules/misc/resh"
	"github.com/virzz/virzz/modules/web/githack"
	"github.com/virzz/virzz/modules/web/gopher"
	"github.com/virzz/virzz/modules/web/jwttool"
)

var commands = []*cli.Command{}

func init() {
	// Add SubCommands
	commands = append(commands, aliasCmd,
		// Crypto
		basex.Cmd,
		basic.Cmd,
		classical.Cmd,
		hash.Cmd,
		hash.BcryptCmd,
		hashpwd.Cmd,
		// Misc
		domain.Cmd,
		nettool.Cmd,
		hashpow.Cmd,
		qrcode.Cmd,
		parser.Cmd,
		ghext.Cmd,
		resh.Cmd,
		// Web
		githack.Cmd,
		gopher.Cmd,
		jwttool.Cmd,
		// Ext
		gostrip.Cmd,
	)
}
