package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	var completionCmd = &cobra.Command{
		Use:                   "completion [bash|zsh]",
		Short:                 "Generate completion script",
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh"},
		Args:                  cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				cmd.Root().GenZshCompletion(os.Stdout)
			}
		},
	}

	rootCmd.AddCommand(completionCmd)
}
