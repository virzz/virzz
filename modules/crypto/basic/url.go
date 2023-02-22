package basic

import (
	"net/url"
	"strings"
)

// URLEncode -
func URLEncode(s string, raw bool) (string, error) {
	s = url.QueryEscape(s)
	if raw {
		s = strings.ReplaceAll(s, "+", "%20")
	}
	return s, nil
}

// URLDecode -
func URLDecode(s string) (string, error) {
	return url.QueryUnescape(s)
}
