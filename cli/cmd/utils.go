package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	outBuf.Flush()
	return os.Stdout.Close()
}

// FIXME: I don't know why "virzz b64e README.md | virzz b64d" was faild
