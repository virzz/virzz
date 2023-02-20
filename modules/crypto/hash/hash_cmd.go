package hash

import (
	"github.com/spf13/cobra"
	"github.com/virzz/virzz/utils"
)

var (
	sha3Size   int
	sha512Size int
	isFile     bool
)

var md2Cmd = &cobra.Command{
	Use:   "md2",
	Short: "MD2 hash algorithm",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := utils.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := md2Hash(s)
		if err != nil {
			return err
		}
		return utils.Output(r)
	},
}

var md4Cmd = &cobra.Command{
	Use:   "md4",
	Short: "MD4 hash algorithm",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := utils.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := md4Hash(s)
		if err != nil {
			return err
		}
		return utils.Output(r)
	},
}

var md5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "MD5 hash algorithm",
	RunE: func(cmd *cobra.Command, args []string) error {
		var r string
		if isFile {
			data, err := utils.GetFileBytes(args[0])
			if err != nil {
				return err
			}
			r = EMd5Hash(data)
		} else {
			s, err := utils.GetFirstArg(args)
			if err != nil {
				return err
			}
			r, _ = md5Hash(s)
		}
		return utils.Output(r)
	},
}

var sha1Cmd = &cobra.Command{
	Use:   "sha1",
	Short: "SHA1 hash algorithm",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := utils.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := sha1Hash(s)
		if err != nil {
			return err
		}
		return utils.Output(r)
	},
}

var sha3Cmd = &cobra.Command{
	Use:   "sha3",
	Short: "SHA3 hash algorithm",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := utils.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := sha3Hash(s, sha3Size)
		if err != nil {
			return err
		}
		return utils.Output(r)
	},
}

var sha224Cmd = &cobra.Command{
	Use:   "sha224",
	Short: "SHA224 hash algorithm",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := utils.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := sha224Hash(s)
		if err != nil {
			return err
		}
		return utils.Output(r)
	},
}

var sha256Cmd = &cobra.Command{
	Use:   "sha256",
	Short: "SHA256 hash algorithm",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := utils.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := sha256Hash(s)
		if err != nil {
			return err
		}
		return utils.Output(r)
	},
}

var sha384Cmd = &cobra.Command{
	Use:   "sha384",
	Short: "SHA384 hash algorithm",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := utils.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := sha512Hash(s, 384)
		if err != nil {
			return err
		}
		return utils.Output(r)
	},
}
var sha512Cmd = &cobra.Command{
	Use:   "sha512",
	Short: "SHA512 hash algorithm",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := utils.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := sha512Hash(s, sha512Size)
		if err != nil {
			return err
		}
		return utils.Output(r)
	},
}

var ripemd160Cmd = &cobra.Command{
	Use:   "ripemd160",
	Short: "RIPEMD160 hash algorithm",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := utils.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := ripemd160Hash(s)
		if err != nil {
			return err
		}
		return utils.Output(r)
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
	hashCmd.PersistentFlags().BoolVarP(&isFile, "file", "f", false, "read args from file")
	sha3Cmd.Flags().IntVarP(&sha3Size, "size", "s", 256, "size: 224/256/384/512")
	sha512Cmd.Flags().IntVarP(&sha512Size, "size", "s", 512, "size: 224/256/512")

	hashCmd.AddCommand(
		// mdX
		md2Cmd,
		md4Cmd,
		md5Cmd,

		// shaX
		sha1Cmd,
		sha224Cmd,
		sha384Cmd,
		sha256Cmd,
		sha3Cmd,
		sha512Cmd,

		// ripemd160
		ripemd160Cmd,

		// gmsm
		sm3Cmd,
	)
}
