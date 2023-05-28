package hash

import (
	"encoding/hex"

	//lint:ignore SA1019 Ignore deprecated md4 package
	"golang.org/x/crypto/md4"
)

func utf16le(s []byte) []byte {
	u16 := make([]byte, len(s)*2)
	for i, b := range s {
		u16[i*2] = b
		u16[i*2+1] = 0
	}
	return u16
}
func NTLMv1Hash(s []byte) string {
	hash := md4.New()
	hash.Write(utf16le(s))
	ntlmHash := hash.Sum(nil)
	return hex.EncodeToString(ntlmHash[:])
}
