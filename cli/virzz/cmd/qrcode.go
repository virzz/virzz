package cmd

import (
	"github.com/spf13/cobra"
	cm "github.com/virink/virzz/common"
	"github.com/virink/virzz/misc/image"
)

func init() {
	var (
		reverse bool
		dstname string
	)

	// zeroOneToQrcodeCmd -
	var zeroOneToQrcodeCmd = &cobra.Command{
		Use:   "bs2qr",
		Short: "Bin String (0,1) to Qrcode Image",
		// Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			r, err := image.ZeroOneToQrcode(s, reverse, dstname)
			if err != nil {
				return err
			}
			return cm.Output(r, true)
		},
	}

	zeroOneToQrcodeCmd.Flags().BoolVarP(&reverse, "reverse", "r", false, "Reverse color")
	zeroOneToQrcodeCmd.Flags().StringVarP(&dstname, "dstname", "d", "virzz_qrcode.png", "Dest Filename")

	rootCmd.AddCommand(zeroOneToQrcodeCmd)
}
