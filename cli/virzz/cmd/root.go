package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/virink/virzz/common"
	log "github.com/virink/vlogger"
)

var (
	showVersion  bool = false
	setDebugMode bool
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("%s %s", common.AppName, common.Version))
	},
}

var rootCmd = &cobra.Command{
	Use:           common.BinName,
	Short:         "The Cyber Swiss Army Knife for terminal",
	SilenceErrors: true,
	SilenceUsage:  true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		level := log.LevelError
		if setDebugMode {
			common.DebugMode = true
			level = log.LevelDebug
		}
		common.InitLogger(level)
		// FIX ~
		if strings.HasPrefix(args[0], "~/") {
			if home := os.Getenv("HOME"); home != "" {
				args[0] = home
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
	rootCmd.PersistentFlags().BoolVarP(&setDebugMode, "debug", "D", false, "Set Debug Mode")
	rootCmd.SuggestionsMinimumDistance = 1
	rootCmd.AddCommand(versionCmd)
}

// Execute -
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Set(color.FgHiRed, color.Bold)
		defer color.Unset()
		// Error
		fmt.Fprintf(os.Stderr, "[-] %v\n", err)
		// UsageHelp
		os.Exit(1)
	}
}
