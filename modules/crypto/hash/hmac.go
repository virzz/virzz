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

	"github.com/htruong/go-md2"

	//lint:ignore SA1019 Ignore deprecated md4 package
	"golang.org/x/crypto/md4"
	//lint:ignore SA1019 Ignore deprecated ripemd160 package
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

func HmacMDHash(s, key []byte, typ int) (string, error) {
	var res []byte
	switch typ {
	case 2:
		h := hmac.New(md2.New, key)
		h.Write(s)
		res = h.Sum(nil)[:]
	case 4:
		h := hmac.New(md4.New, key)
		h.Write(s)
		res = h.Sum(nil)[:]
	case 5:
		h := hmac.New(md5.New, key)
		h.Write(s)
		res = h.Sum(nil)[:]
	default:
		return "", fmt.Errorf("not found md%d", typ)
	}
	return hex.EncodeToString(res), nil
}

func HmacSha1Hash(s, key []byte) (string, error) {
	h := hmac.New(sha1.New, key)
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil)[:]), nil
}

func HmacSha2Hash(s, key []byte, size int) (string, error) {
	var f func() hash.Hash
	switch size {
	case 224:
		f = sha256.New224
	case 256:
		f = sha256.New
	case 384:
		f = sha512.New384
	case 512:
		f = sha512.New
	case 512224:
		f = sha512.New512_224
	case 512256:
		f = sha512.New512_256
	default:
		return "", fmt.Errorf("not method sha%d", size)
	}
	h := hmac.New(f, key)
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil)), nil
}

func HmacSha3Hash(s, key []byte, size int) (string, error) {
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
	h := hmac.New(f, key)
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil)), nil
}

func HmacRipemd160Hash(s, key []byte) (string, error) {
	h := hmac.New(ripemd160.New, key)
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil)[:]), nil
}
