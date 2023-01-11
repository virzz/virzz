package jwttool

import (
	"bytes"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/dgrijalva/jwt-go"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/utils"
)

func crackJWTV2(s string, minLen, maxLen int, alphabet, prefix, suffix []byte) (string, error) {
	count := utils.CalcPermutationMore(len(alphabet), minLen, maxLen)
	logger.WarnF("Total crack count: %d", count)
	secrets := make(chan []byte, 100000)
	done := make(chan struct{})

	var insertTable func([]byte, byte, int, int)
	insertTable = func(prefix []byte, c byte, m, l int) {
		if len(prefix) < m {
			for _, cc := range alphabet {
				insertTable(append(prefix, c), cc, m, l)
			}
			return
		}
		if len(prefix) >= l {
			return
		}
		secrets <- append(prefix, c)
	}

	go func() {
		for i := range alphabet {
			insertTable([]byte{}, alphabet[i], minLen, maxLen)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	tokenStr := strings.TrimSpace(s)
	res := ""

	for {
		select {
		case s := <-secrets:
			logger.Info(s)
			_secret := bytes.Buffer{}
			if prefix != nil {
				_secret.Write(prefix)
			}
			_secret.Write(s)
			if suffix != nil {
				_secret.Write(suffix)
			}
			tt, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				return _secret.Bytes(), nil
			})
			if tt.Valid {
				done <- struct{}{}
				close(secrets)
				close(done)
				close(interrupt)
				logger.Debug("success")
				return _secret.String(), nil
			}
		case <-done:
			logger.Debug("done")
			close(done)
		case <-interrupt:
			logger.Debug("interrupt")
			close(done)
		}
	}

	if len(res) > 0 {
		return res, nil
	}
	return res, fmt.Errorf("no secret found")
}
