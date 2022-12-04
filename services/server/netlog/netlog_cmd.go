package netlog

import (
	"github.com/mozhu1024/virzz/common"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var netlogCmd = &cobra.Command{
	Use:   "netlog",
	Short: "Netlog Server",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
var (
	host      string
	port      int
	timeout   int
	localHost bool = false
)
var tcpLogCmd = &cobra.Command{
	Use:   "tcp",
	Short: "TCPLog Server",
	RunE: func(cmd *cobra.Command, args []string) error {
		if localHost {
			host = "127.0.0.1"
		}
		r, err := runTCPLogServer(host, port, timeout)
		if err != nil {
			return err
		}
		return common.Output(r)
	},
}

func init() {
	flagSet := &pflag.FlagSet{}
	flagSet.StringVarP(&host, "host", "H", "0.0.0.0", "Host(ip), default: 0.0.0.0")
	flagSet.IntVarP(&port, "port", "p", 6789, "Port, default: 6789")
	flagSet.IntVarP(&timeout, "timeout", "t", 5, "Timeout, default: 5s")
	flagSet.BoolVarP(&localHost, "localhost", "L", false, "Set host=127.0.0.1")
	// Common Persistent Flags
	// netlogCmd.PersistentFlags().AddFlagSet(flagSet)

	tcpLogCmd.Flags().AddFlagSet(flagSet)
	// tcpLogCmd.MarkFlagRequired("host")

	netlogCmd.AddCommand(tcpLogCmd)
}

func ExportCommand() []*cobra.Command {
	return []*cobra.Command{
		netlogCmd,
	}
}
