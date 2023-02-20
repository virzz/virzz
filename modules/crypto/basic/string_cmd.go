package basic

import (
	"fmt"

	"github.com/urfave/cli/v3"
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
		r, err := HexToDec(c.Args().First())
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
