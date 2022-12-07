package hashpow

import (
	"bytes"
	"io"
	"math/rand"
	"time"

	"github.com/virzz/virzz/logger"
	"github.com/virzz/virzz/modules/crypto/hash"
)

const (
	letter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	done   = make(chan struct{})
	result = make(chan string, 1)
)

type randbo struct {
	rand.Source
}

func (r *randbo) Read(p []byte) (n int, err error) {
	todo := len(p)
	for {
		val := r.Int63()
		for todo > 0 {
			p[todo-1] = letter[int(val&(1<<6-1))%52]
			todo--
			if todo == 0 {
				return len(p), nil
			}
			val >>= 6
		}
	}
}

func newFrom(src rand.Source) io.Reader {
	return &randbo{src}
}

func newRandbo() io.Reader {
	return newFrom(rand.NewSource(time.Now().UnixNano()))
}

// brute -
// wg *sync.WaitGroup,
func brute(code, prefix, suffix, method string, pos, posend int) {
	// defer wg.Done()
	var _hash func(str []byte) string
	if method == "sha1" {
		_hash = hash.ESha1Hash
	} else if method == "md5" {
		_hash = hash.EMd5Hash
	} else {
		result <- "Error hash type!"
		return
	}
	var buffer bytes.Buffer
	r := newRandbo()
	for {
		select {
		case <-done:
			return
		default:
			buffer.Reset()
			tmp := make([]byte, 8)
			if _, err := r.Read(tmp); err != nil {
				close(done)
				return
			}
			if len(prefix) > 0 {
				buffer.WriteString(prefix)
			}
			buffer.Write(tmp)
			if len(suffix) > 0 {
				buffer.WriteString(suffix)
			}
			if m := _hash(buffer.Bytes()); m[pos:posend] == code {
				logger.Debug(string(tmp))
				result <- string(tmp)
				logger.InfoF("hash = %s result = %s", m, string(tmp))
				close(result)
				close(done)
				return
			}
		}
	}
}

func doBrute(code, prefix, suffix, method string, pos int) string {
	posend := len(code) + pos
	for i := 0; i < 16; i++ {
		go brute(code, prefix, suffix, method, pos, posend)
	}
	return <-result
}
