package basex

import (
	"encoding/ascii85"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/virzz/virzz/modules/crypto/internal/base100"
	"github.com/virzz/virzz/modules/crypto/internal/base36"
	"github.com/virzz/virzz/modules/crypto/internal/base58"
	"github.com/virzz/virzz/modules/crypto/internal/base62"
	"github.com/virzz/virzz/modules/crypto/internal/base91"
	"github.com/virzz/virzz/modules/crypto/internal/base92"
)

// base16Encode hex.EncodeToString
func base16Encode(s string) (string, error) {
	return hex.EncodeToString([]byte(strings.TrimSuffix(s, "\n"))), nil
}

// base16Decode hex.DecodeString
func base16Decode(s string) (string, error) {
	out, err := hex.DecodeString(strings.TrimSuffix(s, "\n"))
	if err != nil {
		return "", fmt.Errorf("failed to decode input: %w", err)
	}
	return string(out), nil
}

// base32Encode -
func base32Encode(s string) (string, error) {
	return base32.StdEncoding.EncodeToString([]byte(s)), nil
}

// base32Decode -
func base32Decode(s string) (string, error) {
	s = basePadding(strings.TrimSpace(s), 8)
	res, err := base32.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// base36Encode -
func base36Encode(s string) (string, error) {
	return base36.Encode(s), nil
}

// base36Decode -
func base36Decode(s string) (string, error) {
	return base36.Decode(strings.TrimSpace(s)), nil
}

// base58Encode -
func base58Encode(s string, enc ...string) (string, error) {
	if len(enc) > 0 {
		if enc[0] == "flickr" {
			return base58.FlickrEncoding.EncodeToString([]byte(s)), nil
		} else if enc[0] == "ripple" {
			return base58.RippleEncoding.EncodeToString([]byte(s)), nil
		}
	}
	return base58.BitcoinEncoding.EncodeToString([]byte(s)), nil
}

// base58Decode -
func base58Decode(s string, enc ...string) (string, error) {
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
		} else {
			return "", fmt.Errorf("enc is not support")
		}
	} else {
		res, err = base58.BitcoinEncoding.DecodeString(s)
	}
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// base62Encode
func base62Encode(s string) (string, error) {
	return string(base62.StdEncoding.Encode([]byte(s))), nil
}

// base62Decode -
func base62Decode(s string) (string, error) {
	res, err := base62.StdEncoding.Decode([]byte(s))
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// base64Encode Base64 Encode
func base64Encode(s string, safe ...bool) (string, error) {
	if len(safe) > 0 && safe[0] {
		return base64.URLEncoding.EncodeToString([]byte(s)), nil
	}
	return base64.StdEncoding.EncodeToString([]byte(s)), nil
}

// base64Decode -
func base64Decode(s string) (string, error) {
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

// base85Encode -
func base85Encode(s string) (string, error) {
	bs := []byte(s)
	dst := make([]byte, ascii85.MaxEncodedLen(len(bs)))
	n := ascii85.Encode(dst, bs)
	return string(dst[:n]), nil
}

// base85Decode -
func base85Decode(s string) (string, error) {
	bs := []byte(strings.TrimSpace(s))
	dst := make([]byte, len(bs)*4)
	n, _, err := ascii85.Decode(dst, bs, true)
	if err != nil {
		return "", err
	}
	return string(dst[:n]), nil
}

// base91Encode
func base91Encode(s string) (string, error) {
	return base91.StdEncoding.EncodeToString([]byte(s)), nil
}

// base91Decode -
func base91Decode(s string) (string, error) {
	res, err := base91.StdEncoding.DecodeString(strings.TrimSpace(s))
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// base92Encode
func base92Encode(s string) (string, error) {
	return base92.StdEncoding.EncodeToString([]byte(s)), nil
}

// base92Decode -
func base92Decode(s string) (string, error) {
	res, err := base92.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// base100Encode
func base100Encode(s string) (string, error) {
	return base100.Encode([]byte(s)), nil
}

// base100Decode -
func base100Decode(s string) (string, error) {
	res, err := base100.Decode(s)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
