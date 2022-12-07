package basic

import (
	"encoding/hex"
)

// BinToHex Bin -> Hex
func BinToHex(s []byte) (string, error) {
	return padHex(hex.EncodeToString(s)), nil
}

// BinStrToHex BinString -> Hex
func BinStrToHex(s string) (string, error) {
	return "TODO", nil
	// return PadHex(hex.EncodeToString(s)), nil
}
