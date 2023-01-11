package hash

import (
	"github.com/spf13/cobra"
	"github.com/virzz/virzz/common"
)

var md5Cmd = &cobra.Command{
	Use:   "md5 [string]",
	Short: "MD5 hash algorithm",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := md5Hash(args[0])
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var sha1Cmd = &cobra.Command{
	Use:   "sha1 [string]",
	Short: "SHA1 hash algorithm",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := sha1Hash(args[0])
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var sha224Cmd = &cobra.Command{
	Use:   "sha224 [string]",
	Short: "SHA224 hash algorithm",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := sha224Hash(args[0])
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var sha256Cmd = &cobra.Command{
	Use:   "sha256 [string]",
	Short: "SHA256 hash algorithm",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := sha256Hash(args[0])
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var sha384Cmd = &cobra.Command{
	Use:   "sha384 [string]",
	Short: "SHA384 hash algorithm",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := sha384Hash(args[0])
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}
var sha512Cmd = &cobra.Command{
	Use:   "sha512 [string]",
	Short: "SHA512 hash algorithm",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := sha512Hash(args[0])
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var sha512_224Cmd = &cobra.Command{
	Use:   "sha512_224 [string]",
	Short: "SHA512/224 hash algorithm",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := sha512_224Hash(args[0])
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var sha512_256Cmd = &cobra.Command{
	Use:   "sha512_256 [string]",
	Short: "SHA512/256 hash algorithm",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := sha512_256Hash(args[0])
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "Some hash function",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	hashCmd.AddCommand(
		md5Cmd,
		sha1Cmd, sha256Cmd, sha512Cmd,
		sha224Cmd, sha384Cmd,
		sha512_224Cmd, sha512_256Cmd,
	)
}

func ExportCommand() []*cobra.Command {
	return []*cobra.Command{
		hashCmd,
	}
}
