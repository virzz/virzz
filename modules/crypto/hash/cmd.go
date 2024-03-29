package hash

import (
	"fmt"

	"github.com/urfave/cli/v3"
	"github.com/virzz/virzz/utils"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "hmac",
		Aliases: []string{"s"},
		Usage:   "Hmac Key",
		Value:   "",
	},
	&cli.BoolFlag{
		Name:    "raw",
		Usage:   "Print raw data",
		Aliases: []string{"r"},
	},
}

var Cmd = &cli.Command{
	Category: "Crypto",
	Name:     "hash",
	Usage:    "Hash Function",
	Commands: []*cli.Command{
		// md5
		&cli.Command{
			Category: "Hash",
			Name:     "md5",
			Usage:    "MD5 algorithm",
			Flags:    flags,
			Action: func(c *cli.Context) (err error) {
				code, err := utils.GetArgFilePipe(c.Args().First())
				if err != nil {
					return err
				}
				key := c.String("hmac")
				var r string
				if key != "" {
					r, _ = HmacMDHash(code, []byte(key), 5)
				} else {
					r, _ = MDHash(code, 5, c.Bool("raw"))
				}
				if c.Bool("raw") {
					_, err = fmt.Print(r)
					return
				}
				_, err = fmt.Println(r)
				return
			},
		},
		// md4
		&cli.Command{
			Category: "Hash",
			Name:     "md4",
			Usage:    "MD4 algorithm",
			Flags:    flags,
			Action: func(c *cli.Context) (err error) {
				code, err := utils.GetArgFilePipe(c.Args().First())
				if err != nil {
					return err
				}
				key := c.String("hmac")
				var r string
				if key != "" {
					r, _ = HmacMDHash(code, []byte(key), 4)
				} else {
					r, _ = MDHash(code, 4, c.Bool("raw"))
				}
				if c.Bool("raw") {
					_, err = fmt.Print(r)
					return
				}
				_, err = fmt.Println(r)
				return
			},
		},
		// md2
		&cli.Command{
			Category: "Hash",
			Name:     "md2",
			Usage:    "MD2 algorithm",
			Flags:    flags,
			Action: func(c *cli.Context) (err error) {
				code, err := utils.GetArgFilePipe(c.Args().First())
				if err != nil {
					return err
				}
				key := c.String("hmac")
				var r string
				if key != "" {
					r, _ = HmacMDHash(code, []byte(key), 2)
				} else {
					r, _ = MDHash(code, 2, c.Bool("raw"))
				}
				if c.Bool("raw") {
					_, err = fmt.Print(r)
					return
				}
				_, err = fmt.Println(r)
				return
			},
		},

		// sha1
		&cli.Command{
			Category: "Hash",
			Name:     "sha1",
			Usage:    "SHA1 algorithm",
			Flags:    flags,
			Action: func(c *cli.Context) (err error) {
				code, err := utils.GetArgFilePipe(c.Args().First())
				if err != nil {
					return err
				}
				key := c.String("hmac")
				var r string
				if key != "" {
					r, _ = HmacSha1Hash(code, []byte(key))
				} else {
					r, _ = Sha1Hash(code, c.Bool("raw"))
				}
				if c.Bool("raw") {
					_, err = fmt.Print(r)
					return
				}
				_, err = fmt.Println(r)
				return
			},
		},

		// sha2
		&cli.Command{
			Category: "Hash",
			Name:     "sha2",
			Usage:    "SHA2 224|256|384|512|512224|512256",
			Flags: append(flags,
				&cli.IntFlag{
					Name:    "type",
					Aliases: []string{"t"},
					Usage:   "Type of hash",
					Value:   256,
					Action: func(c *cli.Context, t int) error {
						if t == 224 || t == 256 || t == 384 || t == 512 ||
							t == 512224 || t == 512256 {
							return nil
						}
						return fmt.Errorf("invalid type: %d", t)
					},
				},
			),
			Action: func(c *cli.Context) (err error) {
				code, err := utils.GetArgFilePipe(c.Args().First())
				if err != nil {
					return err
				}
				key := c.String("hmac")
				var r string
				if key != "" {
					r, _ = HmacSha2Hash(code, []byte(key), c.Int("type"))
				} else {
					r, _ = Sha2Hash(code, c.Int("type"))
				}
				_, err = fmt.Println(r)
				return
			},
		},

		// sha3
		&cli.Command{
			Category: "Hash",
			Name:     "sha3",
			Usage:    "SHA3 224|256|384|512",
			Flags: append(flags,
				&cli.IntFlag{
					Name:    "type",
					Aliases: []string{"t"},
					Usage:   "Type of hash",
					Value:   256,
					Action: func(c *cli.Context, s int) error {
						if s == 224 || s == 256 || s == 384 || s == 512 {
							return nil
						}
						return fmt.Errorf("invalid size: %d", s)
					},
				},
			),
			Action: func(c *cli.Context) (err error) {
				code, err := utils.GetArgFilePipe(c.Args().First())
				if err != nil {
					return err
				}
				key := c.String("hmac")
				var r string
				if key != "" {
					r, _ = HmacSha3Hash(code, []byte(key), c.Int("type"))
				} else {
					r, _ = Sha3Hash(code, c.Int("type"))
				}
				_, err = fmt.Println(r)
				return
			},
		},

		// ripemd160
		&cli.Command{
			Category: "Hash",
			Name:     "ripemd",
			Aliases:  []string{"ripemd160"},
			Usage:    "RIPEMD160 algorithm",
			Flags: append(flags,
				&cli.IntFlag{
					Name:    "type",
					Aliases: []string{"t"},
					Usage:   "Type of hash",
					Value:   256,
					Action: func(c *cli.Context, s int) error {
						if s == 224 || s == 256 || s == 384 || s == 512 {
							return nil
						}
						return fmt.Errorf("invalid size: %d", s)
					},
				},
			),
			Action: func(c *cli.Context) (err error) {
				code, err := utils.GetArgFilePipe(c.Args().First())
				if err != nil {
					return err
				}
				key := c.String("hmac")
				var r string
				if key != "" {
					r, _ = HmacRipemd160Hash(code, []byte(key))
				} else {
					r, _ = Ripemd160Hash(code)
				}
				_, err = fmt.Println(r)
				return
			},
		},

		// mysql
		&cli.Command{
			Category: "Hash",
			Name:     "mysql",
			Usage:    "MySQL Hash password using before 4.1",
			Action: func(c *cli.Context) (err error) {
				code, err := utils.GetArgFilePipe(c.Args().First())
				if err != nil {
					return err
				}
				_, err = fmt.Println(MySQLHash(code))
				return
			},
		},

		// mysql5
		&cli.Command{
			Category: "Hash",
			Name:     "mysql5",
			Usage:    "MySQL5 Hash password using 4.1+ method (SHA1)",
			Action: func(c *cli.Context) (err error) {
				code, err := utils.GetArgFilePipe(c.Args().First())
				if err != nil {
					return err
				}
				_, err = fmt.Println(MySQL5Hash(code))
				return
			},
		},

		// ntlm
		&cli.Command{
			Category: "Hash",
			Name:     "ntlm",
			Aliases:  []string{"ntlm"},
			Usage:    "NTLM Hash password (MD4(utf16))",
			Action: func(c *cli.Context) (err error) {
				code, err := utils.GetArgFilePipe(c.Args().First())
				if err != nil {
					return err
				}
				_, err = fmt.Println(NTLMv1Hash(code))
				return
			},
		},
	},
}
