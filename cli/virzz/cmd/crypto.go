package cmd

import (
	"github.com/spf13/cobra"
	"github.com/virink/virzz/crypto"
	"github.com/virink/virzz/common"
)

// caesarCmd
var caesarCmd = &cobra.Command{
	Use:   "caesar",
	Short: "Caesar Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetArgs(args)
		if err != nil {
			return err
		}
		r, _ := crypto.Caesar(s)
		return common.Output(r)
	},
}

// rot13Cmd
var rot13Cmd = &cobra.Command{
	Use:   "rot13",
	Short: "Rot13 By Caesar Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetArgs(args)
		if err != nil {
			return err
		}
		r, _ := crypto.Rot13(s)
		return common.Output(r)
	},
}

// morseCmd
var morseCmd = &cobra.Command{
	Use:   "morse",
	Short: "Morse Code 摩尔斯电码",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetArgs(args)
		if err != nil {
			return err
		}
		r, err := crypto.Morse(s, decode, sep)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// morseDecodeCmd
var morseDecodeCmd = &cobra.Command{
	Use:   "morsed",
	Short: "Morse Code Decode 摩尔斯电码",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetArgs(args)
		if err != nil {
			return err
		}
		r, err := crypto.Morse(s, true, sep)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var (
	decode bool
	sep    string
)

func init() {
	morseCmd.Flags().BoolVarP(&decode, "decode", "d", false, "Decode")
	morseCmd.Flags().StringVarP(&sep, "sep", "s", "/", "Delimiter")
	morseDecodeCmd.Flags().StringVarP(&sep, "sep", "s", "/", "Delimiter")
	rootCmd.AddCommand(caesarCmd)
	rootCmd.AddCommand(rot13Cmd)
	rootCmd.AddCommand(morseCmd)
	rootCmd.AddCommand(morseDecodeCmd)
}
