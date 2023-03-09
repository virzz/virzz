package hash

import (
	"testing"

	"github.com/virzz/logger"
)

func TestNTLMv1Hash(t *testing.T) {
	r := NTLMv1Hash([]byte("testasdvafsd"))
	logger.InfoF("NTLM Hash: %s", r)
}
