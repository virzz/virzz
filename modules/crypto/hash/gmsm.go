package hash

import (
	"encoding/hex"

	"github.com/emmansun/gmsm/sm3"
)

func sm3Hash(s string) (string, error) {
	h := sm3.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil)), nil
}
