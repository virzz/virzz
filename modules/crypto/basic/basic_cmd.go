package basic

import (
	"github.com/spf13/cobra"
)

var basicCmd = &cobra.Command{
	Use:   "basic",
	Short: "Some basic encodings",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	basicCmd.AddCommand(
		// string
		str2asciiCmd, ascii2strCmd,
		hex2strCmd, str2hexCmd,
		hex2decCmd, dec2hexCmd, hex2bytesCmd,
		bytes2hexCmd, bytes2strCmd,

		// url
		urlencodeCmd, urldecodeCmd,

		// bin
		bin2hexCmd, hex2binCmd,
	)
}

func ExportCommand() []*cobra.Command {
	return []*cobra.Command{basicCmd}
}
