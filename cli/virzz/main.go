package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/urfave/cli/v3"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/common"
	"github.com/virzz/virzz/modules/crypto/basex"
	"github.com/virzz/virzz/modules/crypto/basic"
	"github.com/virzz/virzz/modules/crypto/hash"
	"github.com/virzz/virzz/modules/crypto/hashpow"
	"github.com/virzz/virzz/modules/parser"
	"github.com/virzz/virzz/modules/tools/domain"
	"github.com/virzz/virzz/modules/web/gopher"
	"github.com/virzz/virzz/modules/web/jwttool"
	"github.com/virzz/virzz/modules/web/leakcode/githack"
	"github.com/virzz/virzz/utils"
)

const (
	AppName = "Virzz"
	BinName = "virzz"
)

var (
	Version  string = "latest"
	BuildID  string = "0"
	Revision string = ""
)

func init() {
	// rootCmd.AddCommand(versionCmd, common.CompletionCommand(), aliasCommand())

	// CMD
	// Crypto
	// rootCmd.AddCommand(basex.ExportCommand()...)
	// rootCmd.AddCommand(basic.ExportCommand()...)
	// rootCmd.AddCommand(classical.ExportCommand()...)
	// rootCmd.AddCommand(hash.ExportCommand()...)
	// // Tools
	// rootCmd.AddCommand(qrcode.ExportCommand()...)
	// rootCmd.AddCommand(netool.ExportCommand()...)
	// rootCmd.AddCommand(dsstore.ExportCommand()...)
}

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("Ver: %s (build-%s) revision=%s\n", c.App.Version, BuildID, Revision)
	}
	app := &cli.App{
		Name:                       BinName,
		Authors:                    []any{fmt.Sprintf("%s <%s>", common.Author, common.Email)},
		Usage:                      "The Cyber Swiss Army Knife for terminal",
		Version:                    Version,
		Suggest:                    true,
		EnableShellCompletion:      true,
		HideHelpCommand:            true,
		ShellCompletionCommandName: "completion",
		Action: func(c *cli.Context) error {
			// Link Binary to run subCommand
			runName := path.Base(os.Args[0])
			// Remove .{ext}
			runName = strings.TrimSuffix(runName, path.Ext(runName))
			if runName != BinName {
				if cmd := c.Command.Command(runName); cmd != nil {
					cmd.HelpName = runName
					c.Command = cmd
					if err := cmd.Action(c); err != nil {
						return err
					}
					return nil
				}
				cli.ShowAppHelp(c)
				return fmt.Errorf("not found command: %s", runName)
			} else if c.NArg() > 0 {
				cli.ShowAppHelp(c)
				return fmt.Errorf("not found command: %s", c.Args().First())
			}
			return cli.ShowAppHelp(c)
		},
	}
	// Add SubCommands
	app.Commands = append(app.Commands, aliasCmd,
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
	// HideHelpCommand
	utils.HideHelpCommand(app.Commands)

	// TODO(app.Name, app.Commands)

	if err := app.Run(os.Args); err != nil {
		logger.Error(err)
	}
}
