package cmd

import (
	"github.com/spf13/cobra"
	cm "github.com/virink/virzz/common"
	"github.com/virink/virzz/misc/network"
)

func init() {
	// ip2octCmd
	var ip2octCmd = &cobra.Command{
		Use:   "ip2oct",
		Short: "IPv4 -> Oct",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			r, err := network.IPToOct(s)
			if err != nil {
				return err
			}
			return cm.Output(r)
		},
	}

	// ip2decCmd
	var ip2decCmd = &cobra.Command{
		Use:   "ip2dec",
		Short: "IPv4 -> Dec",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			r, err := network.IPToDec(s)
			if err != nil {
				return err
			}
			return cm.Output(r)
		},
	}

	// ip2hexCmd
	var ip2hexCmd = &cobra.Command{
		Use:   "ip2hex",
		Short: "IPv4 -> Hex",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			r, err := network.IPToHex(s)
			if err != nil {
				return err
			}
			return cm.Output(r)
		},
	}

	// ip2dotoctCmd
	var ip2dotoctCmd = &cobra.Command{
		Use:   "ip2dotoct",
		Short: "IPv4 -> DotOct",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			r, err := network.IPToDotOct(s)
			if err != nil {
				return err
			}
			return cm.Output(r)
		},
	}

	// ip2dothexCmd
	var ip2dothexCmd = &cobra.Command{
		Use:   "ip2dothex",
		Short: "IPv4 -> DotHex",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			r, err := network.IPToDotHex(s)
			if err != nil {
				return err
			}
			return cm.Output(r)
		},
	}

	// ip2allCmd
	var ip2allCmd = &cobra.Command{
		Use:   "ip2all",
		Short: "IPv4 -> Oct,Dec,Hex",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			r, err := network.IPToAll(s)
			if err != nil {
				return err
			}
			return cm.Output(r)
		},
	}

	// oct2ipCmd
	var oct2ipCmd = &cobra.Command{
		Use:   "oct2ip",
		Short: "Oct -> IPv4",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			r, err := network.OctToIP(s)
			if err != nil {
				return err
			}
			return cm.Output(r)
		},
	}

	// dec2ipCmd
	var dec2ipCmd = &cobra.Command{
		Use:   "dec2ip",
		Short: "Dec -> IPv4",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			r, err := network.DecToIP(s)
			if err != nil {
				return err
			}
			return cm.Output(r)
		},
	}

	// hex2ipCmd
	var hex2ipCmd = &cobra.Command{
		Use:   "hex2ip",
		Short: "Hex -> IPv4",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			r, err := network.HexToIP(s)
			if err != nil {
				return err
			}
			return cm.Output(r)
		},
	}

	// mac2decCmd
	var mac2decCmd = &cobra.Command{
		Use:   "mac2dec",
		Short: "MAC -> Dec",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			r, err := network.MACToDec(s)
			if err != nil {
				return err
			}
			return cm.Output(r)
		},
	}

	// dec2macCmd
	var dec2macCmd = &cobra.Command{
		Use:   "dec2mac",
		Short: "Dec -> MAC",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cm.GetArgs(args)
			if err != nil {
				return err
			}
			r, err := network.DecToMAC(s)
			if err != nil {
				return err
			}
			return cm.Output(r)
		},
	}

	// IP ->
	rootCmd.AddCommand(ip2octCmd, ip2decCmd, ip2hexCmd)
	// IP -> Plus
	rootCmd.AddCommand(ip2dotoctCmd, ip2dothexCmd, ip2allCmd)
	// -> IP
	rootCmd.AddCommand(oct2ipCmd, dec2ipCmd, hex2ipCmd)
	// Mac
	rootCmd.AddCommand(mac2decCmd, dec2macCmd)
}
