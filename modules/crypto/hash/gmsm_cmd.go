package hash

import (
	"github.com/spf13/cobra"
	"github.com/virzz/virzz/utils"
)

var sm3Cmd = &cobra.Command{
	Use:   "sm3",
	Short: "SM3 hash algorithm",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := utils.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := sm3Hash(s)
		if err != nil {
			return err
		}
		return utils.Output(r)
	},
}

// var gmsmCmd = &cobra.Command{
// 	Use:   "gmsm",
// 	Short: "Some gmsm function",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		cmd.Help()
// 	},
// }

// func init() {
// 	gmsmCmd.AddCommand(
// 		sm3Cmd,
// 	)
// }
