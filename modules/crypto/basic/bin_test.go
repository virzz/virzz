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

func TestBinStrToHex(t *testing.T) {
	r, err := BinStrToHex("100101000101011101")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(r)
}
