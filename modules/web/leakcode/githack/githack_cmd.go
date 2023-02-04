package githack

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/go-playground/validator.v9"
)

var (
	limit   int64
	delay   int64
	timeout int64
)

var GithackCmd = &cobra.Command{
	Use:   "githack",
	Short: "A `.git` folder disclosure exploit",
	Long:  `A Git source leak exploit tool that restores the entire Git repository, including data from stash, for white-box auditing and analysis of developers' mind`,
	Args: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) < 1 {
			return errors.New("requires a url argument")
		}
		if err = validator.New().Var(args[0], "url"); err != nil {
			return
		}
		if !strings.HasPrefix(args[0], "http") {
			return errors.New("must be http(s)")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return gitHack(args[0], limit, delay)
	},
}

func init() {
	GithackCmd.Flags().Int64VarP(&limit, "limit", "l", 10, "Request limit (N times one second)")
	GithackCmd.Flags().Int64VarP(&delay, "delay", "d", 0, "Request delay (N times one second)")
	GithackCmd.Flags().Int64VarP(&timeout, "timeout", "t", 10, "Request timeout (second)")
}

func ExportCommand() []*cobra.Command {
	return []*cobra.Command{
		GithackCmd,
	}
}
