package base58

import (
	"bytes"
	"fmt"
	"math/big"
)

// https://github.com/itchyny/base58-go/blob/master/base58.go
// https://github.com/tv42/base58/blob/master/base58.go

// Encoding -
type Encoding struct {
	encode    [58]byte
	decodeMap [256]byte
}

const (
	encodeFlickr  = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
	encodeRipple  = "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"
	encodeBitcoin = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

func newEncoding(encoder string) *Encoding {
	e := new(Encoding)
	copy(e.encode[:], encoder)
	for i := 0; i < len(e.decodeMap); i++ {
		e.decodeMap[i] = 0xFF
	}
	for i, b := range encoder {
		e.decodeMap[b] = byte(i)
	}
	return e
}

var (
	// FlickrEncoding -
	FlickrEncoding = newEncoding(encodeFlickr)
	// RippleEncoding -
	RippleEncoding = newEncoding(encodeRipple)
	// BitcoinEncoding -
	BitcoinEncoding = newEncoding(encodeBitcoin)
)

// Encode -
func (enc *Encoding) Encode(dst, src []byte) {
	if len(src) == 0 {
		return
	}
	n := new(big.Int).SetBytes(src)
	radix := big.NewInt(58)
	zero := big.NewInt(0)
	for i := 0; n.Cmp(zero) > 0; i++ {
		mod := new(big.Int)
		n.DivMod(n, radix, mod)
		dst[i] = enc.encode[mod.Int64()]
	}
	// Reverse
	for i, j := 0, len(dst)-1; i < j; i, j = i+1, j-1 {
		dst[i], dst[j] = dst[j], dst[i]
	}
}

// EncodeToString returns the base58 encoding of src.
func (enc *Encoding) EncodeToString(src []byte) string {
	buf := make([]byte, (len(src)*138/100 + 1))
	enc.Encode(buf, src)
	return string(buf)
}

// Decode decodes src using the encoding enc.
func (enc *Encoding) Decode(dst, src []byte) (int, error) {
	if len(src) == 0 {
		return 0, nil
	}
	n := new(big.Int)
	radix := big.NewInt(58)
	for _, s := range src {
		b := enc.decodeMap[s]
		if b == 0xFF {
			return 0, fmt.Errorf("illegal base58 data at input byte 0xFF - %c", s)
		}
		n.Add(n.Mul(n, radix), big.NewInt(int64(b)))
	}
	r := bytes.Trim(n.Bytes(), "\x00")
	copy(dst, r)
	return len(r), nil
}

// DecodeString returns the bytes represented by the base58 string s.
func (enc *Encoding) DecodeString(s string) ([]byte, error) {
	dbuf := make([]byte, (len(s) * 733 / 1000))
	n, err := enc.Decode(dbuf, []byte(s))
	return dbuf[:n], err
}
