package basic

import (
	"fmt"
	"strings"

	"github.com/mozhu1024/virzz/common"
	"github.com/mozhu1024/virzz/logger"
	"github.com/spf13/cobra"
)

// b64eCmd -
var b64eCmd = &cobra.Command{
	Use:   "b64e",
	Short: "Base64 Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := Base64Encode(s, safe)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b64dCmd -
var b64dCmd = &cobra.Command{
	Use:   "b64d",
	Short: "Base64 Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := Base64Decode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b32eCmd -
var b32eCmd = &cobra.Command{
	Use:   "b32e",
	Short: "Base32 Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := Base32Encode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b32dCmd -
var b32dCmd = &cobra.Command{
	Use:   "b32d",
	Short: "Base64 Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := Base32Decode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b58eCmd -
var b58eCmd = &cobra.Command{
	Use:   "b58e",
	Short: "Base58 Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := Base58Encode(s, strings.ToLower(enc))
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b58dCmd -
var b58dCmd = &cobra.Command{
	Use:   "b58d",
	Short: "Base58 Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := Base58Decode(s, strings.ToLower(enc))
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var baseXCmd = &cobra.Command{
	Use:   "basex",
	Short: "Auto Base-X Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		var r string
		if r, err = Base64Decode(s); err != nil {
			if r, err = Base32Decode(s); err != nil {
				if r, err = Base58Decode(s); err != nil {
					return fmt.Errorf("auto decode base-x fail")
				}
				logger.Info("Base58")
			}
			logger.Info("Base32")
		}
		logger.Info("Base64")
		return common.Output(r)
	},
}
