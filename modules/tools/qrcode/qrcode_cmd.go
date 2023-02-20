package qrcode

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/virzz/virzz/utils"
)

var (
	exchange bool   = false
	output   string = "-"
)

var qrcodeCmd = &cobra.Command{
	Use:   "qrcode",
	Short: "A qrcode tool for terminal",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var zeroOneCmd = &cobra.Command{
	Use:   "bs",
	Short: "Bin String (0,1) to Qrcode Image",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := utils.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := zeroOneToQrcode(s, exchange, output)
		if err != nil {
			return err
		}
		return utils.Output(r)
	},
}

var parseQrcodeCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse qrcode image",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := utils.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := parseQrcode(s)
		if err != nil {
			return err
		}
		return utils.Output(r)
	},
}

var generateQrcodeCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate qrcode image",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := utils.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := generateQrcode(s, output)
		if err != nil {
			return err
		}
		return utils.Output(r)
	},
}

func init() {
	flagSet := &pflag.FlagSet{}
	flagSet.StringVarP(&output, "output", "o", "-", "To terminal:'-' / File:'Filename'")

	qrcodeCmd.PersistentFlags().AddFlagSet(flagSet)

	generateQrcodeCmd.Flags().AddFlagSet(flagSet)

	zeroOneCmd.Flags().AddFlagSet(flagSet)
	zeroOneCmd.Flags().BoolVarP(&exchange, "exchange", "e", false, "Exchange 0/1")

	qrcodeCmd.AddCommand(zeroOneCmd, parseQrcodeCmd, generateQrcodeCmd)
	qrcodeCmd.SuggestionsMinimumDistance = 1
}

func ExportCommand() []*cobra.Command {
	return []*cobra.Command{
		qrcodeCmd,
	}
}
