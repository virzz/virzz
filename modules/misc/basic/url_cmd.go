package basic

import (
	"github.com/mozhu1024/virzz/common"
	"github.com/spf13/cobra"
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
