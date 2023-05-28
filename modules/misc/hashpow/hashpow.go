package hashpow

import (
	"bytes"

	"github.com/virzz/logger"
	"github.com/virzz/virzz/modules/crypto/hash"
	"github.com/virzz/virzz/utils"
	"github.com/virzz/virzz/utils/pool"
)

var (
	done   chan struct{}
	result chan string
)

type bruteArg struct {
	code, prefix, suffix, method string
	start, end                   int
	hash                         *func([]byte, ...bool) (string, error)
}

func hashBrute(arg bruteArg) bool {
	var buffer bytes.Buffer
	for {
		select {
		case <-done:
			return true
		default:
			buffer.Reset()
			tmp := utils.RandomBytesByLength(8)
			if len(arg.prefix) > 0 {
				buffer.WriteString(arg.prefix)
			}
			buffer.Write(tmp)
			if len(arg.suffix) > 0 {
				buffer.WriteString(arg.suffix)
			}
			// logger.Debug(arg.hash)
			// logger.Debug(*arg.hash)
			// logger.Debug(&arg.hash)
			if m, _ := (*arg.hash)(buffer.Bytes()); m[arg.start:arg.end] == arg.code {
				res := buffer.String()
				result <- res
				done <- struct{}{}
				logger.SuccessF("method: %s hash = %s result = %s", arg.method, m, res)
				close(result)
				close(done)
				return true
			}
		}
	}
}

// HashPoW Brute Hash Power of Work with md5/sha1
func HashPoW(code, prefix, suffix, method string, start int) string {
	done = make(chan struct{}, 1)
	result = make(chan string, 1)
	var _hash func([]byte, ...bool) (string, error)
	switch method {
	case "sha1":
		_hash = hash.Sha1Hash
	case "md5":
		_hash = func(s []byte, isRaw ...bool) (string, error) {
			return hash.MDHash(s, 5)
		}
	default:
		return "Error hash type!"
	}
	pool.Start(
		hashBrute,
		bruteArg{code, prefix, suffix, method, start, len(code) + start, &_hash},
	)
	return <-result
}
