package cmd

import (
	"github.com/spf13/cobra"
	"github.com/virink/virzz/misc/basic"
)

// str2asciiCmd
var str2asciiCmd = &cobra.Command{
	Use:   "str2ascii",
	Short: "String -> ASCII",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := basic.StringToASCII(s)
		if err != nil {
			return err
		}
		return output(r)
	},
}

// ascii2strCmd
var ascii2strCmd = &cobra.Command{
	Use:   "ascii2str",
	Short: "String -> ASCII",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := basic.ASCIIToString(s)
		if err != nil {
			return err
		}
		return output(r)
	},
}

// hex2strCmd
var hex2strCmd = &cobra.Command{
	Use:   "hex2str",
	Short: "Hex -> String",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := basic.ASCIIToString(s)
		if err != nil {
			return err
		}
		return output(r)
	},
}

// str2hexCmd
var str2hexCmd = &cobra.Command{
	Use:   "str2hex",
	Short: "String -> Hex",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := basic.StringToHex(s)
		if err != nil {
			return err
		}
		return output(r)
	},
}

// hex2decCmd
var hex2decCmd = &cobra.Command{
	Use:   "hex2dec",
	Short: "Hex -> Dec",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := basic.HexToDec(s)
		if err != nil {
			return err
		}
		return output(r)
	},
}

// dec2hexCmd
var dec2hexCmd = &cobra.Command{
	Use:   "dec2hex",
	Short: "Dec -> Hex",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := basic.DecToHex(s)
		if err != nil {
			return err
		}
		return output(r)
	},
}

// hex2bytesCmd
var hex2bytesCmd = &cobra.Command{
	Use:   "hex2bytes",
	Short: "Hex -> Bytes String",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := basic.HexToByteString(s)
		if err != nil {
			return err
		}
		return output(r)
	},
}

// bytes2hexCmd
var bytes2hexCmd = &cobra.Command{
	Use:   "bytes2hex",
	Short: "ByteString -> Hex",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := basic.ByteStringToHex(s)
		if err != nil {
			return err
		}
		return output(r)
	},
}

// bytes2strCmd
var bytes2strCmd = &cobra.Command{
	Use:   "bytes2str",
	Short: "ByteString -> String",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := basic.ByteStringToString(s)
		if err != nil {
			return err
		}
		return output(r)
	},
}

// urlencodeCmd
var urlencodeCmd = &cobra.Command{
	Use:   "urle",
	Short: "URL Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := basic.URLEncode(s)
		if err != nil {
			return err
		}
		return output(r)
	},
}

// urldecodeCmd
var urldecodeCmd = &cobra.Command{
	Use:   "urld",
	Short: "URL Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := basic.URLDecode(s)
		if err != nil {
			return err
		}
		return output(r)
	},
}

// // stringxCmd
// var stringxCmd = &cobra.Command{
// 	Use:   "stringx",
// 	Short: "A brief description of your command",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		cmd.Help()
// 	},
// }

func init() {
	rootCmd.AddCommand(str2asciiCmd)
	rootCmd.AddCommand(ascii2strCmd)
	rootCmd.AddCommand(hex2strCmd)
	rootCmd.AddCommand(str2hexCmd)
	rootCmd.AddCommand(hex2decCmd)
	rootCmd.AddCommand(dec2hexCmd)
	rootCmd.AddCommand(hex2bytesCmd)
	rootCmd.AddCommand(bytes2hexCmd)
	rootCmd.AddCommand(bytes2strCmd)

	rootCmd.AddCommand(urlencodeCmd)
	rootCmd.AddCommand(urldecodeCmd)
	// rootCmd.AddCommand(stringxCmd)
}
