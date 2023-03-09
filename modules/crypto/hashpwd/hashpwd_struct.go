//go:generate stringer -type=EncMethed
//go:generate go run -tags hashpwd ./script/
//go:generate go fmt ./

package hashpwd

import (
	"encoding/hex"
	"fmt"

	"github.com/virzz/logger"
	"github.com/virzz/virzz/modules/crypto/hash"
)

type EncMethed int

const (
	MySQL EncMethed = iota
	NTLMv1
	MD5 // MD5MID // == MD5[8:-8]
	MD5x2
	MD5x3
	SHA1MD5
	SHA1
	SHA1x2 // MySQL5 // == SHA1x2
	MD5SHA1
	SHA256
	MD5SHA256
)

var MethedLen = map[EncMethed]int{
	MySQL:     17,
	NTLMv1:    32,
	MD5:       32,
	MD5x2:     32,
	MD5x3:     32,
	SHA1MD5:   32,
	SHA1:      40,
	SHA1x2:    40,
	MD5SHA1:   40,
	SHA256:    64,
	MD5SHA256: 64,
}

func GetPlaintext(line string) string {
	l := 0
	for _, s := range MethedLen {
		l += s
	}
	logger.Debug(line)
	return line[l:]
}

// FIXME: 多重加密使用原始数据加密，可能会有问题
// golang: md5(md5(xxx))
// == php: md5(md5(xxx,true))
// != php: md5(md5(xxx))
func Encrypt(passwd string) string {
	bytePwd := []byte(passwd)
	_hash := make(map[EncMethed]string, len(MethedLen))

	_hash[MySQL] = hash.MySQLHash(bytePwd)
	_hash[NTLMv1] = hash.NTLMv1Hash(bytePwd)
	_hash[SHA256], _ = hash.Sha2Hash(bytePwd, 256)

	tmp, _ := hash.MDHash(bytePwd, 5, true) // md5 x1
	_hash[MD5] = hex.EncodeToString([]byte(tmp))
	_hash[MD5SHA1], _ = hash.Sha1Hash([]byte(tmp))
	_hash[MD5SHA256], _ = hash.Sha2Hash([]byte(tmp), 256)

	tmp, _ = hash.MDHash([]byte(tmp), 5, true) // md5 x2
	_hash[MD5x2] = hex.EncodeToString([]byte(tmp))
	_hash[MD5x3], _ = hash.MDHash([]byte(tmp), 5) // md5 x3

	tmp, _ = hash.Sha1Hash(bytePwd, true) // sha1 x1
	_hash[SHA1] = hex.EncodeToString([]byte(tmp))
	_hash[SHA1MD5], _ = hash.MDHash([]byte(tmp), 5)
	_hash[SHA1x2], _ = hash.Sha1Hash([]byte(tmp)) // sha1 x2

	var t = 0
	var format = ``
	for j := 0; j < len(MethedLen); j++ {
		l := MethedLen[EncMethed(j)]
		format += fmt.Sprintf(fmt.Sprintf("%%%ds", l), _hash[EncMethed(j)])
		t += l
	}
	return fmt.Sprintf("%s%s", format, passwd)
}
