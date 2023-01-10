package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/virzz/logger"
	"github.com/virzz/virzz/common"
	"github.com/virzz/virzz/modules/crypto/basex"
	"github.com/virzz/virzz/modules/crypto/basic"
	"github.com/virzz/virzz/modules/crypto/classical"
	"github.com/virzz/virzz/modules/crypto/hash"
	"github.com/virzz/virzz/modules/crypto/hashpow"
	"github.com/virzz/virzz/modules/parser"
	"github.com/virzz/virzz/modules/tools/domain"
	"github.com/virzz/virzz/modules/tools/dsstore"
	"github.com/virzz/virzz/modules/tools/netool"
	"github.com/virzz/virzz/modules/tools/qrcode"
	"github.com/virzz/virzz/modules/web/gopher"
	"github.com/virzz/virzz/modules/web/jwttool"
	"github.com/virzz/virzz/modules/web/leakcode/githack"
	"github.com/virzz/virzz/services/server/netlog"
)

var (
	showVersion bool = false

	AppName        = "Virzz"
	BinName        = "virzz"
	Version string = "dev" // git tag | tail -1
	BuildID string = "0"   // head .buildid
)

var versionCmd = common.VersionCommand(AppName, Version, BuildID)

var rootCmd = &cobra.Command{
	Use:           BinName,
	Short:         "The Cyber Swiss Army Knife for terminal",
	SilenceErrors: true,
	SilenceUsage:  true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if strings.HasPrefix(args[0], "~/") {
				if home := os.Getenv("HOME"); home != "" {
					args[0] = filepath.Join(home, args[0][2:len(args[0])])
				}
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			versionCmd.Run(cmd, args)
			return
		}
		cmd.Help()
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&showVersion, "version", "V", false, "Show Version")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(common.CompletionCommand())
	rootCmd.AddCommand(common.AliasCommand())

	// CMD
	// Crypto
	rootCmd.AddCommand(basex.ExportCommand()...)
	rootCmd.AddCommand(basic.ExportCommand()...)
	rootCmd.AddCommand(classical.ExportCommand()...)
	rootCmd.AddCommand(hash.ExportCommand()...)
	rootCmd.AddCommand(hashpow.ExportCommand()...)
	// Web
	rootCmd.AddCommand(githack.ExportCommand()...)
	rootCmd.AddCommand(gopher.ExportCommand()...)
	rootCmd.AddCommand(jwttool.ExportCommand()...)
	// Tools
	rootCmd.AddCommand(qrcode.ExportCommand()...)
	rootCmd.AddCommand(netool.ExportCommand()...)
	rootCmd.AddCommand(dsstore.ExportCommand()...)
	rootCmd.AddCommand(domain.ExportCommand()...)
	// Services
	// -> server
	rootCmd.AddCommand(netlog.ExportCommand()...)
	// Parser
	rootCmd.AddCommand(parser.ExportCommand()...)

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
