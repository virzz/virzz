package utils

import (
	"math/rand"
	"regexp"
)

var defaultLetters = []byte("0123456789abcdefghijklmnopqrstuvwxyz")

// RandomString 随机字符串
func RandomString() string {
	return RandomStringByLength(8)
}

// RandomBytesByLength 随机字符串
func RandomBytesByLength(n int, allowedChars ...[]byte) []byte {
	var letters []byte
	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}
	l := len(letters)
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(l)]
	}
	return b
}

// RandomStringByLength 随机字符串
func RandomStringByLength(n int, allowedChars ...[]byte) string {
	return string(RandomBytesByLength(n, allowedChars...))
}

func AnsiStrip(str string) string {
	const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"
	return regexp.MustCompile(ansi).ReplaceAllString(str, "")
}
