package domain

import (
	"github.com/mozhu1024/virzz/common"
	"github.com/spf13/cobra"
)

var ctfrCmd = &cobra.Command{
	Use:   "ctfr",
	Short: "滥用证书透明记录 By https://crt.sh",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := ctfr(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Some tools for Domain/SubDomain",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	domainCmd.AddCommand(ctfrCmd)
	domainCmd.SuggestionsMinimumDistance = 1
}

func ExportCommand() []*cobra.Command {
	return []*cobra.Command{
		domainCmd,
	}
}
