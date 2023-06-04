package gostrip

import (
	"github.com/urfave/cli/v3"
)

var Cmd = &cli.Command{
	Category:    "Ext",
	Name:        "gostrip",
	Aliases:     []string{"strip"},
	Usage:       "Strip golang binary file",
	Description: `Modify from [Go-strip @w8ay] https://github.com/boy-hack/go-strip`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "force",
			Aliases: []string{"f"},
			Usage:   "Force to strip 'Go Struct Name' (未完全测试,请谨慎使用)",
			Value:   false,
		},
	},
	Action: func(c *cli.Context) error {
		opts := &Opts{}
		if c.Bool("force") {
			opts.IsForce = true
		}
		return Strip(c.Args().First(), opts)
	},
}
