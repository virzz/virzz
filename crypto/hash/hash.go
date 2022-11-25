package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func _md5Hash(s []byte) string {
	has := md5.Sum(s)
	return hex.EncodeToString(has[:])
}

func _sha1Hash(s []byte) string {
	has := sha1.Sum(s)
	return hex.EncodeToString(has[:])
}

func md5Hash(s string) (string, error) {
	return _md5Hash([]byte(s)), nil
}

func sha1Hash(s string) (string, error) {
	return _sha1Hash([]byte(s)), nil
}

func sha224Hash(s string) (string, error) {
	has := sha256.Sum224([]byte(s))
	return hex.EncodeToString(has[:]), nil
}

func sha256Hash(s string) (string, error) {
	has := sha256.Sum256([]byte(s))
	return hex.EncodeToString(has[:]), nil
}

func sha384Hash(s string) (string, error) {
	has := sha512.Sum384([]byte(s))
	return hex.EncodeToString(has[:]), nil
}

func sha512Hash(s string) (string, error) {
	has := sha512.Sum512([]byte(s))
	return hex.EncodeToString(has[:]), nil
}

func sha512_224Hash(s string) (string, error) {
	has := sha512.Sum512_224([]byte(s))
	return hex.EncodeToString(has[:]), nil
}

func sha512_256Hash(s string) (string, error) {
	has := sha512.Sum512_256([]byte(s))
	return hex.EncodeToString(has[:]), nil
}
