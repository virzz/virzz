package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/virink/virzz/common"
	"github.com/virink/virzz/pentest/proxy"
)

var (
	timeout    int = 5
	localPort  int
	remotePort int
	localHost  string
	remoteHost string
	localAddr  string
	remoteAddr string
)

// lcxCmd
var lcxCmd = &cobra.Command{
	Use:   "lcx",
	Short: "Lcx",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

// listenCmd
var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Listen [local port] [remote port]",
	Args: func(cmd *cobra.Command, args []string) (err error) {
		for _, arg := range args {
			_, err := strconv.ParseInt(arg, 10, 64)
			if err != nil {
				return err
			}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// lcx listen [local port] [remote port]
		var lport, rport int
		if len(args) > 0 {
			_port, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}
			lport = int(_port)
		}
		if len(args) > 1 {
			_port, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}
			rport = int(_port)
		}
		if localPort > 0 {
			lport = localPort
		}
		if remotePort > 0 {
			rport = remotePort
		}
		if err := common.CheckPort(lport); err != nil {
			return err
		}
		if err := common.CheckPort(lport); err != nil {
			return err
		}
		common.Logger.Debug(lport, rport, timeout)
		return proxy.LcxListen(lport, rport, timeout)
	},
}

// tranCmd
var tranCmd = &cobra.Command{
	Use:   "tran",
	Short: "tran [local port] [remote addr]",
	RunE: func(cmd *cobra.Command, args []string) error {
		// lcx tran [local port] [remote addr]
		var (
			lport, rport int
			rhost        string
		)
		if len(args) > 0 {
			_port, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return err
			}
			lport = int(_port)
		}
		if len(args) > 1 {
			_ip, _port, err := common.ParseAddr(args[1])
			if err != nil {
				return err
			}
			rport = int(_port)
			rhost = _ip
		}
		if localPort > 0 {
			lport = localPort
		}
		if remotePort > 0 {
			rport = remotePort
		}
		if len(remoteHost) > 0 {
			rhost = remoteHost
		}
		if err := common.CheckPort(lport); err != nil {
			return err
		}
		if err := common.CheckPort(rport); err != nil {
			return err
		}
		return proxy.LcxTran(lport, rhost, rport, timeout)
	},
}

func init() {
	lcxCmd.PersistentFlags().IntVarP(&timeout, "timeout", "t", 5, "Connect delay")
	lcxCmd.PersistentFlags().IntVarP(&localPort, "lport", "l", 0, "Local Port")
	lcxCmd.PersistentFlags().IntVarP(&remotePort, "rport", "r", 0, "Remote Port")
	tranCmd.Flags().StringVarP(&remoteHost, "rhost", "H", "", "Remote Host")
	// tranCmd.Flags().StringVarP(&localHost, "lhost", "h", "", "Local Host")
	lcxCmd.AddCommand(listenCmd)
	lcxCmd.AddCommand(tranCmd)
	rootCmd.AddCommand(lcxCmd)
}
