package basex

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v3"
	"github.com/virzz/virzz/utils"
)

var b16eCmd = &cli.Command{
	Category: "String",
	Name:     "b16e",
	Usage:    "Base16 Encode",
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, _ := Base16Encode(code)
		_, err = fmt.Println(r)
		return
	},
}

var b16dCmd = &cli.Command{
	Category: "String",
	Name:     "b16d",
	Usage:    "Base16 Decode",
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, err := Base16Decode(code)
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}

var b32eCmd = &cli.Command{
	Category: "String",
	Name:     "b32e",
	Usage:    "Base32 Encode",
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, _ := Base32Encode(code)
		_, err = fmt.Println(r)
		return
	},
}

var b32dCmd = &cli.Command{
	Category: "String",
	Name:     "b32d",
	Usage:    "Base32 Decode",
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, err := Base32Decode(code)
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}

var b36eCmd = &cli.Command{
	Category: "String",
	Name:     "b36e",
	Usage:    "Base36 Encode",
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, _ := Base36Encode(code)
		_, err = fmt.Println(r)
		return
	},
}

var b36dCmd = &cli.Command{
	Category: "String",
	Name:     "b36d",
	Usage:    "Base36 Decode",
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, err := Base36Decode(code)
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}

var base58EncFlag = &cli.StringFlag{
	Name:    "enc",
	Aliases: []string{"e"},
	Usage:   "Base58 encoder <|flickr|ripple|bitcoin>",
	// Action: func(c *cli.Context, enc string) error {
	// 	if enc == "" || enc == "flickr" || enc == "ripple" || enc == "bitcoin" {
	// 		return nil
	// 	}
	// 	return fmt.Errorf("unknown encoder : %s", enc)
	// },
}

var b58eCmd = &cli.Command{
	Category: "String",
	Name:     "b58e",
	Usage:    "Base58 Encode",
	Flags: []cli.Flag{
		base58EncFlag,
	},
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, _ := Base58Encode(code, strings.ToLower(c.String("enc")))
		_, err = fmt.Println(r)
		return
	},
}

var b58dCmd = &cli.Command{
	Category: "String",
	Name:     "b58d",
	Usage:    "Base58 Decode",
	Flags:    []cli.Flag{base58EncFlag},
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, err := Base58Decode(code, strings.ToLower(c.String("enc")))
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}

var b62eCmd = &cli.Command{
	Category: "String",
	Name:     "b62e",
	Usage:    "Base62 Encode",
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, _ := Base62Encode(code)
		_, err = fmt.Println(r)
		return
	},
}

var b62dCmd = &cli.Command{
	Category: "String",
	Name:     "b62d",
	Usage:    "Base62 Decode",
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, err := Base62Decode(code)
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}

var b64eCmd = &cli.Command{
	Category: "String",
	Name:     "b64e",
	Usage:    "Base64 Encode",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "safe",
			Aliases: []string{"url", "s"},
			Usage:   "Base64 URLEncoding (defined in RFC 4648)",
		},
	},
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, _ := Base64Encode(code, c.Bool("safe"))
		_, err = fmt.Println(r)
		return
	},
}

var b64dCmd = &cli.Command{
	Category: "String",
	Name:     "b64d",
	Usage:    "Base64 Decode",
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, err := Base64Decode(code)
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}

var b85eCmd = &cli.Command{
	Category: "String",
	Name:     "b85e",
	Usage:    "Base85 Encode",
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, _ := Base85Encode(code)
		_, err = fmt.Println(r)
		return
	},
}

var b85dCmd = &cli.Command{
	Category: "String",
	Name:     "b85d",
	Usage:    "Base85 Decode",
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, err := Base85Decode(code)
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}

var b91eCmd = &cli.Command{
	Category: "String",
	Name:     "b91e",
	Usage:    "Base91 Encode",
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, _ := Base91Encode(code)
		_, err = fmt.Println(r)
		return
	},
}

var b91dCmd = &cli.Command{
	Category: "String",
	Name:     "b91d",
	Usage:    "Base91 Decode",
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, err := Base91Decode(code)
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}

var b92eCmd = &cli.Command{
	Category: "String",
	Name:     "b92e",
	Usage:    "Base92 Encode",
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, _ := Base92Encode(code)
		_, err = fmt.Println(r)
		return
	},
}

var b92dCmd = &cli.Command{
	Category: "String",
	Name:     "b92d",
	Usage:    "Base92 Decode",
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, err := Base92Decode(code)
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}

var b100eCmd = &cli.Command{
	Category: "String",
	Name:     "b100e",
	Usage:    "Base100 Encode",
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, _ := Base100Encode(code)
		_, err = fmt.Println(r)
		return
	},
}

var b100dCmd = &cli.Command{
	Category: "String",
	Name:     "b100d",
	Usage:    "Base100 Decode",
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, err := Base100Decode(code)
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}

var baseNCmd = &cli.Command{
	Category: "String",
	Name:     "auto",
	Usage:    "Auto Base-N Decode",
	Action: func(c *cli.Context) (err error) {
		code, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		var r string
		var res []string
		if r, err = Base16Decode(code); err == nil {
			res = append(res, fmt.Sprintf("base16          : %s", r))
		}
		if r, err = Base32Decode(code); err == nil {
			res = append(res, fmt.Sprintf("base32          : %s", r))
		}
		if r, err = Base36Decode(code); err == nil {
			res = append(res, fmt.Sprintf("base36          : %s", r))
		}
		for _, enc := range []string{"flickr", "ripple", "bitcoin"} {
			if r, err = Base58Decode(code, enc); err == nil {
				res = append(res, fmt.Sprintf("base58(%7s) : %s", enc, r))
			}
		}
		if r, err = Base62Decode(code); err == nil {
			res = append(res, fmt.Sprintf("base62          : %s", r))
		}
		if r, err = Base64Decode(code); err == nil {
			res = append(res, fmt.Sprintf("base64          : %s", r))
		}
		if r, err = Base85Decode(code); err == nil {
			res = append(res, fmt.Sprintf("base85          : %s", r))
		}
		if r, err = Base91Decode(code); err == nil {
			res = append(res, fmt.Sprintf("base91          : %s", r))
		}
		if r, err = Base92Decode(code); err == nil {
			res = append(res, fmt.Sprintf("base92          : %s", r))
		}
		if r, err = Base100Decode(code); err == nil {
			res = append(res, fmt.Sprintf("base100         : %s", r))
		}
		// for i, c := range res {
		// 	res[i] = strings.ReplaceAll(res[i], "\r", `\\r`)
		// 	res[i] = strings.ReplaceAll(res[i], "\n", `\\n`)
		// 	r, _ := basic.StringToASCII(c)
		// 	res[i] += r
		// }
		_, err = fmt.Println(strings.Join(res, "\n"))
		return
	},
}

var Cmd = &cli.Command{
	Category: "Crypto",
	Name:     "basex",
	Usage:    "Base 16/32/58/62/64/85/91/92/100 Encode/Decode",
	Commands: []*cli.Command{
		b16eCmd, b16dCmd,
		b32eCmd, b32dCmd,
		b36eCmd, b36dCmd,
		b58eCmd, b58dCmd,
		b62eCmd, b62dCmd,
		b64eCmd, b64dCmd,
		b85eCmd, b85dCmd,
		b91eCmd, b91dCmd,
		b92eCmd, b92dCmd,
		b100eCmd, b100dCmd,
		baseNCmd,
	},
}
