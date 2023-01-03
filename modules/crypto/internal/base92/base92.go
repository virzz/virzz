// Package base92 fork from https://github.com/teal-finance/BaseXX
// with customizable encoding alphabet.
package base92

import (
	"fmt"
)

// Encoding alphabet is an optimized form of the encoding characters.
type Encoding struct {
	EncChars []byte
	DecMap   [128]int8
}

const (
	Radix       = 92 // approximation of ceil(log(256)/log(base)).
	numerator   = 16 // power of two -> speed up DecodeString()
	denominator = 13
)

const alphabet = " !" + // double-quote " removed
	"#$%&'()*+,-./0123456789:" + // semi-colon ; removed
	"<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[" + // back-slash \ removed
	"]^_`abcdefghijklmnopqrstuvwxyz{|}~"

// StdEncoding is the default encoding enc.
var StdEncoding = NewEncoding(alphabet)

// NewEncoding creates a new alphabet mapping.
//
// It panics if the passed string does not meet all requirements:
// its length (in bytes) must be the same as the base,
// all runes must be valid ASCII characters,
// and all characters must be different.
// Encoder string with non-printable characters are accepted.
func NewEncoding(encoder string) *Encoding {
	ret := new(Encoding)
	ret.EncChars = []byte(encoder)
	for i := range ret.DecMap {
		ret.DecMap[i] = -1
	}
	for i, b := range ret.EncChars {
		ret.DecMap[b] = int8(i)
	}
	return ret
}

// EncodeToString encodes binary bytes into Base92 bytes.
func (enc *Encoding) EncodeToString(bin []byte) string {
	return string(enc.Encode(bin))
}

// EncodeToString encodes binary bytes into a Base92 string.
func (enc *Encoding) Encode(bin []byte) []byte {
	size := len(bin)
	zcount := 0
	for zcount < size && bin[zcount] == 0 {
		zcount++
	}
	// It is crucial to make this as short as possible, especially for
	// the usual case of bitcoin addrs
	size = zcount +
		// This is an integer simplification of
		// ceil(log(256)/log(base))
		(size-zcount)*numerator/denominator + 1
	out := make([]byte, size)
	var i, high int
	var carry uint32
	high = size - 1
	for _, b := range bin {
		i = size - 1
		for carry = uint32(b); i > high || carry != 0; i-- {
			carry += 256 * uint32(out[i])
			out[i] = byte(carry % uint32(Radix))
			carry /= uint32(Radix)
		}
		high = i
	}
	// Determine the additional "zero-gap" in the buffer (aside from zcount)
	for i = zcount; i < size && out[i] == 0; i++ {
	}
	// Now encode the values with actual alphabet in-place
	val := out[i-zcount:]
	size = len(val)
	for i = 0; i < size; i++ {
		out[i] = enc.EncChars[val[i]]
	}
	return out[:size]
}

// DecodeString decodes a Base92 string into binary bytes.
func (enc *Encoding) DecodeString(str string) ([]byte, error) {
	if len(str) == 0 {
		return nil, nil
	}
	zero := enc.EncChars[0]
	strLen := len(str)
	var zcount int
	for i := 0; i < strLen && str[i] == zero; i++ {
		zcount++
	}
	var t, c uint64
	// the 32bit algo stretches the result up to 2 times
	binu := make([]byte, 2*((strLen*denominator/numerator)+1))
	outi := make([]uint32, (strLen+3)/4)
	for _, r := range str {
		if r > 127 {
			return nil, fmt.Errorf("base%d: high-bit set on invalid digit", Radix)
		}
		if enc.DecMap[r] == -1 {
			return nil, fmt.Errorf("base%d: invalid digit %q", Radix, r)
		}
		c = uint64(enc.DecMap[r])
		for j := len(outi) - 1; j >= 0; j-- {
			t = uint64(outi[j])*uint64(Radix) + c
			c = t >> 32
			outi[j] = uint32(t & 0xffffffff)
		}
	}
	// initial mask depends on b92sz, on further loops it always starts at 24 bits
	mask := (uint(strLen%4) * 8)
	if mask == 0 {
		mask = 32
	}
	mask -= 8
	outLen := 0
	for j := 0; j < len(outi); j++ {
		for mask < 32 { // loop relies on uint overflow
			binu[outLen] = byte(outi[j] >> mask)
			mask -= 8
			outLen++
		}
		mask = 24
	}
	// find the most significant byte post-decode, if any
	for msb := zcount; msb < len(binu); msb++ {
		if binu[msb] > 0 {
			return binu[msb-zcount : outLen], nil
		}
	}
	// it's all zeroes
	return binu[:outLen], nil
}
