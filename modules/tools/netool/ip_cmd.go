package netool

import (
	"github.com/spf13/cobra"
	"github.com/virzz/virzz/common"
)

var ip2octCmd = &cobra.Command{
	Use:   "ip2oct",
	Short: "IPv4 -> Oct",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := ipToOct(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// ip2decCmd
var ip2decCmd = &cobra.Command{
	Use:   "ip2dec",
	Short: "IPv4 -> Dec",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := ipToDec(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// ip2hexCmd
var ip2hexCmd = &cobra.Command{
	Use:   "ip2hex",
	Short: "IPv4 -> Hex",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := ipToHex(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// ip2dotoctCmd
var ip2dotoctCmd = &cobra.Command{
	Use:   "ip2dotoct",
	Short: "IPv4 -> DotOct",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := ipToDotOct(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// ip2dothexCmd
var ip2dothexCmd = &cobra.Command{
	Use:   "ip2dothex",
	Short: "IPv4 -> DotHex",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := ipToDotHex(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// ip2allCmd
var ip2allCmd = &cobra.Command{
	Use:   "ip2all",
	Short: "IPv4 -> Oct,Dec,Hex",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := ipToAll(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// oct2ipCmd
var oct2ipCmd = &cobra.Command{
	Use:   "oct2ip",
	Short: "Oct -> IPv4",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := octToIP(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// dec2ipCmd
var dec2ipCmd = &cobra.Command{
	Use:   "dec2ip",
	Short: "Dec -> IPv4",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := decToIP(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// hex2ipCmd
var hex2ipCmd = &cobra.Command{
	Use:   "hex2ip",
	Short: "Hex -> IPv4",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := hexToIP(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// mac2decCmd
var mac2decCmd = &cobra.Command{
	Use:   "mac2dec",
	Short: "MAC -> Dec",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := macToDec(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

// dec2macCmd
var dec2macCmd = &cobra.Command{
	Use:   "dec2mac",
	Short: "Dec -> MAC",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := common.GetFirstArg(args)
		if err != nil {
			return err
		}
		r, err := decToMAC(s)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

var netoolCmd = &cobra.Command{
	Use:   "netool",
	Short: "Some net utils",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	netoolCmd.AddCommand(
		// IP ->
		ip2octCmd, ip2decCmd, ip2hexCmd,
		// IP -> Plus
		ip2dotoctCmd, ip2dothexCmd, ip2allCmd,
		// -> IP
		oct2ipCmd, dec2ipCmd, hex2ipCmd,
		// Mac
		mac2decCmd, dec2macCmd,
	)
}

func ExportCommand() []*cobra.Command {
	return []*cobra.Command{netoolCmd}
}
