package basic

import (
	"fmt"
	"os"
	"testing"
)

func TestStringToASCII(t *testing.T) {
	r, err := StringToASCII("test_string_virzz")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
	if r != "116,101,115,116,95,115,116,114,105,110,103,95,118,105,114,122,122" {
		t.Fail()
	}
}

func TestASCIIToString(t *testing.T) {
	r, err := ASCIIToString("116,101,115,116,95,115,116,114,105,110,103,95,118,105,114,122,122")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
	if r != "test_string_virzz" {
		t.Fail()
	}
}

func TestHexToString(t *testing.T) {
	r, err := HexToString("0x746573745f737472696e675f7669727a7a")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
	if r != "test_string_virzz" {
		t.Fail()
	}
}

func TestStringToHex(t *testing.T) {
	r, err := StringToHex("test_string_virzz")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
	if r != "0x746573745f737472696e675f7669727a7a" {
		t.Fail()
	}
}

func TestDecToHex(t *testing.T) {
	r, err := DecToHex("1234567890987654321")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
	if r != "0x112210f4b16c1cb1" {
		t.Fail()
	}
}

func TestHexToDec(t *testing.T) {
	r, err := HexToDecStr("0x112210f4b16c1cb1")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
	if r != "1234567890987654321" {
		t.Fail()
	}
}

func TestHexToByteString(t *testing.T) {
	r, err := HexToByteString("0x746573745f11aa22bb33cc44dd55ee66ff7788995f737472696e67")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
	if r != `b'test_\x11\xaa"\xbb3\xccD\xddU\xeef\xffw\x88\x99_string'` {
		t.Fail()
	}
}

func TestByteStringToHex(t *testing.T) {
	r, err := ByteStringToHex(`b'test_\x11\xaa"\xbb3\xccD\xddU\xeef\xffw\x88\x99_string'`)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
	if r != "0x746573745f11aa22bb33cc44dd55ee66ff7788995f737472696e67" {
		t.Fail()
	}
}

func TestByteStringToString(t *testing.T) {
	r, err := ByteStringToString(`b'test_\x11\xaa"\xbb3\xccD\xddU\xeef\xffw\x88\x99_string'`)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
	if r != "test_\x11\xaa\"\xbb3\xccD\xddU\xeef\xffw\x88\x99_string" {
		t.Fail()
	}
}
