package parser

import (
	"fmt"

	"github.com/urfave/cli/v3"
	"github.com/virzz/virzz/utils"
)

var Cmd = &cli.Command{
	Category: "Misc",
	Name:     "parser",
	Usage:    "Parse some file",
	Commands: []*cli.Command{
		&cli.Command{
			Category: "Parser",
			Name:     "procnet",
			Aliases:  []string{"net"},
			Usage:    "Parse /proc/net/tcp|udp",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "filepath",
					Aliases: []string{"f"},
					Usage:   "Which file for parse",
				},
			},
			Action: func(c *cli.Context) (err error) {
				filepath := c.String("filepath")
				if filepath == "" {
					if c.NArg() > 0 {
						filepath = c.Args().First()
					} else if data, err := utils.GetFromPipe(); err == nil {
						filepath = string(data)
					} else {
						return fmt.Errorf("invalid filepath")
					}
				}
				r, err := ParseProcNet(filepath)
				if err != nil {
					return err
				}
				_, err = fmt.Println(r)
				return
			},
		},
		&cli.Command{
			Category: "Parser",
			Name:     "todo",
			Usage:    "Parse todo",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "filepath",
					Aliases: []string{"f"},
					Usage:   "Which file for parse",
				},
			},
			Action: func(c *cli.Context) (err error) {
				return fmt.Errorf("TODO")
			},
		},
	},
}
