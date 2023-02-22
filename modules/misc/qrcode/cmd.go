package qrcode

import (
	"fmt"

	"github.com/urfave/cli/v3"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/utils"
)

var Cmd = &cli.Command{
	Category: "Misc",
	Name:     "qrcode",
	Aliases:  []string{"qr"},
	Usage:    "A qrcode tool for terminal",
	Commands: []*cli.Command{
		// bs
		&cli.Command{
			Name:  "bs",
			Usage: "Bin String (0,1) to Qrcode Image",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "exchange",
					Aliases: []string{"c"},
					Usage:   "Exchange 0/1",
				},
				&cli.StringFlag{
					Name:    "output",
					Aliases: []string{"o"},
					Usage:   "Which file for parse",
				},
			},
			Action: func(c *cli.Context) (err error) {
				content, err := utils.GetArgFilePipe(c.Args().First())
				if err != nil {
					return err
				}
				r, err := ZeroOneToQrcode(string(content), c.Bool("exchange"), c.String("output"))
				if err != nil {
					return err
				}
				_, err = fmt.Println(r)
				return
			},
		},
		// parse
		&cli.Command{
			Name:    "parse",
			Usage:   "Parse qrcode image",
			Aliases: []string{"p"},
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "terminal",
					Aliases: []string{"t"},
					Usage:   "Also print qrimage to terminal",
				},
			},
			Action: func(c *cli.Context) error {
				target := c.Args().First()
				logger.Debug(target)
				if err := utils.ValidArg(target, "url|file"); err != nil {
					return err
				}
				r, err := ParseQrcode(target, c.Bool("terminal"))
				if err != nil {
					return err
				}
				logger.Success("Parsed: ", r)
				_, err = fmt.Println(r)
				return err
			},
		},
		// generate
		&cli.Command{
			Name:    "generate",
			Usage:   "Generate qrcode image",
			Aliases: []string{"gen", "g"},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "output",
					Aliases: []string{"o"},
					Usage:   "save to output",
				},
			},
			Action: func(c *cli.Context) (err error) {
				content, err := utils.GetArgFilePipe(c.Args().First())
				if err != nil {
					return err
				}
				r, err := GenerateQrcode(string(content), c.String("output"))
				if err != nil {
					return err
				}
				_, err = fmt.Println(r)
				return
			},
		},
	},
}
