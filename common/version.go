package common

import (
	"fmt"

	"github.com/spf13/cobra"
)

func VersionCommand(name, version, buildID string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version",
		Run: func(cmd *cobra.Command, args []string) {
			if buildID == "0" {
				buildID = "dev"
			}
			fmt.Printf("%s %s build-%s\n", name, version, buildID)
		},
	}
}
