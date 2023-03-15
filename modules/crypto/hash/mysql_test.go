package hash

import (
	"testing"

	"github.com/virzz/logger"
)

func TestMySQLHash(t *testing.T) {
	pwd := MySQLHash([]byte("test"))
	logger.Info(pwd)
}

func TestMySQL5Hash(t *testing.T) {
	pwd := MySQL5Hash([]byte("test"))
	logger.Info(pwd)
}
