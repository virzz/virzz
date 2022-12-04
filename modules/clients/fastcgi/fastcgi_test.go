package fastcgi

import (
	"fmt"
	"net/url"
	"testing"
)

func TestNewFastCGIRecord(t *testing.T) {
	env := map[string]string{
		"REMOTE_ADDR": "127.0.0.1",
		"PHP_VALUE":   "allow_url_include = On",
	}
	r := NewFastCGIRecord(env, []byte("orzzzzz"))
	rs := string(r)
	fmt.Println(rs)
	fmt.Println(url.QueryEscape(rs))
}
