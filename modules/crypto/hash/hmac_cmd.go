package hash

import (
	"github.com/spf13/cobra"
	"github.com/virzz/virzz/common"
)

var (
	macKey string
)

var hMd2Cmd = &cobra.Command{
	Use:   "md2",
	Short: "Hmac-MD2 hash algorithm",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := hmacMd2Hash(s, macKey)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}
var hMd4Cmd = &cobra.Command{
	Use:   "md4",
	Short: "Hmac-MD4 hash algorithm",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := hmacMd4Hash(s, macKey)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}
var hMd5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "Hmac-MD5 hash algorithm",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := hmacMd5Hash(s, macKey)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var hSha1Cmd = &cobra.Command{
	Use:   "sha1",
	Short: "Hmac-SHA1 hash algorithm",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := hmacSha1Hash(s, macKey)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var hSha224Cmd = &cobra.Command{
	Use:   "sha224",
	Short: "Hmac-SHA224 hash algorithm",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := hmacSha224Hash(s, macKey)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}
var hSha256Cmd = &cobra.Command{
	Use:   "sha256",
	Short: "Hmac-SHA256 hash algorithm",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := hmacSha256Hash(s, macKey)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}
var hSha384Cmd = &cobra.Command{
	Use:   "sha384",
	Short: "Hmac-SHA384 hash algorithm",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := hmacSha384Hash(s, macKey)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var hSha3Cmd = &cobra.Command{
	Use:   "sha3",
	Short: "Hmac-SHA3 hash algorithm",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := hmacSha3Hash(s, macKey, sha3Size)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}
var hSha512Cmd = &cobra.Command{
	Use:   "sha3",
	Short: "Hmac-SHA3 hash algorithm",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := hmacSha512Hash(s, macKey, sha512Size)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var hRipemd160Cmd = &cobra.Command{
	Use:   "ripemd160",
	Short: "Hmac-Ripemd160 hash algorithm",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := hmacRipemd160Hash(s, macKey)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var hSm3Cmd = &cobra.Command{
	Use:   "sm3",
	Short: "Hmac-SM3 hash algorithm",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := hmacSm3Hash(s, macKey)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var hmacCmd = &cobra.Command{
	Use:   "hmac",
	Short: "Some Hmac function",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	hmacCmd.PersistentFlags().StringVarP(&macKey, "key", "k", "", "MAC Key")
	hmacCmd.MarkFlagRequired("key")

	hSha3Cmd.Flags().IntVarP(&sha3Size, "size", "s", 256, "size: 224/256/384/512")
	hSha512Cmd.Flags().IntVarP(&sha512Size, "size", "s", 512, "size: 224/256/512")

	hmacCmd.AddCommand(
		// mdX
		hMd2Cmd,
		hMd4Cmd,
		hMd5Cmd,

		// shaX
		hSha1Cmd,
		hSha224Cmd,
		hSha384Cmd,
		hSha256Cmd,
		hSha3Cmd,
		hSha512Cmd,

		// ripemd160
		hRipemd160Cmd,

		// gmsm
		hSm3Cmd,
	)
}
