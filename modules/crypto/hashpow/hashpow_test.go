package hashpow

import (
	"fmt"
	"testing"

	"github.com/virzz/logger"
)

var prefix = "orzz"

func init() {
	logger.SetDebug(true)
}
func TestHashPoWMore(t *testing.T) {
	for i := 0; i < 10; i++ {
		HashPoW("aaaaa", fmt.Sprintf("%s%d", prefix, i), "", "md5", 0)
	}
}

func TestHashPoW(t *testing.T) {
	HashPoW("aaaaa", prefix, "", "md5", 0)
}
