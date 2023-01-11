package hash

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/common"
	"golang.org/x/crypto/bcrypt"
)

var bcryptGenerateCmd = &cobra.Command{
	Use:   "gen",
	Short: "Bcrypt Generate",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
			return fmt.Errorf("cost must be between %d and %d", bcrypt.MinCost, bcrypt.MaxCost)
		}
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := bcryptGenerate(s, cost)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}
var bcryptCompareCmd = &cobra.Command{
	Use:   "comp [hashed] [password]",
	Short: "Bcrypt Compare",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := bcryptCompare(args[0], args[1])
		if err != nil {
			return err
		}
		logger.Success("Compare OK")
		return nil
	},
}

var bcryptCmd = &cobra.Command{
	Use:   "bcrypt",
	Short: "Bcrypt Generate/Compare",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var (
	cost int
)

func init() {
	bcryptGenerateCmd.Flags().IntVarP(&cost, "cost", "c", bcrypt.DefaultCost, "bcrypt cost")
	bcryptCmd.AddCommand(
		bcryptGenerateCmd,
		bcryptCompareCmd,
	)
}
