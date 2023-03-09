package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/htruong/go-md2"

	//lint:ignore SA1019 Ignore deprecated md4 package
	"golang.org/x/crypto/md4"
	//lint:ignore SA1019 Ignore deprecated ripemd160 package
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

func Sha1Hash(s []byte, isRaw ...bool) (string, error) {
	has := sha1.Sum(s)
	if len(isRaw) > 0 && isRaw[0] {
		return string(has[:]), nil
	}
	return hex.EncodeToString(has[:]), nil
}

func Sha2Hash(b []byte, size int) (string, error) {
	var res []byte
	switch size {
	case 224:
		dst := sha256.Sum224(b)
		res = dst[:]
	case 256:
		dst := sha256.Sum256(b)
		res = dst[:]
	case 384:
		dst := sha512.Sum384(b)
		res = dst[:]
	case 512:
		dst := sha512.Sum512(b)
		res = dst[:]
	case 512224:
		dst := sha512.Sum512_224(b)
		res = dst[:]
	case 512256:
		dst := sha512.Sum512_256(b)
		res = dst[:]
	default:
		return "", fmt.Errorf("not found size: %d", size)
	}
	return hex.EncodeToString(res), nil
}

func Sha3Hash(b []byte, size int) (string, error) {
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

func Ripemd160Hash(s []byte) (string, error) {
	h := ripemd160.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil)), nil
}

func MDHash(s []byte, typ int, isRaw ...bool) (string, error) {
	var res []byte
	switch typ {
	case 2:
		h := md2.New()
		h.Write(s)
		res = h.Sum(nil)[:]
	case 4:
		h := md4.New()
		h.Write(s)
		res = h.Sum(nil)[:]
	case 5:
		dst := md5.Sum(s)
		res = dst[:]
	default:
		return "", fmt.Errorf("not found md%d", typ)
	}
	if len(isRaw) > 0 && isRaw[0] {
		return string(res), nil
	}
	return hex.EncodeToString(res), nil
}
