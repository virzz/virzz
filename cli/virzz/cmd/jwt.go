package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	cm "github.com/virink/virzz/common"
	"github.com/virink/virzz/web/jwt"
)

func getSecret(s string) string {
	f, err := os.Stat(s)
	if err == nil && !f.IsDir() && f.Size() > 0 {
		data, err := ioutil.ReadFile(s)
		if err == nil {
			if cm.DebugMode {
				fmt.Fprintln(os.Stderr, "secret", string(data))
			}
			return string(data)
		}
	}
	return s
}

func init() {

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

	// printCmd
	var printCmd = &cobra.Command{
		Use:   "jwt",
		Short: "JWT Print",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			secret = getSecret(secret)
			r, err := jwt.PrintJWT(s, secret)
			if err != nil {
				return err
			}
			return cm.Output(r)
		},
	}

	// crackCmd
	var crackCmd = &cobra.Command{
		Use:   "jwtc",
		Short: "JWT Crack",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			r, err := jwt.CrackJWT(s, minLen, maxLen, alphabet, prefix, suffix)
			if err != nil {
				return err
			}
			return cm.Output(r)
		},
	}

	// modifyCmd
	var modifyCmd = &cobra.Command{
		Use:   "jwtm",
		Short: "JWT Modify",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			secret = getSecret(secret)
			r, err := jwt.ModifyJWT(s, none, secret, claims, method)
			if err != nil {
				return err
			}
			return cm.Output(r)
		},
	}

	printCmd.Flags().StringVarP(&secret, "secret", "s", "", "the secret")

	crackCmd.Flags().IntVarP(&minLen, "min", "m", 4, "the min length secret for crack")
	crackCmd.Flags().IntVarP(&maxLen, "max", "l", 4, "the max length secret for crack")
	crackCmd.Flags().StringVarP(&alphabet, "alphabet", "a", alphabet, "the alphabet for the brute")
	crackCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "prefixed to the secret")
	crackCmd.Flags().StringVarP(&suffix, "suffix", "s", "", "suffixed to the secret")

	modifyCmd.Flags().BoolVarP(&none, "none", "n", false, "set none method and no signature")
	modifyCmd.Flags().StringVarP(&secret, "secret", "s", "", "the secret")
	modifyCmd.Flags().StringVarP(&method, "method", "m", method, "set method")
	modifyCmd.Flags().StringToStringVarP(&claims, "claims", "c", claims, "modify or add claims")

	rootCmd.AddCommand(printCmd, crackCmd, modifyCmd)
}
