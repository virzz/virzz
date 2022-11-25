package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/mozhu1024/virzz/crypto/hash"
	"github.com/mozhu1024/virzz/logger"
	"github.com/mozhu1024/virzz/misc/basic"
	"github.com/mozhu1024/virzz/web/gopher"
	"github.com/mozhu1024/virzz/web/jwttool"
	"github.com/mozhu1024/virzz/web/leakcode/githack"

	"github.com/mozhu1024/virzz/common"
	"github.com/spf13/cobra"
)

var (
	showVersion  bool = false
	setDebugMode bool

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
		// Args
		if len(args) > 0 {
			// FIX ~ as HOME
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

var completionCmd = &cobra.Command{
	Use:                   "completion [bash|zsh]",
	Short:                 "Generate completion script",
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		}
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&showVersion, "version", "V", false, "Show Version")
	rootCmd.PersistentFlags().BoolVarP(&setDebugMode, "debug", "D", false, "Set Debug Mode")
	rootCmd.SuggestionsMinimumDistance = 1
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(completionCmd)

	// CMD
	// Crypto
	rootCmd.AddCommand(basic.BasicCmd()...)
	// Web
	rootCmd.AddCommand(githack.GithackCmd)
	rootCmd.AddCommand(gopher.GopherCmd)
	rootCmd.AddCommand(jwttool.JWTToolCmd)
	rootCmd.AddCommand(hash.HashCmd...)
	rootCmd.AddCommand(hash.HashPowCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
