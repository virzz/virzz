package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/alexeyco/simpletable"
	"github.com/urfave/cli/v3"
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

// GetArgFilePipe Get arg from c.Arg > File > Pipe
func GetArgFilePipe(arg string) ([]byte, error) {
	if arg == "" {
		return GetFromPipe()
	}
	data, err := GetFileBytes(arg)
	if err != nil {
		return []byte(arg), nil
	}
	return data, nil
}

func GetFileString(path string) (string, error) {
	data, err := GetFileBytes(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func GetFromPipe() ([]byte, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}
	if (fi.Mode() & os.ModeNamedPipe) == os.ModeNamedPipe {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	return nil, fmt.Errorf("not found from pipe")
}

func GetFileOrPipe(args []string) ([]byte, error) {
	if len(args) > 0 {
		return GetFileBytes(args[0])
	}
	return GetFromPipe()
}

// ArgOrPipe -
func ArgOrPipe(args string) (string, error) {
	// Priority: Args > Stdin > nil
	if len(args) > 0 {
		return args, nil
	}
	// Stdin
	if data, err := GetFromPipe(); err == nil {
		return string(data), err
	}
	return "", fmt.Errorf("not found arg")
}

// GetFirstArg -
func GetFirstArg(args []string) (string, error) {
	// Priority: Args > Stdin > nil
	if len(args) > 0 {
		return args[0], nil
	}
	// Stdin
	if data, err := GetFromPipe(); err == nil {
		return string(data), err
	}
	return "", fmt.Errorf("not found args")
}

// OutputBytes -
func OutputBytes(s []byte) error {
	outBuf := bufio.NewWriter(os.Stdout)
	outBuf.Write(s)
	outBuf.Flush()
	return os.Stdout.Close()
}

func TableOutput(datas [][]*simpletable.Cell, header, footer []*simpletable.Cell) string {
	table := simpletable.New()
	// Body
	table.Body.Cells = append(table.Body.Cells, datas...)
	if header != nil {
		table.Header = &simpletable.Header{Cells: header}
	}
	if footer != nil {
		table.Footer = &simpletable.Footer{Cells: footer}
	}
	return table.String()
}

var HideHelpCommand func(c []*cli.Command)

func init() {
	HideHelpCommand = func(c []*cli.Command) {
		for _, cmd := range c {
			cmd.HideHelpCommand = true
			HideHelpCommand(cmd.Commands)
		}
	}

}
