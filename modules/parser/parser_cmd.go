package parser

import (
	"github.com/spf13/cobra"
	"github.com/virzz/virzz/common"
)

var parserCmd = &cobra.Command{
	Use:   "parser",
	Short: "Parse some file",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var parserProcNetTcpCmd = &cobra.Command{
	Use:   "tcp",
	Short: "Parse /proc/net/tcp",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := parseProcNetTcp(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

func init() {
	parserCmd.AddCommand(
		// /proc/net/tcp
		parserProcNetTcpCmd,
	)
}

func ExportCommand() []*cobra.Command {
	return []*cobra.Command{parserCmd}
}
