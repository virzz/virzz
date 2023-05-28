package parser

import (
	"fmt"

	"github.com/urfave/cli/v3"
	"github.com/virzz/virzz/utils"
)

var DsStoreCmd = &cli.Command{
	Name:  "dsstore",
	Usage: ".DS_Store Parser",
	Action: func(c *cli.Context) (err error) {
		target := c.Args().First()
		if err := utils.ValidArg(target, "url|file"); err != nil {
			return err
		}
		r, err := DSStore(target)
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}
