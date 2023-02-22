package basic

import (
	"fmt"

	"github.com/urfave/cli/v3"
	"github.com/virzz/virzz/utils"
)

var bin2hexCmd = &cli.Command{
	Category: "Bin",
	Name:     "bin2hex",
	Usage:    "Bin -> Hex",
	Action: func(c *cli.Context) (err error) {
		data, err := utils.GetFileOrPipe(c.Args().Slice())
		if err != nil {
			return err
		}
		r, err := BinToHex(data)
		if err != nil {
			return err
		}
		_, err = fmt.Println(r)
		return
	},
}

var hex2binCmd = &cli.Command{
	Category: "Bin",
	Name:     "hex2bin",
	Usage:    "Hex -> Bin",
	Action: func(c *cli.Context) (err error) {
		data, err := utils.ArgOrPipe(c.Args().First())
		if err != nil {
			return err
		}
		r, err := HexToBin(data)
		if err != nil {
			return err
		}
		_, err = fmt.Println(string(r))
		return
	},
}
