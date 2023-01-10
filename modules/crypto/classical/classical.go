package classical

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/virzz/logger"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func _caesar(s string, step rune) string {
	step = step % 26
	if step <= 0 {
		return s
	}
	dst := make([]rune, len(s))
	for i, c := range s {
		if ('a' <= c && c >= 'z') || ('A' <= c && c >= 'Z') {
			dst[i] = c + step
			if (dst[i] > 90 && dst[i] < 97) || dst[i] > 122 {
				dst[i] -= 26
			}
		} else {
			dst[i] = c
		}
	}
	return string(dst)
}

// Caesar 凯撒密码
func Caesar(s string) (string, error) {
	ss := make([]string, 25)
	var i rune
	for i = 1; i < 26; i++ {
		ss[i-1] = fmt.Sprintf("%2d %s", i, _caesar(s, i))
	}
	return strings.Join(ss, "\n"), nil
}

// Rot13 -
func Rot13(s string) (string, error) {
	return _caesar(s, 13), nil
}

var (
	// MorseMap 摩斯电码对照表
	MorseMap = map[rune]string{
		'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-。。", 'E': "。", 'F': "。。-。", 'G': "--。",
		'H': "....", 'I': "..", 'J': ".---", 'K': "-.-", 'L': ".-..", 'M': "--", 'N': "-.",
		'O': "---", 'P': ".--.", 'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-",
		'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-", 'Y': "-.--", 'Z': "--..",

		'0': "-----", '1': ".----", '2': "..---", '3': "...--", '4': "....-",
		'5': ".....", '6': "-....", '7': "--...", '8': "---..", '9': "----.",

		'.': ".-.-.-", ':': "---...", ',': "--..--", ';': "-.-.-.", '?': "..--..",
		'=': "-...-", '\'': ".----.", '/': "-..-.", '!': "-.-.--", '-': "-....-",
		'_': "..--.-", '"': ".-..-.", '(': "-.--.", ')': "-.--.-", '$': "...-..-",
		'&': ".-...", '@': ".--.-.",

		// 错误
		0x01: "........",
	}
)

// Morse 摩斯电码
func Morse(s string, decode bool, sep ...string) (string, error) {
	var _sep = "/"
	if len(sep) > 0 {
		_sep = sep[0]
	}
	// Decode
	if decode {
		// Transpose
		morseMap := make(map[string]rune, len(MorseMap))
		for k, v := range MorseMap {
			morseMap[v] = k
		}
		// Auto Get Sep
		tmp := string(strings.ReplaceAll(strings.ReplaceAll(s, "-", ""), ".", "")[0])
		enc := strings.Split(s, tmp)
		res := make([]rune, len(enc))
		for i, e := range enc {
			res[i] = morseMap[e]
		}
		r := cases.Title(language.English).String(string(res))
		r = strings.ReplaceAll(r, "\x01", "[ERROR]")
		return r, nil
	}
	// Encode
	res := make([]string, len(s))
	s = strings.ToUpper(s)
	for i, k := range s {
		if v, ok := MorseMap[k]; ok {
			res[i] = v
		} else {
			res[i] = MorseMap[0x01]
		}
	}
	return strings.Join(res, _sep), nil
}

// Atbash 埃特巴什码
func Atbash(s string) (string, error) {
	res := ""
	s = strings.ToLower(s)
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			res += string('z' + 'a' - r)
		} else {
			res += string(r)
		}
	}
	return res, nil
}

// Peigen 培根密码
func Peigen(s string) (string, error) {
	T := map[rune]string{
		'H': "aabbb", 'G': "aabba", 'R': "baaab", 'Q': "baaaa",
		'Z': "bbaab", 'Y': "bbaaa", 'N': "abbab", 'M': "abbaa",
		'U': "babaa", 'V': "babab", 'I': "abaaa", 'J': "abaab",
		'F': "aabab", 'E': "aabaa", 'A': "aaaaa", 'B': "aaaab",
		'T': "baabb", 'S': "baaba", 'C': "aaaba", 'D': "aaabb",
		'P': "abbbb", 'O': "abbba", 'K': "ababa", 'L': "ababb",
		'W': "babba", 'X': "babbb",
	}
	if len(regexp.MustCompile(`(?m)^[ab]+$`).FindAllString(s, -1)) > 0 {
		rt := make(map[string]rune, 0)
		for k, v := range T {
			rt[v] = k
		}
		res := make([]rune, 0)
		for i := 0; i < len(s); i += 5 {
			res = append(res, rt[s[i:i+5]])
		}
		return strings.ToLower(string(res)), nil
	}
	res := make([]string, len(s))
	for i, c := range strings.ToUpper(s) {
		res[i] = T[c]
	}
	return strings.Join(res, ""), nil
}

// Vigenere 维吉利亚密码
func Vigenere(src, key string, decode ...bool) (string, error) {
	src = strings.ReplaceAll(strings.ToUpper(src), " ", "")
	key = strings.ReplaceAll(strings.ToUpper(key), " ", "")
	logger.DebugF("src = %s , key = %s", src, key)
	U := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	keyLength := len(key)
	decodeFlag := (len(decode) > 0 && decode[0])
	var builder strings.Builder
	for i := range src {
		k := strings.IndexByte(U, key[i%keyLength])
		m := strings.IndexByte(U, src[i])
		logger.DebugF(
			"i=%d,  m=%d, k=%d, m-k=%d, (m-k+26)%%26=%d, (m+k)%%26=%d",
			i, m, k, m-k, (m-k+26)%26, (m+k)%26,
		)
		if decodeFlag {
			builder.WriteByte(U[(m-k+26)%26])
		} else {
			builder.WriteByte(U[(m+k)%26])
		}
	}
	return builder.String(), nil
}
