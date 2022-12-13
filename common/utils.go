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

func GetFileBytes(path string) ([]byte, error) {
	f, err := os.Stat(path)
	if err == nil && !f.IsDir() {
		if f.Size() < 104857600 { // 100M
			data, err := os.ReadFile(path)
			if err != nil {
				return nil, err
			}
			return data, nil
		}
		return nil, fmt.Errorf("file is too bigger.(must <= 100M)")
	}
	return nil, fmt.Errorf("not found file")
}

func GetFileString(path string) (string, error) {
	data, err := GetFileBytes(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// GetFirstArg -
func GetFirstArg(args []string) (string, error) {
	// Priority: Args > Stdin > nil
	if len(args) > 0 {
		return args[0], nil
	}
	// Stdin
	fi, err := os.Stdin.Stat()
	if err != nil {
		return "", err
	}
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

func TableOutputV2(datas [][]*simpletable.Cell, header_footer ...[]*simpletable.Cell) string {
	table := simpletable.New()
	// Body
	table.Body.Cells = append(table.Body.Cells, datas...)
	if len(header_footer) > 0 {
		table.Header = &simpletable.Header{Cells: header_footer[0]}
	}
	if len(header_footer) > 1 {
		table.Footer = &simpletable.Footer{Cells: header_footer[1]}
	}
	return table.String()
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
		for _, footer := range header_footer[1] {
			cell = append(cell, &simpletable.Cell{
				Align: simpletable.AlignCenter,
				Text:  footer,
			})
		}
		table.Footer = &simpletable.Footer{Cells: cell}
	}
	return table.String()
}

func getEnvDefault(key, value string) string {
	v := os.Getenv(key)
	if v == "" {
		return value
	}
	return v
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

func sliceContains[T comparable](inputSlice []T, element T) bool {
	for _, inputValue := range inputSlice {
		if inputValue == element {
			return true
		}
	}
	return false
}

func getCommandAlias(prefix string, cmd *cobra.Command) []string {
	var res = make([]string, 0)
	for _, c := range cmd.Commands() {
		if c.HasSubCommands() && sliceContains(allowAlias, c.Name()) {
			res = append(res, getCommandAlias(fmt.Sprintf("%s %s", prefix, c.Name()), c)...)
		} else if sliceContains([]string{"help", "completion", "alias", "version"}, c.Name()) {
			continue
		} else {
			// alias cmd='prefix cmd'
			res = append(res, fmt.Sprintf("alias %s='%s %s'", c.Name(), prefix, c.Name()))
		}
	}
	return res
}

func AliasCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "alias",
		Short: "Print the version",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd = cmd.Root()
			res := getCommandAlias(cmd.CommandPath(), cmd)
			return Output(strings.Join(res, "\n"))
		},
	}
}
