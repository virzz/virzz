package hash

import (
	"github.com/mozhu1024/virzz/common"
	"github.com/spf13/cobra"
)

var (
	pos                        int
	code, prefix, suffix, hash string
)

var HashPowCmd = &cobra.Command{
	Use:   "hashpow",
	Short: "A tool for ctfer which make hash collision faster",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.ValidateRequiredFlags(); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return common.Output(doBrute(code, prefix, suffix, hash, pos))
	},
}

func init() {
	HashPowCmd.Flags().IntVarP(&pos, "pos", "i", 0, "starting position of hash")
	HashPowCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "text prefix")
	HashPowCmd.Flags().StringVarP(&suffix, "suffix", "s", "", "text suffix")
	HashPowCmd.Flags().StringVarP(&code, "code", "c", "", "part of hash code")
	HashPowCmd.Flags().StringVarP(&hash, "hash", "t", "md5", "hash type : md5 sha1")

	HashPowCmd.MarkFlagRequired("code")

	HashPowCmd.SuggestionsMinimumDistance = 1
}
