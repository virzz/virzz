package dsstore

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
	r, err := dsStore(filename, false)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(r)
}

func TestWebDSStore(t *testing.T) {
	r, err := dsStore("http://www.virzz.com/", false)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(r)
}
