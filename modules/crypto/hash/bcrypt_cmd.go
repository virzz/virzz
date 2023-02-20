package hash

import (
	"fmt"

	"github.com/urfave/cli/v3"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/utils"
	"golang.org/x/crypto/bcrypt"
)

var BcryptCmd = &cli.Command{
	Category: "Crypto",
	Name:     "bcrypt",
	Usage:    "Bcrypt Generate/Compare",
	Commands: []*cli.Command{
		&cli.Command{
			Name:    "generate",
			Aliases: []string{"gen", "g"},
			Usage:   "Generate",
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:    "cost",
					Aliases: []string{"c"},
					Usage:   "bcrypt cost",
					Value:   bcrypt.DefaultCost,
					Action: func(c *cli.Context, cost int) error {
						if cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
							return fmt.Errorf("cost must be between %d and %d", bcrypt.MinCost, bcrypt.MaxCost)
						}
						return nil
					},
				},
			},
			Action: func(c *cli.Context) (err error) {
				code, err := utils.ArgOrPipe(c.Args().First())
				if err != nil {
					return err
				}
				r, err := BcryptGenerate(code, c.Int("cost"))
				if err != nil {
					return err
				}
				_, err = fmt.Println(r)
				return
			},
		},
		&cli.Command{
			Name:      "compare",
			Aliases:   []string{"comp", "c"},
			Usage:     "Compare",
			ArgsUsage: "hashed password",
			Action: func(c *cli.Context) (err error) {
				if c.NArg() < 2 {
					return fmt.Errorf("missing parameter")
				}
				err = BcryptCompare(c.Args().Get(0), c.Args().Get(1))
				if err != nil {
					return err
				}
				logger.Success("Compare OK")
				return nil
			},
		},
	},
}
