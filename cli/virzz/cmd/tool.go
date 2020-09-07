package cmd

import (
	"github.com/spf13/cobra"
	"github.com/virink/virzz/tools/bilibili"
)

// toolCmd
var toolCmd = &cobra.Command{
	Use:   "tool",
	Short: "tool",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// bilibiliCmd
var bilibiliCmd = &cobra.Command{
	Use:   "bili [bv/url]",
	Short: "Bilibili Download",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return bilibili.Bilibilis(args...)
	},
}

func init() {
	toolCmd.AddCommand(bilibiliCmd)
	rootCmd.AddCommand(toolCmd)
}
