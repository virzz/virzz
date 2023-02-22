package classical

import (
	"fmt"

	"github.com/urfave/cli/v3"
	"github.com/virzz/virzz/utils"
)

var Cmd = &cli.Command{
	Category: "Crypto",
	Name:     "classical",
	Usage:    "Classical cryptography",
	Commands: []*cli.Command{
		// caesar
		&cli.Command{
			Category: "Classical",
			Name:     "caesar",
			Usage:    "Caesar Encode",
			Action: func(c *cli.Context) (err error) {
				code, err := utils.ArgOrPipe(c.Args().First())
				if err != nil {
					return err
				}
				r, _ := Caesar(code)
				_, err = fmt.Println(r)
				return
			},
		},
		// rot13
		&cli.Command{
			Category: "Classical",
			Name:     "rot13",
			Usage:    "Rot13 Encode",
			Action: func(c *cli.Context) (err error) {
				code, err := utils.ArgOrPipe(c.Args().First())
				if err != nil {
					return err
				}
				r, _ := Rot13(code)
				_, err = fmt.Println(r)
				return
			},
		},
		// atbash
		&cli.Command{
			Category: "Classical",
			Name:     "atbash",
			Usage:    "Atbash 埃特巴什码",
			Action: func(c *cli.Context) (err error) {
				code, err := utils.ArgOrPipe(c.Args().First())
				if err != nil {
					return err
				}
				r, _ := Atbash(code)
				_, err = fmt.Println(r)
				return
			},
		},

		// peigen
		&cli.Command{
			Category: "Classical",
			Name:     "peigen",
			Usage:    "Peigen 培根密码",
			Action: func(c *cli.Context) (err error) {
				code, err := utils.ArgOrPipe(c.Args().First())
				if err != nil {
					return err
				}
				r, _ := Peigen(code)
				_, err = fmt.Println(r)
				return
			},
		},

		// vigenere
		&cli.Command{
			Category: "Classical",
			Name:     "vigenere",
			Usage:    "Vigenere 维吉利亚密码",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "key",
					Aliases: []string{"k"},
					Usage:   "Vigenere Key",
					Value:   "virzz",
				},
				&cli.StringFlag{
					Name:    "decode",
					Aliases: []string{"d"},
					Usage:   "Decode",
				},
			},
			Action: func(c *cli.Context) (err error) {
				code, err := utils.ArgOrPipe(c.Args().First())
				if err != nil {
					return err
				}
				r, _ := Vigenere(code, c.String("key"), c.Bool("decode"))
				_, err = fmt.Println(r)
				return
			},
		},
		// morse
		&cli.Command{
			Category: "Classical",
			Name:     "morse",
			Usage:    "Morse Code 摩尔斯电码",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "sep",
					Aliases: []string{"s"},
					Usage:   "Delimiter 分隔符",
					Value:   "/",
				},
				&cli.StringFlag{
					Name:    "decode",
					Aliases: []string{"d"},
					Usage:   "Decode",
				},
			},
			Action: func(c *cli.Context) (err error) {
				code, err := utils.ArgOrPipe(c.Args().First())
				if err != nil {
					return err
				}
				r, _ := Morse(code, c.Bool("decode"), c.String("sep"))
				_, err = fmt.Println(r)
				return
			},
		},
	},
}
