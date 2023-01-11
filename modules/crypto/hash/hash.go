package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/virzz/logger"

	"github.com/htruong/go-md2"
	//lint:ignore SA1019 Ignore deprecated md4 package
	"golang.org/x/crypto/md4"
	//lint:ignore SA1019 Ignore deprecated ripemd160 package
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

func ESha1Hash(s []byte) string {
	has := sha1.Sum(s)
	return hex.EncodeToString(has[:])
}

func sha1Hash(s string) (string, error) {
	return ESha1Hash([]byte(s)), nil
}

func sha3Hash(s string, size int) (string, error) {
	b := []byte(s)
	var res []byte
	switch size {
	case 224:
		dst := sha3.Sum224(b)
		res = dst[:]
	case 256:
		dst := sha3.Sum256(b)
		res = dst[:]
	case 384:
		dst := sha3.Sum384(b)
		res = dst[:]
	case 512:
		dst := sha3.Sum512(b)
		res = dst[:]
	default:
		return "", fmt.Errorf("not found size: %d", size)
	}
	return hex.EncodeToString(res), nil
}

func sha224Hash(s string) (string, error) {
	has := sha256.Sum224([]byte(s))
	return hex.EncodeToString(has[:]), nil
}

func sha256Hash(s string) (string, error) {
	has := sha256.Sum256([]byte(s))
	return hex.EncodeToString(has[:]), nil
}

func sha512Hash(s string, size int) (string, error) {
	logger.DebugF("size: %d", size)
	b := []byte(s)
	var res []byte
	switch size {
	case 224:
		dst := sha512.Sum512_224(b)
		res = dst[:]
	case 256:
		dst := sha512.Sum512_256(b)
		res = dst[:]
	case 384:
		dst := sha512.Sum384(b)
		res = dst[:]
	case 512:
		dst := sha512.Sum512(b)
		res = dst[:]
	default:
		return "", fmt.Errorf("not found size: %d", size)
	}
	return hex.EncodeToString(res), nil
}

func ripemd160Hash(s string) (string, error) {
	h := ripemd160.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil)), nil
}

func EMd5Hash(s []byte) string {
	has := md5.Sum(s)
	return hex.EncodeToString(has[:])
}

func md5Hash(s string) (string, error) {
	return EMd5Hash([]byte(s)), nil
}

func md4Hash(s string) (string, error) {
	h := md4.New()
	io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil)), nil
}

func md2Hash(s string) (string, error) {
	h := md2.New()
	io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil)), nil
}
