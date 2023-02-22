package githack

import (
	"fmt"

	"github.com/urfave/cli/v3"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/utils"
)

var Cmd = &cli.Command{
	Category: "Web",
	Name:     "githack",
	Usage:    "A `.git` folder disclosure exploit",
	Flags: []cli.Flag{
		&cli.Int64Flag{
			Name:    "limit",
			Aliases: []string{"l"},
			Value:   10,
			Usage:   "Request limit",
		},
		&cli.Int64Flag{
			Name:    "delay",
			Aliases: []string{"d"},
			Value:   0,
			Usage:   "Request delay",
		},
		&cli.Int64Flag{
			Name:    "timeout",
			Aliases: []string{"t"},
			Value:   10,
			Usage:   "Request timeout",
		},
	},
	Action: func(c *cli.Context) error {
		if c.NArg() < 1 {
			cli.ShowSubcommandHelp(c)
			return fmt.Errorf("require a target url arg")
		}
		targetURL := c.Args().First()
		if err := utils.ValidArg(targetURL, "url"); err != nil {
			return err
		}
		logger.DebugF("Target url: %s", targetURL)
		return gitHack(targetURL, c.Int64("limit"), c.Int64("delay"), c.Int64("timeout"))
	},
}
