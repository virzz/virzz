package crypto

import (
	"fmt"
	"strings"
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
		r := strings.Title(strings.ToLower(string(res)))
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
