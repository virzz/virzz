package common

import (
	"os"

	"github.com/spf13/cobra"
)

func CompletionCommand() *cobra.Command {
	return &cobra.Command{
		Use:       "completion [bash|zsh]",
		Short:     "Generate completion script",
		Hidden:    true,
		ValidArgs: []string{"bash", "zsh"},
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				cmd.Root().GenZshCompletion(os.Stdout)
			}
		},
	}
}
