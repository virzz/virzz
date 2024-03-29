// Package base36, form https://github.com/martinlindhe/base36
package base36

import (
	"math/big"
	"strings"
)

var (
	base36 = []byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
		'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
		'U', 'V', 'W', 'X', 'Y', 'Z'}

	uint8Index = []uint64{
		0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 1, 2,
		3, 4, 5, 6, 7, 8, 9, 0, 0, 0,
		0, 0, 0, 0, 10, 11, 12, 13, 14,
		15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
		25, 26, 27, 28, 29, 30, 31, 32, 33, 34,
		35, 0, 0, 0, 0, 0, 0, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 21, 22, 23,
		24, 25, 26, 27, 28, 29, 30, 31, 32, 33,
		34, 35, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, // 256
	}
	pow36Index = []uint64{
		1, 36, 1296, 46656, 1679616, 60466176,
		2176782336, 78364164096, 2821109907456,
		101559956668416, 3656158440062976,
		131621703842267136, 4738381338321616896,
		9223372036854775808,
	}
)

// Encode encodes a number to base36.
func EncodeNumberToString(value uint64) string {
	var res [16]byte
	var i int
	for i = len(res) - 1; ; i-- {
		res[i] = base36[value%36]
		value /= 36
		if value == 0 {
			break
		}
	}

	return string(res[i:])
}

// Decode decodes a base36-encoded string.
func DecodeStringToMember(s string) uint64 {
	if len(s) > 13 {
		s = s[:12]
	}
	res := uint64(0)
	l := len(s) - 1
	for idx := 0; idx < len(s); idx++ {
		c := s[l-idx]
		res += uint8Index[c] * pow36Index[idx]
	}
	return res
}

var bigRadix = big.NewInt(36)
var bigZero = big.NewInt(0)

// EncodeToBytes encodes to base36.
func EncodeToBytes(b string) []byte {
	x := new(big.Int).SetBytes([]byte(b))
	answer := make([]byte, 0, len(b)*136/100)
	for x.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, bigRadix, mod)
		answer = append(answer, base36[mod.Int64()])
	}
	// leading zero bytes
	for _, i := range b {
		if i != 0 {
			break
		}
		answer = append(answer, base36[0])
	}
	// reverse
	alen := len(answer)
	for i := 0; i < alen/2; i++ {
		answer[i], answer[alen-1-i] = answer[alen-1-i], answer[i]
	}
	return answer
}

// DecodeToBytes decodes a base36 string to a byte slice, using alphabet.
func DecodeToBytes(b string) []byte {
	alphabet := string(base36)
	answer := big.NewInt(0)
	j := big.NewInt(1)
	for i := len(b) - 1; i >= 0; i-- {
		tmp := strings.IndexAny(alphabet, string(b[i]))
		if tmp == -1 {
			return []byte("")
		}
		idx := big.NewInt(int64(tmp))
		tmp1 := big.NewInt(0)
		tmp1.Mul(j, idx)
		answer.Add(answer, tmp1)
		j.Mul(j, bigRadix)
	}
	tmpval := answer.Bytes()
	var numZeros int
	for numZeros = 0; numZeros < len(b); numZeros++ {
		if b[numZeros] != alphabet[0] {
			break
		}
	}
	flen := numZeros + len(tmpval)
	val := make([]byte, flen)
	copy(val[numZeros:], tmpval)
	return val
}

// Encode encodes to base36 string.
func Encode(b string) string {
	return string(EncodeToBytes(b))
}

// Decode decodes to string.
func Decode(b string) string {
	return string(DecodeToBytes(b))
}
