package hashpow

import (
	"fmt"

	"github.com/urfave/cli/v3"
	"github.com/virzz/virzz/utils"
)

var methodLength = map[string]int{"md5": 32, "sha1": 40}

var Cmd = &cli.Command{
	Category: "Misc",
	Name:     "hashpow",
	Usage:    "Brute Hash Power of Work with md5/sha1",

	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "code",
			Aliases:  []string{"c"},
			Value:    "",
			Usage:    "Request code",
			Required: true,
		},
		&cli.IntFlag{
			Name:    "pos",
			Aliases: []string{"i"},
			Value:   0,
			Usage:   "Starting position of hash",
			Action: func(c *cli.Context, p int) error {
				if p < 0 || methodLength[c.String("method")] < p {
					return fmt.Errorf("invalid position")
				}
				return nil
			},
		},
		&cli.StringFlag{
			Name:    "prefix",
			Aliases: []string{"p"},
			Value:   "",
			Usage:   "Hash prefix",
		},
		&cli.StringFlag{
			Name:    "suffix",
			Aliases: []string{"s"},
			Value:   "",
			Usage:   "Hash suffix",
		},
		&cli.StringFlag{
			Name:    "method",
			Aliases: []string{"m"},
			Value:   "md5",
			Usage:   "Hash method: <sha1|md5>",
			Action: func(_ *cli.Context, m string) error {
				if utils.SliceContains([]string{"md5", "sha1"}, m) {
					return nil
				}
				return fmt.Errorf("invalid method")
			},
		},
	},
	Action: func(c *cli.Context) error {
		r := HashPoW(c.String("code"), c.String("prefix"), c.String("suffix"), c.String("method"), c.Int("pos"))
		_, err := fmt.Println(r)
		return err
	},
}
