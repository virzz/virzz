package qrcode

import (
	"os"
	"testing"

	"github.com/virzz/virzz/common"
)

func TestZeroOneToQrcode(t *testing.T) {
	res, err := zeroOneToQrcode("100110101001000101001100110101001000101001100110101001000101001100110101001000101001", false, "-")
	if err != nil {
		t.Fatal(err)
	}
	err = common.Output(res)
	if err != nil {
		t.Fatal(err)
	}
}

func TestZeroOneToQrcodeFile(t *testing.T) {
	var fn = "../../../tmp/test_qrcode.png"
	res, err := zeroOneToQrcode("100110101001000101001", false, fn)
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
	res, err := parseQrcode("../../../tmp/qrcode10.png")
	if err != nil {
		t.Fatal(err)
	}
	err = common.Output(res)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseQrcodeTermial(t *testing.T) {
	res, err := parseQrcode("../../../tmp/qrcode10.png", true)
	if err != nil {
		t.Fatal(err)
	}
	err = common.Output(res)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerateQrcode(t *testing.T) {
	res, err := generateQrcode("Mozhu233", "-")
	if err != nil {
		t.Fatal(err)
	}
	err = common.Output(res)
	if err != nil {
		t.Fatal(err)
	}
}

func TestA(t *testing.T) {
	t.Log(len("\033[48;5;0m  \033[0m"))
}
