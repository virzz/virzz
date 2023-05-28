package base58

import (
	"fmt"
	"os"
	"testing"
)

var testStr = []byte("test_base58_string")
var encStr = []byte("5q1dAkvfMPRxpkkujHtkssust")

func TestEncode(t *testing.T) {
	l := len(encStr)*138/100 + 1
	res := make([]byte, l)
	// FlickrEncoding
	FlickrEncoding.Encode(res, testStr)
	fmt.Println("Flickr", string(res))
	// RippleEncoding
	RippleEncoding.Encode(res, testStr)
	fmt.Println("Ripple", string(res))
	// BitcoinEncoding
	BitcoinEncoding.Encode(res, testStr)
	fmt.Println("Bitcoin", string(res))
}
func TestEncodeToString(t *testing.T) {
	// FlickrEncoding
	res := FlickrEncoding.EncodeToString(testStr)
	fmt.Println("Flickr", res)
	// RippleEncoding
	res = RippleEncoding.EncodeToString(testStr)
	fmt.Println("Ripple", res)
	// BitcoinEncoding
	res = BitcoinEncoding.EncodeToString(testStr)
	fmt.Println("Bitcoin", res)
}

func TestDecode(t *testing.T) {
	// var res []byte
	l := len(encStr) * 733 / 1000
	res := make([]byte, l)
	// FlickrEncoding
	n, err := FlickrEncoding.Decode(res, encStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(n, string(res))
	// RippleEncoding
	n, err = RippleEncoding.Decode(res, encStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(n, string(res))
	// BitcoinEncoding
	n, err = BitcoinEncoding.Decode(res, encStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(n, string(res))
}

func TestDecodeString(t *testing.T) {
	res, err := FlickrEncoding.DecodeString(string(encStr))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(string(res))
}
