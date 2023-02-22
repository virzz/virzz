package basex

import (
	"encoding/ascii85"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/virzz/virzz/internal/crypto/base100"
	"github.com/virzz/virzz/internal/crypto/base36"
	"github.com/virzz/virzz/internal/crypto/base58"
	"github.com/virzz/virzz/internal/crypto/base62"
	"github.com/virzz/virzz/internal/crypto/base91"
	"github.com/virzz/virzz/internal/crypto/base92"
)

// basePadding RFC 4648
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

// base16Encode hex.EncodeToString
func Base16Encode(s string) (string, error) {
	return hex.EncodeToString([]byte(strings.TrimSuffix(s, "\n"))), nil
}

// base16Decode hex.DecodeString
func Base16Decode(s string) (string, error) {
	out, err := hex.DecodeString(strings.TrimSuffix(s, "\n"))
	if err != nil {
		return "", fmt.Errorf("failed to decode input: %w", err)
	}
	return string(out), nil
}

// base32Encode -
func Base32Encode(s string) (string, error) {
	return base32.StdEncoding.EncodeToString([]byte(s)), nil
}

// base32Decode -
func Base32Decode(s string) (string, error) {
	s = basePadding(strings.TrimSpace(s), 8)
	res, err := base32.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// base36Encode -
func Base36Encode(s string) (string, error) {
	return base36.Encode(s), nil
}

// base36Decode -
func Base36Decode(s string) (string, error) {
	return base36.Decode(strings.TrimSpace(s)), nil
}

// base58Encode -
func Base58Encode(s string, enc string) (string, error) {
	switch enc {
	case "flickr":
		return base58.FlickrEncoding.EncodeToString([]byte(s)), nil
	case "ripple":
		return base58.RippleEncoding.EncodeToString([]byte(s)), nil
	}
	return base58.BitcoinEncoding.EncodeToString([]byte(s)), nil
}

// base58Decode enc = <flickr|ripple|bitcoin>
func Base58Decode(s string, enc string) (string, error) {
	var (
		res []byte
		err error
	)
	s = strings.TrimSpace(s)
	switch enc {
	case "flickr":
		res, err = base58.FlickrEncoding.DecodeString(s)
	case "ripple":
		res, err = base58.RippleEncoding.DecodeString(s)
	case "", "bitcoin":
		res, err = base58.BitcoinEncoding.DecodeString(s)
	default:
		err = fmt.Errorf("unknown encoder: %s", enc)
	}
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// base62Encode
func Base62Encode(s string) (string, error) {
	return string(base62.StdEncoding.Encode([]byte(s))), nil
}

// base62Decode -
func Base62Decode(s string) (string, error) {
	res, err := base62.StdEncoding.Decode([]byte(s))
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// base64Encode Base64 Encode
func Base64Encode(s string, url bool) (string, error) {
	if url {
		return base64.URLEncoding.EncodeToString([]byte(s)), nil
	}
	return base64.StdEncoding.EncodeToString([]byte(s)), nil
}

// base64Decode -
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

// base85Encode -
func Base85Encode(s string) (string, error) {
	bs := []byte(s)
	dst := make([]byte, ascii85.MaxEncodedLen(len(bs)))
	n := ascii85.Encode(dst, bs)
	return string(dst[:n]), nil
}

// base85Decode -
func Base85Decode(s string) (string, error) {
	bs := []byte(strings.TrimSpace(s))
	dst := make([]byte, len(bs)*4)
	n, _, err := ascii85.Decode(dst, bs, true)
	if err != nil {
		return "", err
	}
	return string(dst[:n]), nil
}

// base91Encode
func Base91Encode(s string) (string, error) {
	return base91.StdEncoding.EncodeToString([]byte(s)), nil
}

// base91Decode -
func Base91Decode(s string) (string, error) {
	res, err := base91.StdEncoding.DecodeString(strings.TrimSpace(s))
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// base92Encode
func Base92Encode(s string) (string, error) {
	return base92.StdEncoding.EncodeToString([]byte(s)), nil
}

// base92Decode -
func Base92Decode(s string) (string, error) {
	res, err := base92.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// base100Encode
func Base100Encode(s string) (string, error) {
	return base100.Encode([]byte(s)), nil
}

// base100Decode -
func Base100Decode(s string) (string, error) {
	res, err := base100.Decode(s)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
