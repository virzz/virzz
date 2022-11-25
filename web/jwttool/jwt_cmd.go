package jwttool

import (
	"os"

	"github.com/mozhu1024/virzz/common"
	"github.com/mozhu1024/virzz/logger"
	"github.com/spf13/cobra"
)

var jwtPrintCmd = &cobra.Command{
	Use:   "jwt",
	Short: "JWT Print",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetArgs(args)
		if err != nil {
			return err
		}
		secret = getSecret(secret)
		r, err := printJWT(s, secret)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var jwtCrackCmd = &cobra.Command{
	Use:   "jwtc",
	Short: "JWT Crack",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetArgs(args)
		if err != nil {
			return err
		}
		r, err := crackJWT(s, minLen, maxLen, alphabet, prefix, suffix)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var jwtModifyCmd = &cobra.Command{
	Use:   "jwtm",
	Short: "JWT Modify",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetArgs(args)
		if err != nil {
			return err
		}
		secret = getSecret(secret)
		r, err := modifyJWT(s, none, secret, claims, method)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var (
	minLen   = 4
	maxLen   = 4
	alphabet = "abcdefghijklnmopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	prefix   = ""
	suffix   = ""
	secret   = ""
	none     = false
	claims   map[string]string
	method   = "HS256"
)

func getSecret(s string) string {
	f, err := os.Stat(s)
	if err == nil && !f.IsDir() && f.Size() > 0 {
		data, err := os.ReadFile(s)
		if err == nil {
			logger.Debug("secret", string(data))
			return string(data)
		}
	}
	return s
}

var JWTToolCmd = &cobra.Command{
	Use:   "jwttool",
	Short: "A jwt tool with Print/Crack/Modify",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	jwtPrintCmd.Flags().StringVarP(&secret, "secret", "s", "", "the secret")

	jwtCrackCmd.Flags().IntVarP(&minLen, "min", "m", 4, "the min length secret for crack")
	jwtCrackCmd.Flags().IntVarP(&maxLen, "max", "l", 4, "the max length secret for crack")
	jwtCrackCmd.Flags().StringVarP(&alphabet, "alphabet", "a", alphabet, "the alphabet for the brute")
	jwtCrackCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "prefixed to the secret")
	jwtCrackCmd.Flags().StringVarP(&suffix, "suffix", "s", "", "suffixed to the secret")

	jwtModifyCmd.Flags().BoolVarP(&none, "none", "n", false, "set none method and no signature")
	jwtModifyCmd.Flags().StringVarP(&secret, "secret", "s", "", "the secret")
	jwtModifyCmd.Flags().StringVarP(&method, "method", "m", method, "set method")
	jwtModifyCmd.Flags().StringToStringVarP(&claims, "claims", "c", claims, "modify or add claims")

	JWTToolCmd.AddCommand(jwtPrintCmd, jwtCrackCmd, jwtModifyCmd)
	JWTToolCmd.SuggestionsMinimumDistance = 1
}
