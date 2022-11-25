package basic

import (
	"fmt"
	"strings"

	"github.com/mozhu1024/virzz/common"
	"github.com/mozhu1024/virzz/logger"
	"github.com/spf13/cobra"
)

var (
	safe bool
	enc  string
)

// b64eCmd -
var b64eCmd = &cobra.Command{
	Use:   "b64e",
	Short: "Base64 Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetArgs(args)
		if err != nil {
			return err
		}
		r, err := Base64Encode(s, safe)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b64dCmd -
var b64dCmd = &cobra.Command{
	Use:   "b64d",
	Short: "Base64 Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetArgs(args)
		if err != nil {
			return err
		}
		r, err := Base64Decode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b32eCmd -
var b32eCmd = &cobra.Command{
	Use:   "b32e",
	Short: "Base32 Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetArgs(args)
		if err != nil {
			return err
		}
		r, err := Base32Encode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b32dCmd -
var b32dCmd = &cobra.Command{
	Use:   "b32d",
	Short: "Base64 Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetArgs(args)
		if err != nil {
			return err
		}
		r, err := Base32Decode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b58eCmd -
var b58eCmd = &cobra.Command{
	Use:   "b58e",
	Short: "Base58 Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetArgs(args)
		if err != nil {
			return err
		}
		r, err := Base58Encode(s, strings.ToLower(enc))
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b58dCmd -
var b58dCmd = &cobra.Command{
	Use:   "b58d",
	Short: "Base58 Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetArgs(args)
		if err != nil {
			return err
		}
		r, err := Base58Decode(s, strings.ToLower(enc))
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var baseXCmd = &cobra.Command{
	Use:   "basex",
	Short: "Auto Base-X Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetArgs(args)
		if err != nil {
			return err
		}
		var r string
		if r, err = Base64Decode(s); err != nil {
			if r, err = Base32Decode(s); err != nil {
				if r, err = Base58Decode(s); err != nil {
					return fmt.Errorf("auto decode base-x fail")
				}
				logger.Info("Base58")
			}
			logger.Info("Base32")
		}
		logger.Info("Base64")
		return common.Output(r)
	},
}

// str2asciiCmd
var str2asciiCmd = &cobra.Command{
	Use:     "str2ascii",
	Aliases: []string{"ords"},
	Short:   "String -> ASCII",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetArgs(args)
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
	Use:     "ascii2str",
	Aliases: []string{"chrs"},
	Short:   "String -> ASCII",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetArgs(args)
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
		s, err := common.GetArgs(args)
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
		s, err := common.GetArgs(args)
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
		s, err := common.GetArgs(args)
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
		s, err := common.GetArgs(args)
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
		s, err := common.GetArgs(args)
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
		s, err := common.GetArgs(args)
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
		s, err := common.GetArgs(args)
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

// urlencodeCmd
var urlencodeCmd = &cobra.Command{
	Use:   "urle",
	Short: "URL Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetArgs(args)
		if err != nil {
			return err
		}
		r, err := URLEncode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// urldecodeCmd
var urldecodeCmd = &cobra.Command{
	Use:   "urld",
	Short: "URL Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetArgs(args)
		if err != nil {
			return err
		}
		r, err := URLDecode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

func init() {
	b64eCmd.Flags().BoolVarP(&safe, "safe", "s", false, "Use Safe Encode")
	b58eCmd.Flags().StringVarP(&enc, "enc", "e", "", "Use sepcial EncodeTable: [flickr|ripple]")
}

func BasicCmd() []*cobra.Command {
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
