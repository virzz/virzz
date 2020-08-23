package common

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
)

// GetArgs -
func GetArgs(args []string) (string, error) {
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

// Output -
func Output(s string, color ...bool) error {
	if len(color) > 0 && color[0] {
		Logger.Success(s)
		return nil
	}
	outBuf := bufio.NewWriter(os.Stdout)
	outBuf.WriteString(s)
	outBuf.WriteString("\n")
	outBuf.Flush()
	return os.Stdout.Close()
}

// FIXME: I don't know why "virzz b64e README.md | virzz b64d" was faild

// CheckPort -
func CheckPort(port int) error {
	if port < 1 || port > 65535 {
		return fmt.Errorf("port should be a number and the range is [1,65536)")
	}
	return nil
}

// ParseAddr -
func ParseAddr(addr string) (string, int, error) {
	ipAndPort := strings.Split(addr, ":")
	if len(ipAndPort) != 2 {
		return "", 0, fmt.Errorf("address should be a string like [ip:port]")
	}
	ip := net.ParseIP(ipAndPort[0])
	if ip == nil {
		return "", 0, fmt.Errorf("parse ip faild")
	}
	port, err := strconv.ParseInt(ipAndPort[1], 10, 64)
	if err != nil {
		return "", 0, err
	}
	return ip.String(), int(port), nil
}
