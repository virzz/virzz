package netool

import (
	"fmt"
	"testing"
)

func TestIPToDec(t *testing.T) {
	r, err := ipToDec("192.168.1.1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "3232235777" {
		t.Fail()
	}
}
func TestIPToOct(t *testing.T) {
	r, err := ipToOct("192.168.1.1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "030052000401" {
		t.Fail()
	}
}

func TestIPToHex(t *testing.T) {
	r, err := ipToHex("192.168.1.1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "0xc0a80101" {
		t.Fail()
	}
}

func TestDecToIP(t *testing.T) {
	r, err := decToIP("3232235777")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "192.168.1.1" {
		t.Fail()
	}
}

func TestHexToIP(t *testing.T) {
	r, err := hexToIP("0xc0a80101")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "192.168.1.1" {
		t.Fail()
	}
}

func TestMACToDec(t *testing.T) {
	r, err := macToDec("00:00:5e:00:53:01")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "1577079553" {
		t.Fail()
	}
}

func TestDecToMAC(t *testing.T) {
	r, err := decToMAC("1577079553")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
	if r != "00:00:5e:00:53:01" {
		t.Fail()
	}
}
