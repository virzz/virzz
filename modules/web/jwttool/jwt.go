package jwttool

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/golang-jwt/jwt/v4"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/utils"
)

func printJWT(s string, secret ...string) (string, error) {
	secretByte := []byte("")
	if len(secret) > 0 {
		secretByte = []byte(secret[0])
	}
	token, err := jwt.Parse(strings.TrimSpace(s), func(token *jwt.Token) (interface{}, error) {
		return secretByte, nil
	})
	if err != nil && err.Error() != jwt.ErrSignatureInvalid.Error() {
		return "", err
	}
	res, err := json.MarshalIndent(token, "", "    ")
	return string(res), err
}

// crackJWT - Crack JWT
func crackJWT(s string, minLen, maxLen int, alphabet, prefix, suffix []byte) (string, error) {
	tokenStr := strings.TrimSpace(s)
	var res = ""
	done := make(chan struct{})
	wg := &sync.WaitGroup{}
	var _crack func([]byte)
	_crack = func(secret []byte) {
		wg.Add(1)
		defer wg.Done()
		select {
		case <-done:
			return
		default:
			if len(secret) >= minLen {
				_secret := bytes.Buffer{}
				if prefix != nil {
					_secret.Write(prefix)
				}
				_secret.Write(secret)
				if suffix != nil {
					_secret.Write(suffix)
				}
				tt, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
					return _secret.Bytes(), nil
				})
				if tt.Valid {
					res = _secret.String()
					close(done)
					return
				}
			}
			if len(secret) < maxLen {
				for _, cc := range alphabet {
					_crack(append(secret, cc))
				}
			}
		}
	}

	count := utils.CalcPermutationMore(len(alphabet), minLen, maxLen)
	logger.WarnF("Total crack count: %d", count)

	for _, ct := range alphabet {
		go _crack([]byte{ct})
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
		for {
			select {
			case <-done:
				return
			case <-interrupt:
				logger.Debug("interrupt")
				close(done)
				return
			}
		}
	}()

	wg.Wait()
	if len(res) > 0 {
		return res, nil
	}
	return res, fmt.Errorf("no secret found")
}

func modifyJWT(s string, none bool, secret string, claims map[string]string, method string) (string, error) {
	tokenStr := strings.TrimSpace(s)
	var t *jwt.Token
	newClaims := jwt.MapClaims{}
	if t, _ = jwt.ParseWithClaims(tokenStr, newClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte("dbedd"), nil
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
	// case "RS256":
	// 	newToken = jwt.New(&jwt.SigningMethodRSA{})
	default:
		return "", fmt.Errorf("the method %s is not support", method)
	}
	tokenString, err := newToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// CreatJWT - CreatJ JWT
// NOTE: Maybe not need
func CreatJWT(s string) (string, error) {
	return "", nil
}
