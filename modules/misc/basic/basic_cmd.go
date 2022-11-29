package basic

import (
	"github.com/spf13/cobra"
)

var (
	safe bool
	enc  string
)

func init() {
	b64eCmd.Flags().BoolVarP(&safe, "safe", "s", false, "Use Safe Encode")
	b58eCmd.Flags().StringVarP(&enc, "enc", "e", "", "Use sepcial EncodeTable: [flickr|ripple]")
}

func ExportCommand() []*cobra.Command {
	return []*cobra.Command{
		// base-x
		b64eCmd, b32eCmd, b58eCmd,
		b64dCmd, b32dCmd, b58dCmd,
		baseXCmd,

		// string
		str2asciiCmd, ascii2strCmd,
		hex2strCmd, str2hexCmd,
		hex2decCmd, dec2hexCmd, hex2bytesCmd,
		bytes2hexCmd, bytes2strCmd,

		// url
		urlencodeCmd, urldecodeCmd,

		// bin
	}
}
