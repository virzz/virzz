package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/virzz/logger"
	"github.com/virzz/virzz/modules/crypto/basex"
	"github.com/virzz/virzz/modules/crypto/basic"
	"github.com/virzz/virzz/modules/crypto/classical"
	"github.com/virzz/virzz/modules/crypto/hash"
	"github.com/virzz/virzz/modules/tools/domain"
	"github.com/virzz/virzz/modules/tools/dsstore"
	"github.com/virzz/virzz/modules/tools/netool"
	"github.com/virzz/virzz/modules/tools/qrcode"
)

var (
	AppName        = "God"
	BinName        = "god"
	Version string = "latest"
	BuildID string = "0"
)

var rootCmd = &cobra.Command{
	Use:           BinName,
	Short:         "The Cyber Swiss Army Knife for terminal",
	SilenceErrors: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// CMD
	// Crypto
	rootCmd.AddCommand(basex.ExportCommand()...)
	rootCmd.AddCommand(basic.ExportCommand()...)
	rootCmd.AddCommand(classical.ExportCommand()...)
	rootCmd.AddCommand(hash.ExportCommand()...)
	// rootCmd.AddCommand(hashpow.ExportCommand()...)
	// Web
	// rootCmd.AddCommand(githack.ExportCommand()...)
	// rootCmd.AddCommand(gopher.ExportCommand()...)
	// rootCmd.AddCommand(jwttool.ExportCommand()...)
	// Tools
	rootCmd.AddCommand(qrcode.ExportCommand()...)
	rootCmd.AddCommand(netool.ExportCommand()...)
	rootCmd.AddCommand(dsstore.ExportCommand()...)
	rootCmd.AddCommand(domain.ExportCommand()...)
	// Parser
	// rootCmd.AddCommand(parser.ExportCommand()...)

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
