package resh

import (
	"fmt"

	"github.com/urfave/cli/v3"
)

var Cmd = &cli.Command{
	Category: "Misc",
	Name:     "resh",
	Aliases:  []string{"reshell"},
	Usage:    "Reverse Shell Template Generator",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "ip",
			Aliases: []string{"addr"},
			Usage:   "Your IP Address",
		},
		&cli.IntFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Usage:   "Your Port",
			Value:   1337,
		},
	},
	Action: func(c *cli.Context) (err error) {
		addr := c.String("ip")
		port := c.Int("port")
		if c.NArg() > 0 {
			addr = c.Args().First()
		}
		r, err := ReverseShell(addr, port)
		if err != nil {
			return
		}
		_, err = fmt.Println(r)
		return
	},
}
