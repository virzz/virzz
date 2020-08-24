package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
		// Debug Mode
		if common.DebugMode || setDebugMode {
			setDebugMode = true
			common.DebugMode = true
			level = log.LevelDebug
		}
		// Logger
		common.InitLogger(level)
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

func init() {
	rootCmd.Flags().BoolVarP(&showVersion, "version", "V", false, "Show Version")
	rootCmd.PersistentFlags().BoolVarP(&setDebugMode, "debug", "D", false, "Set Debug Mode")
	rootCmd.SuggestionsMinimumDistance = 1
	rootCmd.AddCommand(versionCmd)
}

// Execute -
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		common.Logger.Error(err)
		os.Exit(1)
	}
}