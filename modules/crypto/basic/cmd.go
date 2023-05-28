package basic

import (
	"github.com/urfave/cli/v3"
)

var Cmd = &cli.Command{
	Category: "Crypto",
	Name:     "basic",
	Usage:    "Some basic encodings",
	Commands: []*cli.Command{
		// bin
		bin2hexCmd, hex2binCmd,
		// url
		urlencodeCmd, urldecodeCmd,
		// string
		str2asciiCmd, ascii2strCmd,
		hex2strCmd, str2hexCmd,
		hex2decCmd, dec2hexCmd, hex2bytesCmd,
		bytes2hexCmd, bytes2strCmd,
		randStrCmd,
	},
}
