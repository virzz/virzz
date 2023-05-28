package hashpow

import (
	"fmt"
	"testing"
)

var prefix = "orzz"

func TestHashPoWMore(t *testing.T) {
	for i := 0; i < 2; i++ {
		HashPoW("aaaaa", fmt.Sprintf("%s%d", prefix, i), "", "md5", 0)
	}
}

func TestHashPoW(t *testing.T) {
	HashPoW("aaaaa", prefix, "", "md5", 0)
}
