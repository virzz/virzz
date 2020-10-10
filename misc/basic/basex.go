package basic

import (
	"encoding/base32"
	"encoding/base64"
	"strings"

	"github.com/virink/virzz/misc/basic/base58"
)

// RFC 4648

func basePadding(s string, bit int) string {
	n := bit - len(s)%bit
	if n == bit {
		return s
	}
	for n > 0 {
		n--
		s += "="
	}
	return s
}

// Base64Encode Base64 Encode
func Base64Encode(s string, safe ...bool) (string, error) {
	if len(safe) > 0 && safe[0] {
		return base64.URLEncoding.EncodeToString([]byte(s)), nil
	}
	return base64.StdEncoding.EncodeToString([]byte(s)), nil
}

// Base64Decode -
func Base64Decode(s string) (string, error) {
	s = basePadding(strings.TrimSpace(s), 4)
	res, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		res, err = base64.URLEncoding.DecodeString(s)
		if err != nil {
			return "", err
		}
	}
	return string(res), nil
}

// Base32Encode -
func Base32Encode(s string) (string, error) {
	return base32.StdEncoding.EncodeToString([]byte(s)), nil
}

// Base32Decode -
func Base32Decode(s string) (string, error) {
	s = basePadding(strings.TrimSpace(s), 8)
	res, err := base32.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// FuzzBase64Table -
// TODO: FuzzBase64Table
func FuzzBase64Table(cipher, plain string) (string, error) {
	return "", nil
}

// TODO: Base 16 36 58 91 92

// Base58Encode -
func Base58Encode(s string, enc ...string) (string, error) {
	if len(enc) > 0 {
		if enc[0] == "flickr" {
			return base58.FlickrEncoding.EncodeToString([]byte(s)), nil
		} else if enc[0] == "ripple" {
			return base58.RippleEncoding.EncodeToString([]byte(s)), nil
		}
	}
	return base58.BitcoinEncoding.EncodeToString([]byte(s)), nil
}

// Base58Decode -
func Base58Decode(s string, enc ...string) (string, error) {
	var (
		res []byte
		err error
	)
	s = strings.TrimSpace(s)
	if len(enc) > 0 {
		if enc[0] == "flickr" {
			res, err = base58.FlickrEncoding.DecodeString(s)
		} else if enc[0] == "ripple" {
			res, err = base58.RippleEncoding.DecodeString(s)
		}
	}
	res, err = base58.BitcoinEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
