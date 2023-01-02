package utils

import (
	"math/rand"
	"regexp"
)

var defaultLetters = []rune("0123456789abcdefghijklmnopqrstuvwxyz")

// RandomString 随机字符串
func RandomString() string {
	return RandomStringByLength(8)
}

// RandomStringByLength 随机字符串
func RandomStringByLength(n int, allowedChars ...[]rune) string {
	var letters []rune
	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func AnsiStrip(str string) string {
	const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"
	return regexp.MustCompile(ansi).ReplaceAllString(str, "")
}
