package basex

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/virzz/virzz/common"
)

// b16eCmd -
var b16eCmd = &cobra.Command{
	Use:   "b16e",
	Short: "Base16 Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := base16Encode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b16dCmd -
var b16dCmd = &cobra.Command{
	Use:   "b16d",
	Short: "Base16 Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := base16Decode(s)
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
		r, err := base32Encode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b32dCmd -
var b32dCmd = &cobra.Command{
	Use:   "b32d",
	Short: "Base32 Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := base32Decode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b36eCmd -
var b36eCmd = &cobra.Command{
	Use:   "b36e",
	Short: "Base36 Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := base36Encode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b36dCmd -
var b36dCmd = &cobra.Command{
	Use:   "b36d",
	Short: "Base36 Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := base36Decode(s)
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
		r, err := base58Encode(s, strings.ToLower(enc))
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
		r, err := base58Decode(s, strings.ToLower(enc))
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b62eCmd -
var b62eCmd = &cobra.Command{
	Use:   "b62e",
	Short: "Base62 Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := base62Encode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b62dCmd -
var b62dCmd = &cobra.Command{
	Use:   "b62d",
	Short: "Base62 Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := base62Decode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b64eCmd -
var b64eCmd = &cobra.Command{
	Use:   "b64e",
	Short: "Base64 Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := base64Encode(s, safe)
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
		r, err := base64Decode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b85eCmd -
var b85eCmd = &cobra.Command{
	Use:   "b85e",
	Short: "Base85 Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := base85Encode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b85dCmd -
var b85dCmd = &cobra.Command{
	Use:   "b85d",
	Short: "Base85 Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := base85Decode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b91eCmd -
var b91eCmd = &cobra.Command{
	Use:   "b91e",
	Short: "Base91 Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := base91Encode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b91dCmd -
var b91dCmd = &cobra.Command{
	Use:   "b91d",
	Short: "Base91 Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := base91Decode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b92eCmd -
var b92eCmd = &cobra.Command{
	Use:   "b92e",
	Short: "Base92 Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := base92Encode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b92dCmd -
var b92dCmd = &cobra.Command{
	Use:   "b92d",
	Short: "Base92 Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := base92Decode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b100eCmd -
var b100eCmd = &cobra.Command{
	Use:   "b100e",
	Short: "Base100 Encode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := base100Encode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// b100dCmd -
var b100dCmd = &cobra.Command{
	Use:   "b100d",
	Short: "Base100 Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := base100Decode(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var baseNCmd = &cobra.Command{
	Use:   "bnd",
	Short: "Auto Base-N Decode",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		var r string
		var res []string
		if r, err = base16Decode(s); err == nil {
			res = append(res, fmt.Sprintf("base16 : %s", r))
		}
		if r, err = base32Decode(s); err == nil {
			res = append(res, fmt.Sprintf("base32 : %s", r))
		}
		if r, err = base36Decode(s); err == nil {
			res = append(res, fmt.Sprintf("base36 : %s", r))
		}
		if r, err = base58Decode(s); err == nil {
			res = append(res, fmt.Sprintf("base58 : %s", r))
		}
		if r, err = base62Decode(s); err == nil {
			res = append(res, fmt.Sprintf("base62 : %s", r))
		}
		if r, err = base64Decode(s); err == nil {
			res = append(res, fmt.Sprintf("base64 : %s", r))
		}
		if r, err = base85Decode(s); err == nil {
			res = append(res, fmt.Sprintf("base85 : %s", r))
		}
		if r, err = base91Decode(s); err == nil {
			res = append(res, fmt.Sprintf("base91 : %s", r))
		}
		if r, err = base92Decode(s); err == nil {
			res = append(res, fmt.Sprintf("base92 : %s", r))
		}
		if r, err = base100Decode(s); err == nil {
			res = append(res, fmt.Sprintf("base100: %s", r))
		}
		return common.Output(strings.Join(res, "\n"))
	},
}

var (
	safe bool
	enc  string
)

var basexCmd = &cobra.Command{
	Use:   "basex",
	Short: "Base 16/32/58/62/64/85/91/92/100 Encode/Decode",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.ValidateRequiredFlags(); err != nil {
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	b64eCmd.Flags().BoolVarP(&safe, "safe", "s", false, "Use Safe Encode")
	b58eCmd.Flags().StringVarP(&enc, "enc", "e", "", "Use sepcial EncodeTable: [flickr|ripple]")

	basexCmd.AddCommand(
		b16eCmd, b16dCmd,
		b32eCmd, b32dCmd,
		b36eCmd, b36dCmd,
		b58eCmd, b58dCmd,
		b62eCmd, b62dCmd,
		b64eCmd, b64dCmd,
		b85eCmd, b85dCmd,
		b91eCmd, b91dCmd,
		b92eCmd, b92dCmd,
		b100eCmd, b100dCmd,
		baseNCmd,
	)
}

func ExportCommand() []*cobra.Command {
	return []*cobra.Command{basexCmd}
}
