package common

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

// GetFirstArg -
func GetFirstArg(args []string) (string, error) {
	// Priority: Args > Stdin > nil
	if len(args) > 0 {
		// File
		f, err := os.Stat(args[0])
		if err == nil && !f.IsDir() {
			if f.Size() < 104857600 { // 100M
				data, err := os.ReadFile(args[0])
				if err != nil {
					return "", err
				}
				return strings.TrimSpace(string(data)), nil
			}
			return "", fmt.Errorf("file is too bigger.(must <= 100M)")
		}
		// string
		return args[0], nil
	}
	// Stdin
	fi, err := os.Stdin.Stat()
	if err != nil {
		return "", err
	}
	// Wait os.Stdin flush
	if (fi.Mode() & os.ModeNamedPipe) == os.ModeNamedPipe {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}
	return "", fmt.Errorf("not found args")
}

// Output -
func Output(s string) error {
	outBuf := bufio.NewWriter(os.Stdout)
	outBuf.WriteString(s)
	outBuf.WriteString("\n")
	outBuf.Flush()
	return os.Stdout.Close()
}

// TableOutput func(data []map[int]string, header,footer []string)
func TableOutput(data []map[int]string, header_footer ...[]string) string {
	table := simpletable.New()
	// Body
	for _, row := range data {
		cell := []*simpletable.Cell{}
		for align, text := range row {
			cell = append(cell, &simpletable.Cell{Align: align, Text: text})
		}
		table.Body.Cells = append(table.Body.Cells, cell)
	}
	if len(header_footer) > 0 {
		// Header
		cell := []*simpletable.Cell{}
		for _, title := range header_footer[0] {
			cell = append(cell, &simpletable.Cell{
				Align: simpletable.AlignCenter,
				Text:  title,
			})
		}
		table.Header = &simpletable.Header{Cells: cell}
	}
	// Footer
	if len(header_footer) > 1 {
		cell := []*simpletable.Cell{}
		for _, title := range header_footer[1] {
			cell = append(cell, &simpletable.Cell{
				Align: simpletable.AlignCenter,
				Text:  title,
			})
		}
		table.Footer = &simpletable.Footer{Cells: cell}
	}
	return table.String()
}

func CompletionCommand() *cobra.Command {
	return &cobra.Command{
		Use:                   "completion [bash|zsh]",
		Short:                 "Generate completion script",
		DisableFlagsInUseLine: true,
		Hidden:                true,
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
}

func getEnvDefault(key, value string) string {
	v := os.Getenv(key)
	if v == "" {
		return value
	}
	return v
}
