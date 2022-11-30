package utils

import (
	"math/rand"
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
