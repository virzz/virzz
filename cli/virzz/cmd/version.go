package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/virink/virzz/common"
)

func printVersion() {
	fmt.Println(fmt.Sprintf("%s %s", common.AppName, common.Version))
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
