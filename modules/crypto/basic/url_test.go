package basic

import (
	"testing"
)

func TestURLEncode(t *testing.T) {
	r, err := URLEncode("argver\x09\t\\=!@#$%^&*")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestURLDecode(t *testing.T) {
	r, err := URLDecode("argver%09%09%5C%3D%21%40%23%24%25%5E%26%2A")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
