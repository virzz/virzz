package qrcode

import (
	"fmt"
	"os"
	"testing"
)

func TestZeroOneToQrcode(t *testing.T) {
	res, err := ZeroOneToQrcode("100110101001000101001100110101001000101001100110101001000101001100110101001000101001", false, "-")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestZeroOneToQrcodeFile(t *testing.T) {
	var fn = "../../../tests/qrcode/test_qrcode.png"
	res, err := ZeroOneToQrcode("100110101001000101001", false, fn)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
	err = os.Remove(fn)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseQrcode(t *testing.T) {
	res, err := ParseQrcode("../../../tests/qrcode/qrcode10.png", false)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestParseQrcodeTermial(t *testing.T) {
	res, err := ParseQrcode("../../../tests/qrcode/qrcode10.png", true)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestGenerateQrcode(t *testing.T) {
	res, err := GenerateQrcode("Mozhu233", "")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
