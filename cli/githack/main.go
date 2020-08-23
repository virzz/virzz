package main

import (
	"errors"
	"os"
	"strings"

	"github.com/siddontang/go/log"
	"github.com/spf13/cobra"
	"github.com/virink/virzz/common"
	"github.com/virink/virzz/web/leakcode/githack"
	"gopkg.in/go-playground/validator.v9"
)

var rootCmd = &cobra.Command{
	Use:   "githack",
	Short: "A `.git` folder disclosure exploit",
	Long:  `A Git source leak exploit tool that restores the entire Git repository, including data from stash, for white-box auditing and analysis of developers' mind`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		debugEnv := os.Getenv("DEBUG")
		level := log.LevelError
		if debugEnv != "" && debugEnv != "0" && debugEnv != "false" {
			common.DebugMode = true
			level = log.LevelDebug
		}
		common.InitLogger(level)
	},
	Args: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) < 1 {
			return errors.New("Requires a url argument")
		}
		if err = validator.New().Var(args[0], "url"); err != nil {
			return
		}
		if !strings.HasPrefix(args[0], "http") {
			return errors.New("must be http(s)")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if delay > 0 {
			limit = 1
		}
		return githack.DoAction(args[0], limit, delay)
	},
}

var (
	limit int64
	delay int64
)

func init() {
	rootCmd.Flags().Int64VarP(&limit, "limit", "l", 10, "Request limit (N times one second)")
	rootCmd.Flags().Int64VarP(&delay, "delay", "d", 0, "Request delay (N times one second)")
	rootCmd.SuggestionsMinimumDistance = 1
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
