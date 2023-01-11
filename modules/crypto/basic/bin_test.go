package basic

import (
	"fmt"
	"os"
	"testing"
)

func TestBintoHex(t *testing.T) {
	r, err := BinToHex([]byte("100101000101011101"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
}

func TestHexToBin(t *testing.T) {
	r, err := HexToBin("0x0cc175b9c0f1b6a831c399e269772661")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	os.Stdout.Write(r)
	os.Stdout.WriteString("\n")
}
