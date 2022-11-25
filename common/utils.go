package common

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetArgs -
func GetArgs(args []string) (string, error) {
	// Priority: Args > Stdin > nil
	// Args
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

// Output -
func Output(s string) error {
	outBuf := bufio.NewWriter(os.Stdout)
	outBuf.WriteString(s)
	outBuf.WriteString("\n")
	outBuf.Flush()
	return os.Stdout.Close()
}

// FIXME: I don't know why "virzz b64e README.md | virzz b64d" was faild
