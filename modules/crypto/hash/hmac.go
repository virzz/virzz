package hash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"

	"github.com/emmansun/gmsm/sm3"
	"github.com/htruong/go-md2"

	//lint:ignore SA1019 Ignore deprecated md4 package
	"golang.org/x/crypto/md4"
	//lint:ignore SA1019 Ignore deprecated ripemd160 package
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

func hmacSm3Hash(s, key string) (string, error) {
	h := hmac.New(sm3.New, []byte(key))
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil)[:]), nil
}

func hmacMd2Hash(s, key string) (string, error) {
	h := hmac.New(md2.New, []byte(key))
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil)[:]), nil
}

func hmacMd4Hash(s, key string) (string, error) {
	h := hmac.New(md4.New, []byte(key))
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil)[:]), nil
}

func hmacMd5Hash(s, key string) (string, error) {
	h := hmac.New(md5.New, []byte(key))
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil)[:]), nil
}

func hmacSha1Hash(s, key string) (string, error) {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil)[:]), nil
}

func hmacSha224Hash(s, key string) (string, error) {
	h := hmac.New(sha256.New224, []byte(key))
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil)[:]), nil
}
func hmacSha256Hash(s, key string) (string, error) {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil)[:]), nil
}
func hmacSha384Hash(s, key string) (string, error) {
	h := hmac.New(sha512.New384, []byte(key))
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil)[:]), nil
}

func hmacSha3Hash(s, key string, size int) (string, error) {
	var f func() hash.Hash
	switch size {
	case 224:
		f = sha3.New224
	case 256:
		f = sha3.New256
	case 384:
		f = sha3.New384
	case 512:
		f = sha3.New512
	default:
		return "", fmt.Errorf("not found size: %d", size)
	}
	h := hmac.New(f, []byte(key))
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil)), nil
}

func hmacSha512Hash(s, key string, size int) (string, error) {
	var f func() hash.Hash
	switch size {
	case 224:
		f = sha3.New224
	case 256:
		f = sha3.New256
	case 512:
		f = sha3.New512
	default:
		return "", fmt.Errorf("not found size: %d", size)
	}
	h := hmac.New(f, []byte(key))
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil)), nil
}

func hmacRipemd160Hash(s, key string) (string, error) {
	h := hmac.New(ripemd160.New, []byte(key))
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil)[:]), nil
}
