package qrcode

import (
	"os"
	"testing"

	"github.com/virzz/virzz/utils"
)

func TestZeroOneToQrcode(t *testing.T) {
	res, err := ZeroOneToQrcode("100110101001000101001100110101001000101001100110101001000101001100110101001000101001", false, "-")
	if err != nil {
		t.Fatal(err)
	}
	err = utils.Output(res)
	if err != nil {
		t.Fatal(err)
	}
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
	err = utils.Output(res)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseQrcodeTermial(t *testing.T) {
	res, err := ParseQrcode("../../../tests/qrcode/qrcode10.png", true)
	if err != nil {
		t.Fatal(err)
	}
	err = utils.Output(res)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerateQrcode(t *testing.T) {
	res, err := GenerateQrcode("Mozhu233", "")
	if err != nil {
		t.Fatal(err)
	}
	err = utils.Output(res)
	if err != nil {
		t.Fatal(err)
	}
}