package dsstore

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/virzz/virzz/utils"
)

var (
	save bool = false
)

// dsstoreCmd represents the dsstore command
var dsstoreCmd = &cobra.Command{
	Use:   "dsstore",
	Short: ".DS_Store Parser",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if save {
			return fmt.Errorf("TODO")
		}
		r, err := dsStore(args[0], save)
		if err != nil {
			return err
		}
		return utils.Output(r)
	},
}

func init() {
	dsstoreCmd.Flags().BoolVarP(&save, "save", "s", false, "Save file what were found by url")
}

func ExportCommand() []*cobra.Command {
	return []*cobra.Command{
		dsstoreCmd,
	}
}
