package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
	"github.com/urfave/cli/v3"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/common"
	"github.com/virzz/virzz/utils"
)

const (
	BinName = "enyo"
)

var (
	Version  string = "latest"
	BuildID  string = "0"
	Revision string = ""
)

func init() {
	viper.AddConfigPath("$HOME/.config/virzz")
	viper.SetConfigName("tongji")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.ReadInConfig()
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

	// Load Commands
	app.Commands = append(app.Commands, commands...)
	// HideHelpCommand
	utils.HideHelpCommand(app.Commands)

	if err := app.Run(os.Args); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
