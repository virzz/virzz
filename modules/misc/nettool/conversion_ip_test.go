package nettool

import (
	"fmt"
	"testing"

	"github.com/virzz/logger"
)

func init() {
	logger.SetDebug(true)
}

func TestIP2Dec(t *testing.T) {
	r, err := IP2Dec("192.168.1.1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "3232235777" {
		t.Fail()
	}
}
func TestIP2Oct(t *testing.T) {
	r, err := IP2Oct("192.168.1.1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "030052000401" {
		t.Fail()
	}
}

func TestIP2Hex(t *testing.T) {
	r, err := IP2Hex("192.168.1.1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "0xc0a80101" {
		t.Fail()
	}
}

func TestDec2IP(t *testing.T) {
	r, err := Dec2IP("3232235777")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "192.168.1.1" {
		t.Fail()
	}
}

func TestHex2IP(t *testing.T) {
	r, err := Hex2IP("0xc0a80101")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "192.168.1.1" {
		t.Fail()
	}
}

func TestIP2DotOct(t *testing.T) {
	r, err := IP2DotOct("127.0.0.1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "0177.0.0.01" {
		t.Fail()
	}
}
func TestDotOct2IP(t *testing.T) {
	r, err := DotOct2IP("0177.0.0.01")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "127.0.0.1" {
		t.Fail()
	}
}

func TestIP2DotHex(t *testing.T) {
	r, err := IP2DotHex("127.0.0.1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "0x7f.0.0.0x1" {
		t.Fail()
	}
}

func TestDotHex2IP(t *testing.T) {
	r, err := DotHex2IP("0x7f.0.0.0x1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "127.0.0.1" {
		t.Fail()
	}
}
