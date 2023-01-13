package common

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/alexeyco/simpletable"
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

func GetFileOrPipe(args []string) ([]byte, error) {
	if len(args) > 0 {
		return GetFileBytes(args[0])
	}
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
	return nil, fmt.Errorf("not found pipe")
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

// OutputBytes -
func OutputBytes(s []byte) error {
	outBuf := bufio.NewWriter(os.Stdout)
	outBuf.Write(s)
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
