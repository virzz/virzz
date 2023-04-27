package basic

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli/v3"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/utils"
)

var str2asciiCmd = &cli.Command{
	Category: "String",
	Name:     "chr2ord",
	Aliases:  []string{"ords"},
	Usage:    "String -> ASCII",
	Action: func(c *cli.Context) (err error) {
		r, err := StringToASCII(c.Args().First())
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}

var ascii2strCmd = &cli.Command{
	Category: "String",
	Name:     "ord2str",
	Aliases:  []string{"chrs"},
	Usage:    "ASCII -> String",
	Action: func(c *cli.Context) (err error) {
		r, err := ASCIIToString(c.Args().First())
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}

var hex2strCmd = &cli.Command{
	Category: "String",
	Name:     "hex2str",
	Usage:    "Hex -> String",
	Aliases:  []string{"chrs"},
	Action: func(c *cli.Context) (err error) {
		r, err := HexToString(c.Args().First())
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}

var str2hexCmd = &cli.Command{
	Category: "String",
	Name:     "str2hex",
	Usage:    "String -> Hex",
	Action: func(c *cli.Context) (err error) {
		r, err := StringToHex(c.Args().First())
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}
var hex2decCmd = &cli.Command{
	Category: "String",
	Name:     "hex2dec",
	Usage:    "Hex -> Dec",
	Action: func(c *cli.Context) (err error) {
		r, err := HexToDecStr(c.Args().First())
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}

var dec2hexCmd = &cli.Command{
	Category: "String",
	Name:     "dec2hex",
	Usage:    "Dec -> Hex",
	Action: func(c *cli.Context) (err error) {
		r, err := DecToHex(c.Args().First())
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}

var hex2bytesCmd = &cli.Command{
	Category: "String",
	Name:     "hex2bytes",
	Usage:    "Hex -> Bytes String",
	Action: func(c *cli.Context) (err error) {
		r, err := HexToByteString(c.Args().First())
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}

var bytes2hexCmd = &cli.Command{
	Category: "String",
	Name:     "bytes2hex",
	Usage:    "ByteString -> Hex",
	Action: func(c *cli.Context) (err error) {
		r, err := ByteStringToHex(c.Args().First())
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}

var bytes2strCmd = &cli.Command{
	Category: "String",
	Name:     "bytes2str",
	Usage:    "ByteString -> String",
	Action: func(c *cli.Context) (err error) {
		r, err := ByteStringToString(c.Args().First())
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}

var randStrCmd = &cli.Command{
	Category: "String",
	Name:     "randstr",
	Aliases:  []string{"rstr"},
	Usage:    "Generate random string (default 8 chars)",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "regex",
			Aliases: []string{"r"},
			Value:   "a-z0-9",
			Usage:   "Regex for random string allowed chars",
		},
		&cli.BoolFlag{
			Name:    "upper",
			Aliases: []string{"u"},
			Usage:   "Upper Alphabet",
		},
		&cli.BoolFlag{
			Name:    "lowwer",
			Aliases: []string{"l"},
			Usage:   "Lowwer Alphabet",
		},
		&cli.BoolFlag{
			Name:    "digit",
			Aliases: []string{"d"},
			Usage:   "Digit Alphabet",
		},
		&cli.BoolFlag{
			Name:    "hex",
			Aliases: []string{"x"},
			Usage:   "Hex Alphabet",
		},
	},
	Action: func(c *cli.Context) (err error) {
		n := 8
		regex := c.String("regex")
		if c.Bool("upper") {
			regex = "A-Z"
		} else if c.Bool("lowwer") {
			regex = "a-z"
		} else if c.Bool("digit") {
			regex = "0-9"
		} else if c.Bool("hex") {
			regex = "0-9a-f"
		}
		if c.NArg() > 0 {
			_n, err := strconv.Atoi(c.Args().First())
			if err != nil {
				logger.WarnF("Invalid length: %s, use default 8", c.Args().First())
			} else {
				n = _n
			}
		}
		_, err = fmt.Println(utils.RandomStringByLength(n, regex))
		return
	},
}
