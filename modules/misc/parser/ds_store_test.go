package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestDSStore(t *testing.T) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		t.Error(err)
		return
	}
	filename := filepath.Join(homePath, ".DS_Store")
	r, err := DSStore(filename)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(r)
}

func TestWebDSStore(t *testing.T) {
	r, err := DSStore("http://www.virzz.com/")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(r)
}
