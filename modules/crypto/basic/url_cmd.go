package basic

import (
	"github.com/spf13/cobra"
	"github.com/virzz/virzz/common"
)

// urlencodeCmd
var urlencodeCmd = &cobra.Command{
	Use:   "urle",
	Short: "URL Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
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
		s, err := common.GetFirstArg(args)
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
