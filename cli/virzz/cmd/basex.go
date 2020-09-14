package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	cm "github.com/virink/virzz/common"
	"github.com/virink/virzz/misc/basic"
)

func init() {
	var (
		safe bool
		enc  string
	)
	// b64eCmd -
	var b64eCmd = &cobra.Command{
		Use:   "b64e",
		Short: "Base64 Encode",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			r, err := basic.Base64Encode(s, safe)
			if err != nil {
				return err
			}
			return cm.Output(r)
		},
	}

	// b64dCmd -
	var b64dCmd = &cobra.Command{
		Use:   "b64d",
		Short: "Base64 Decode",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			r, err := basic.Base64Decode(s)
			if err != nil {
				return err
			}
			return cm.Output(r)
		},
	}

	// b32eCmd -
	var b32eCmd = &cobra.Command{
		Use:   "b32e",
		Short: "Base32 Encode",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			r, err := basic.Base32Encode(s)
			if err != nil {
				return err
			}
			return cm.Output(r)
		},
	}

	// b32dCmd -
	var b32dCmd = &cobra.Command{
		Use:   "b32d",
		Short: "Base64 Decode",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			r, err := basic.Base32Decode(s)
			if err != nil {
				return err
			}
			return cm.Output(r)
		},
	}

	// b58eCmd -
	var b58eCmd = &cobra.Command{
		Use:   "b58e",
		Short: "Base58 Encode",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			r, err := basic.Base58Encode(s, strings.ToLower(enc))
			if err != nil {
				return err
			}
			return cm.Output(r)
		},
	}

	// b58dCmd -
	var b58dCmd = &cobra.Command{
		Use:   "b58d",
		Short: "Base58 Decode",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			r, err := basic.Base58Decode(s, strings.ToLower(enc))
			if err != nil {
				return err
			}
			return cm.Output(r)
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

	b64eCmd.Flags().BoolVarP(&safe, "safe", "s", false, "Use Safe Encode")
	b58eCmd.Flags().StringVarP(&enc, "enc", "e", "", "Use sepcial EncodeTable: [flickr|ripple]")
	rootCmd.AddCommand(b64eCmd, b32eCmd, b58eCmd)
	rootCmd.AddCommand(b64dCmd, b32dCmd, b58dCmd)
	// rootCmd.AddCommand(baseXCmd)
}
