package cmd

import (
	"github.com/spf13/cobra"
	"github.com/virink/virzz/misc/network"
)

// ip2decCmd
var ip2decCmd = &cobra.Command{
	Use:   "ip2dec",
	Short: "IPv4 -> Dec",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := network.IPToDec(s)
		if err != nil {
			return err
		}
		return output(r)
	},
}

// dec2ipCmd
var dec2ipCmd = &cobra.Command{
	Use:   "dec2ip",
	Short: "Dec -> IPv4",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := network.DecToIP(s)
		if err != nil {
			return err
		}
		return output(r)
	},
}

// ip2hexCmd
var ip2hexCmd = &cobra.Command{
	Use:   "ip2hex",
	Short: "IPv4 -> Hex",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := network.IPToHex(s)
		if err != nil {
			return err
		}
		return output(r)
	},
}

// hex2ipCmd
var hex2ipCmd = &cobra.Command{
	Use:   "hex2ip",
	Short: "Hex -> IPv4",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := network.HexToIP(s)
		if err != nil {
			return err
		}
		return output(r)
	},
}

// mac2decCmd
var mac2decCmd = &cobra.Command{
	Use:   "mac2dec",
	Short: "MAC -> Dec",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := network.MACToDec(s)
		if err != nil {
			return err
		}
		return output(r)
	},
}

// dec2macCmd
var dec2macCmd = &cobra.Command{
	Use:   "dec2mac",
	Short: "Dec -> MAC",
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getArgs(args)
		if err != nil {
			return err
		}
		r, err := network.DecToMAC(s)
		if err != nil {
			return err
		}
		return output(r)
	},
}

func init() {
	rootCmd.AddCommand(ip2decCmd)
	rootCmd.AddCommand(dec2ipCmd)
	rootCmd.AddCommand(ip2hexCmd)
	rootCmd.AddCommand(hex2ipCmd)
	rootCmd.AddCommand(mac2decCmd)
	rootCmd.AddCommand(dec2macCmd)
}