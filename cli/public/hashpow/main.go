package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/common"
	"github.com/virzz/virzz/modules/crypto/hashpow"
)

var (
	AppName        = "Hashpow"
	BinName        = "hashpow"
	Version string = "latest"
	BuildID string = "0"
)


func main() {
	cmd := hashpow.Cmd
	app := &cli.App{
		Authors:         []any{fmt.Sprintf("%s <%s>", common.Author, common.Email)},
		Name:            BinName,
		Usage:           cmd.Usage,
		HideVersion:     true,
		HideHelpCommand: true,
		Suggest:         true,
		Action: func(c *cli.Context) error {
			cmd.HelpName = BinName
			c.Command = cmd
			return cmd.Action(c)
		},
	}
	if err := app.Run(os.Args); err != nil {
		logger.Error(err)
	}
}
