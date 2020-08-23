package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/siddontang/go/log"
	"github.com/spf13/cobra"
	"github.com/virink/virzz/common"
)

func getArgs(args []string) (string, error) {
	// Priority: Args > Stdin > nil

	// Args
	if len(args) > 0 {
		f, err := os.Stat(args[0])
		if err == nil && !f.IsDir() {
			if f.Size() < 104857600 { // 100M
				data, err := ioutil.ReadFile(args[0])
				if err != nil {
					return "", err
				}
				return strings.TrimSpace(string(data)), nil
			}
			return "", fmt.Errorf("file is too bigger.(must <= 100M)")
		}
		return args[0], nil
	}
	// Stdin
	if fi, err := os.Stdin.Stat(); err == nil &&
		(fi.Mode()&os.ModeNamedPipe) == os.ModeNamedPipe && fi.Size() > 0 {
		inBuf := bufio.NewReaderSize(os.Stdin, int(fi.Size()))
		data := make([]byte, fi.Size())
		_, err = inBuf.Read(data)
		if err != nil {
			return "", err
		}
		os.Stdin.Close()
		return string(data), nil
	}
	return "", fmt.Errorf("not found args")
}

func output(s string) error {
	outBuf := bufio.NewWriter(os.Stdout)
	outBuf.WriteString(s)
	outBuf.WriteString("\n")
	outBuf.Flush()
	return os.Stdout.Close()
}

var rootCmd = &cobra.Command{
	Use:   "jwttool",
	Short: "A jwt tool with Print/Crack/Modify",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		level := log.LevelError
		// Debug Mode
		if common.DebugMode {
			level = log.LevelDebug
		}
		// Logger
		common.InitLogger(level)
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.SuggestionsMinimumDistance = 1
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
