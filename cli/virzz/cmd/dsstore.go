package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	cm "github.com/virink/virzz/common"
	"github.com/virink/virzz/misc/files/dsstore"
)

// dsstoreCmd represents the dsstore command
var dsstoreCmd = &cobra.Command{
	Use:   "dsstore",
	Short: ".DS_Store Parser",
	// Long:  `A Git source leak exploit tool that restores the entire Git repository, including data from stash, for white-box auditing and analysis of developers' mind`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if save {
			return fmt.Errorf("TODO")
		}
		r, err := dsstore.DSStore(args[0], save)
		if err != nil {
			return err
		}
		return cm.Output(r)
	},
}

var (
	save bool = false
)

func init() {
	dsstoreCmd.Flags().BoolVarP(&save, "save", "s", false, "Save file what were found by url")
	rootCmd.AddCommand(dsstoreCmd)
}
