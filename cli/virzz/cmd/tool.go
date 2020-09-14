package cmd

import (
	"github.com/spf13/cobra"
	"github.com/virink/virzz/tools/bilibili"
)

func init() {
	// bilibiliCmd
	var bilibiliCmd = &cobra.Command{
		Use:   "bilibili [bv/url]",
		Short: "Download Bilibili video By bv/av/url",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return bilibili.Bilibilis(args...)
		},
	}

	rootCmd.AddCommand(bilibiliCmd)
}
