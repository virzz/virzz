package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/virink/virzz/misc/basic"
)

// b64eCmd -
var b64eCmd = &cobra.Command{
	Use:   "b64e",
	Short: "Base64 Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := basic.Base64Encode(s, safe)
		if err != nil {
			return err
		}
		return output(r)
	},
}

// b64dCmd -
var b64dCmd = &cobra.Command{
	Use:   "b64d",
	Short: "Base64 Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := basic.Base64Decode(s)
		if err != nil {
			return err
		}
		return output(r)
	},
}

// b32eCmd -
var b32eCmd = &cobra.Command{
	Use:   "b32e",
	Short: "Base32 Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := basic.Base32Encode(s)
		if err != nil {
			return err
		}
		return output(r)
	},
}

// b32dCmd -
var b32dCmd = &cobra.Command{
	Use:   "b32d",
	Short: "Base64 Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := basic.Base32Decode(s)
		if err != nil {
			return err
		}
		return output(r)
	},
}

// b58eCmd -
var b58eCmd = &cobra.Command{
	Use:   "b58e",
	Short: "Base58 Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := basic.Base58Encode(s, strings.ToLower(enc))
		if err != nil {
			return err
		}
		return output(r)
	},
}

// b58dCmd -
var b58dCmd = &cobra.Command{
	Use:   "b58d",
	Short: "Base58 Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := basic.Base58Decode(s, strings.ToLower(enc))
		if err != nil {
			return err
		}
		return output(r)
	},
}

// TODO: baseXCmd
// var baseXCmd = &cobra.Command{
// 	Use:   "basex",
// 	Short: "Auto Base-X Decode",
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		return nil
// 	},
// }

var safe bool
var enc string

func init() {
	b64eCmd.Flags().BoolVarP(&safe, "safe", "s", false, "Use Safe Encode")
	b58eCmd.Flags().StringVarP(&enc, "enc", "e", "", "Use sepcial EncodeTable: [flickr|ripple]")
	rootCmd.AddCommand(b64eCmd)
	rootCmd.AddCommand(b64dCmd)
	rootCmd.AddCommand(b32eCmd)
	rootCmd.AddCommand(b32dCmd)
	rootCmd.AddCommand(b58eCmd)
	rootCmd.AddCommand(b58dCmd)
	// rootCmd.AddCommand(baseXCmd)
}
