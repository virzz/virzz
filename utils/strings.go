package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"sync"

	"golang.org/x/exp/constraints"
)

var (
	lock          sync.Mutex
	alphabetCache = make(map[string][]byte)
)

func GenerateAlphabet(regex string) []byte {
	if v, ok := alphabetCache[regex]; ok {
		return v
	}
	if strings.HasPrefix(regex, "[") && strings.HasSuffix(regex, "]") {
		regex = regex[1 : len(regex)-1]
	}
	letters := make([]byte, 256)
	for i := 0; i < 256; i++ {
		letters[i] = byte(i)
	}
	lock.Lock()
	alphabetCache[regex] = regexp.MustCompile(fmt.Sprintf(`[^%s]`, regex)).ReplaceAll(letters, nil)
	lock.Unlock()
	return alphabetCache[regex]
}

// RandomBytesByLength 随机字符串
func RandomBytesByLength[T constraints.Integer](n T, regex ...string) []byte {
	var letters []byte
	if len(regex) > 0 && len(regex[0]) > 0 {
		letters = GenerateAlphabet(regex[0])
	}
	if len(letters) == 0 {
		letters = GenerateAlphabet(`a-z0-9`)
	}
	l := len(letters)
	if n == 0 {
		n = 8
	}
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(l)]
	}
	return b
}

// RandomStringByLength 随机字符串
func RandomStringByLength[T constraints.Integer](n T, regex ...string) string {
	return string(RandomBytesByLength(n, regex...))
}

// RandomString 随机字符串
func RandomString() string {
	return RandomStringByLength(8)
}

func AnsiStrip(str string) string {
	const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"
	return regexp.MustCompile(ansi).ReplaceAllString(str, "")
}
