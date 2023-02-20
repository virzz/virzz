package base36

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var raw = []uint64{0, 50, 100, 999, 1000, 1111, 5959, 99999,
	123456789, 5481594952936519619, math.MaxInt64 / 2048, math.MaxInt64 / 512,
	math.MaxInt64, math.MaxUint64}

var encoded = []string{"0", "1E", "2S", "RR", "RS", "UV", "4LJ", "255R",
	"21I3V9", "15N9Z8L3AU4EB", "18CE53UN18F", "4XDKKFEK4XR",
	"1Y2P0IJ32E8E7", "3W5E11264SGSF"}

func TestEncode(t *testing.T) {
	for i, v := range raw {
		e := EncodeNumberToString(v)
		fmt.Printf("plain = %-15s, encode = %s\n", encoded[i], e)
		assert.Equal(t, encoded[i], e)
	}
}

func TestDecode(t *testing.T) {
	for i, v := range encoded {
		d1 := DecodeStringToMember(v)
		d2 := DecodeStringToMember(strings.ToLower(v))
		fmt.Printf("plain = %-20d, d1 = %-20d, d2 = %-20d\n", raw[i], d1, d2)
		assert.Equal(t, raw[i], d1)
		assert.Equal(t, raw[i], d2)
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncodeNumberToString(5481594952936519619)
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DecodeStringToMember("1Y2P0IJ32E8E7")
	}
}
