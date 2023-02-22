package jwttool

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/utils"
	"github.com/virzz/virzz/utils/pool"
)

const defaultAlphabet = "abcdefghijklnmopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// JWTPrint Print JWT
func JWTPrint(token string, secret ...string) (string, error) {
	secretByte := []byte("")
	if len(secret) > 0 {
		secretByte = []byte(secret[0])
	}
	_token, err := jwt.Parse(strings.TrimSpace(token), func(*jwt.Token) (interface{}, error) {
		return secretByte, nil
	})
	if err != nil && err.Error() != jwt.ErrSignatureInvalid.Error() {
		return "", err
	}
	res, err := json.MarshalIndent(_token, "", "    ")
	return string(res), err
}

// JWTModify Modify JWT
func JWTModify(s string, none bool, secret string, claims map[string]string, method string) (string, error) {
	var t *jwt.Token
	newClaims := jwt.MapClaims{}
	if t, _ = jwt.ParseWithClaims(strings.TrimSpace(s), newClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}); t == nil {
		return "", fmt.Errorf("token is malformed")
	}
	for k, v := range claims {
		newClaims[k] = v
	}
	if none {
		t.Header["alg"] = "none"
		h, err := json.Marshal(t.Header)
		if err != nil {
			return "", err
		}
		c, err := json.Marshal(newClaims)
		if err != nil {
			return "", err
		}
		b64e := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString
		return fmt.Sprintf("%s.%s.", b64e(h), b64e(c)), nil
	}

	var newToken *jwt.Token
	switch method {
	case "HS256":
		newToken = jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	case "HS384":
		newToken = jwt.NewWithClaims(jwt.SigningMethodHS384, newClaims)
	case "HS512":
		newToken = jwt.NewWithClaims(jwt.SigningMethodHS512, newClaims)
	default:
		return "", fmt.Errorf("the method %s is not support", method)
	}
	tokenString, err := newToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// JWTCreate - Creat JWT
func JWTCreate(claims map[string]string, method string, secret string) (string, error) {
	newClaims := jwt.MapClaims{}
	for k, v := range claims {
		newClaims[k] = v
	}
	var newToken *jwt.Token
	switch method {
	case "HS256":
		newToken = jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	case "HS384":
		newToken = jwt.NewWithClaims(jwt.SigningMethodHS384, newClaims)
	case "HS512":
		newToken = jwt.NewWithClaims(jwt.SigningMethodHS512, newClaims)
	default:
		return "", fmt.Errorf("the method %s is not support", method)
	}
	tokenString, err := newToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// JWTCrack Crack JWT
func JWTCrack(s string, minLen, maxLen int, alphabet, prefix, suffix string) (string, error) {
	tokenStr := strings.TrimSpace(s)
	alphabetByte := []byte(alphabet)

	done := make(chan struct{}, 1)
	result := make(chan []byte, 1)

	count := utils.CalcPermutationMore(len(alphabet), minLen, maxLen)
	logger.WarnF("Total crack count: %d", count)

	var _crack func([]byte)
	_crack = func(secret []byte) {
		select {
		case <-done:
			return
		default:
			_secret := bytes.Buffer{}
			if len(prefix) > 0 {
				_secret.WriteString(prefix)
			}
			_secret.Write(secret)
			if len(suffix) > 0 {
				_secret.WriteString(suffix)
			}
			secretByte := _secret.Bytes()

			if len(secretByte) >= minLen {
				t, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
					return secretByte, nil
				})
				if t.Valid {
					result <- secretByte
					done <- struct{}{}
					return
				}
			}
			if len(secretByte) < maxLen {
				for _, cc := range alphabetByte {
					_crack(append(secret, cc))
				}
			}
		}
	}

	var doCrackV2 = func(arg string) bool {
		_crack([]byte(arg))
		return false
	}
	args := make([]string, len(alphabet))
	for i := range alphabet {
		args[i] = string(alphabet[i])
	}

	pool.Start(doCrackV2, args...)

	select {
	case r := <-result:
		close(result)
		code := string(r)
		logger.SuccessF("JWT Secret Cracked: %s", code)
		return code, nil
	case <-done:
		close(done)
		return "", fmt.Errorf("crack failed")
	}
}
