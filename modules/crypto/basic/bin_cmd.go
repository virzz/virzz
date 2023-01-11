package basic

import (
	"github.com/spf13/cobra"
	"github.com/virzz/virzz/common"
)

// bin2hexCmd
var bin2hexCmd = &cobra.Command{
	Use:   "bin2hex",
	Short: "Bin -> Hex",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := common.GetFileOrPipe(args)
		if err != nil {
			return err
		}
		r, err := BinToHex(data)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// hex2binCmd
var hex2binCmd = &cobra.Command{
	Use:   "hex2bin",
	Short: "Hex -> Bin",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := HexToBin(s)
		if err != nil {
			return err
		}
		return common.OutputBytes(r)
	},
}
