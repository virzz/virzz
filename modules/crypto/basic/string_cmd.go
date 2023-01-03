package basic

import (
	"github.com/spf13/cobra"
	"github.com/virzz/virzz/common"
)

// str2asciiCmd
var str2asciiCmd = &cobra.Command{
	Use:     "chr2ord",
	Aliases: []string{"ords"},
	Short:   "String -> ASCII",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := StringToASCII(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// ascii2strCmd
var ascii2strCmd = &cobra.Command{
	Use:     "ord2str",
	Aliases: []string{"chrs"},
	Short:   "ASCII -> String",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := ASCIIToString(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// hex2strCmd
var hex2strCmd = &cobra.Command{
	Use:   "hex2str",
	Short: "Hex -> String",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := HexToString(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// str2hexCmd
var str2hexCmd = &cobra.Command{
	Use:   "str2hex",
	Short: "String -> Hex",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := StringToHex(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// hex2decCmd
var hex2decCmd = &cobra.Command{
	Use:   "hex2dec",
	Short: "Hex -> Dec",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := HexToDec(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// dec2hexCmd
var dec2hexCmd = &cobra.Command{
	Use:   "dec2hex",
	Short: "Dec -> Hex",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := DecToHex(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// hex2bytesCmd
var hex2bytesCmd = &cobra.Command{
	Use:   "hex2bytes",
	Short: "Hex -> Bytes String",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := HexToByteString(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// bytes2hexCmd
var bytes2hexCmd = &cobra.Command{
	Use:   "bytes2hex",
	Short: "ByteString -> Hex",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := ByteStringToHex(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// bytes2strCmd
var bytes2strCmd = &cobra.Command{
	Use:   "bytes2str",
	Short: "ByteString -> String",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := ByteStringToString(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}
