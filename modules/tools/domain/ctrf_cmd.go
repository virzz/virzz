package domain

import (
	"fmt"

	"github.com/urfave/cli/v3"
)

var Cmd = &cli.Command{
	Category: "Misc",
	Name:     "domain",
	Usage:    "Some tools for Domain/SubDomain",
	Commands: []*cli.Command{
		&cli.Command{
			Category: "Domain",
			Name:     "ctfr",
			Usage:    "滥用证书透明记录 By https://crt.sh",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "domain",
					Aliases: []string{"f"},
					Usage:   "Which file for parse",
				},
			},
			Action: func(c *cli.Context) (err error) {
				domain := c.String("domain")
				if domain == "" {
					if c.NArg() > 0 {
						domain = c.Args().First()
					} else {
						return fmt.Errorf("invalid domain")
					}
				}
				r, err := Ctfr(domain)
				if err != nil {
					return err
				}
				_, err = fmt.Println(r)
				return
			},
		},
		&cli.Command{
			Category: "Domain",
			Name:     "todo",
			Usage:    "TODO",
			Flags:    []cli.Flag{},
			Action: func(c *cli.Context) (err error) {
				return fmt.Errorf("TODO")
			},
		},
	},
}
