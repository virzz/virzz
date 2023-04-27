package basic

import (
	"fmt"

	"github.com/urfave/cli/v3"
)

var urlencodeCmd = &cli.Command{
	Category: "URL",
	Name:     "urle",
	Aliases:  []string{"urlencode"},
	Usage:    "URL Encode",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "raw",
			Aliases: []string{"s", "safe"},
			Usage:   "Raw encode. + -> %20",
		},
	},
	Action: func(c *cli.Context) (err error) {
		r, err := URLEncode(c.Args().First(), c.Bool("raw"))
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}

var urldecodeCmd = &cli.Command{
	Category: "URL",
	Name:     "urld",
	Aliases:  []string{"urldecode"},
	Usage:    "URL Decode",
	Action: func(c *cli.Context) (err error) {
		r, err := URLDecode(c.Args().First())
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}
