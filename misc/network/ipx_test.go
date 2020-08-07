package network

import (
	"fmt"
	"testing"
)

func TestIPToDec(t *testing.T) {
	r, err := IPToDec("192.168.1.1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "3232235777" {
		t.Fail()
	}
}

func TestIPToHex(t *testing.T) {
	r, err := IPToHex("192.168.1.1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "0xc0a80101" {
		t.Fail()
	}
}

func TestDecToIP(t *testing.T) {
	r, err := DecToIP("3232235777")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "192.168.1.1" {
		t.Fail()
	}
}

func TestHexToIP(t *testing.T) {
	r, err := HexToIP("0xc0a80101")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "192.168.1.1" {
		t.Fail()
	}
}
