package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/dgrijalva/jwt-go"
)

// PrintJWT - Print JWT
func PrintJWT(s string, secret ...string) (string, error) {
	secretByte := []byte("")
	if len(secret) > 0 {
		secretByte = []byte(secret[0])
	}
	token, _ := jwt.Parse(strings.TrimSpace(s), func(token *jwt.Token) (interface{}, error) {
		return secretByte, nil
	})
	res, err := json.MarshalIndent(token, "", "    ")
	// if claims, ok := token.Claims.(jwt.MapClaims); ok {}
	return string(res), err
}

// CrackJWT - Crack JWT
// args: [start=4] [end=4] [table=alphabet&number] [prefix] [suffix]
func CrackJWT(s string, args ...interface{}) (string, error) {
	tokenStr := strings.TrimSpace(s)
	if t, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("dbedd"), nil
	}); t == nil {
		return "", fmt.Errorf("token is malformed")
	}
	var (
		start  = 4
		end    = 4
		table  = []byte("abcdefghijklnmopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
		prefix = []byte("")
		suffix = []byte("")
	)
	if len(args) > 0 && args[0].(int) <= 8 {
		start = args[0].(int)
		if start > end {
			end = start
		}
	}
	if len(args) > 1 && args[1].(int) <= 8 && args[1].(int) >= start {
		end = args[1].(int)
	}
	if len(args) > 2 && len(args[2].(string)) >= end {
		table = []byte(args[2].(string))
	}
	if len(args) > 3 && len(args[3].(string)) > 0 {
		prefix = []byte(args[3].(string))
	}
	if len(args) > 4 && len(args[4].(string)) > 0 {
		suffix = []byte(args[4].(string))
	}
	var res = ""
	done := make(chan struct{})
	wg := &sync.WaitGroup{}

	for _, ct := range table {
		wg.Add(1)
		go func(ct byte) {
			var helper func([]byte)
			helper = func(secret []byte) {
				select {
				case <-done:
					return
				default:
					// if len(secret) > 1 && secret[0] == 'c' && secret[1] == 'b' {
					// 	fmt.Println(string(secret))
					// }
					_secret := append(append(prefix, secret...), suffix...)
					if len(secret) >= start {
						tt, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
							return _secret, nil
						})
						// if err != nil && err != jwt.ErrSignatureInvalid {
						// 	fmt.Fprintln(os.Stderr, err)
						// }
						if tt.Valid {
							res = string(_secret)
							close(done)
							return
						}
					}
					if len(secret) < end {
						for _, cc := range table {
							helper(append(secret, cc))
						}
					}
				}
			}
			// fmt.Println("go work", string(ct))
			helper([]byte{ct})
			// fmt.Println("gone work", string(ct))
			wg.Done()
		}(ct)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)
		for {
			select {
			case <-done:
				return
			case <-interrupt:
				close(done)
				return
			}
		}
	}()

	wg.Wait()
	if len(res) > 0 {
		return res, nil
	}
	return res, fmt.Errorf("No secret found")
}

// ModifyJWT - ModifyJ JWT
func ModifyJWT(s string, none bool, secret string, claims map[string]string, method string) (string, error) {
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
