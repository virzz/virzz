package hashpow

import (
	"github.com/mozhu1024/virzz/common"
	"github.com/spf13/cobra"
)

var (
	pos                          int
	code, prefix, suffix, method string
)

var hashPowCmd = &cobra.Command{
	Use:   "hashpow",
	Short: "A tool for ctfer which make hash collision faster",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.ValidateRequiredFlags(); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return common.Output(doBrute(code, prefix, suffix, method, pos))
	},
}

func init() {
	hashPowCmd.Flags().IntVarP(&pos, "pos", "i", 0, "starting position of hash")
	hashPowCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "text prefix")
	hashPowCmd.Flags().StringVarP(&suffix, "suffix", "s", "", "text suffix")
	hashPowCmd.Flags().StringVarP(&code, "code", "c", "", "part of hash code")
	hashPowCmd.Flags().StringVarP(&method, "hash", "t", "md5", "hash type : md5 sha1")

	hashPowCmd.MarkFlagRequired("code")

	hashPowCmd.SuggestionsMinimumDistance = 1
}

func ExportCommand() []*cobra.Command {
	return []*cobra.Command{
		hashPowCmd,
	}
}
