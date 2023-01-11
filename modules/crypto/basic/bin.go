package basic

import (
	"encoding/hex"
)

// BinToHex Bin -> Hex
func BinToHex(s []byte) (string, error) {
	return padHex(hex.EncodeToString(s)), nil
}

// HexToBin Hex -> Bin
func HexToBin(s string) ([]byte, error) {
	return hex.DecodeString(upPadHex(s))
}
