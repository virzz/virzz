package basic

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

// // StringX -
// type StringX struct{}

// UpPadHex Remove 0x
func UpPadHex(s string) string {
	s = strings.ToLower(s)
	if strings.HasPrefix(s, "0x") {
		s = s[2:]
	}
	return s
}

// PadHex Append 0x
func PadHex(s string) string {
	return fmt.Sprintf("0x%s", s)
}

// StringToASCII 字符串 -> ASCII
func StringToASCII(s string) (string, error) {
	var res = make([]string, len(s))
	var (
		i int
		c rune
	)
	for i, c = range []rune(s) {
		res[i] = strconv.Itoa(int(c))
	}
	if i < len(s) {
		res = res[:i+1]
	}
	return strings.Join(res, ","), nil
}

// ASCIIToString ASCII -> 字符串
func ASCIIToString(s string) (string, error) {
	ss := strings.Split(strings.TrimSpace(s), ",")
	l := len(ss)
	res := make([]string, l)
	for i, c := range ss {
		a, err := strconv.Atoi(c)
		if err != nil {
			res[i] = "?"
		} else {
			res[i] = string(rune(a))
		}
	}
	return strings.Join(res, ""), nil
}

// HexToString Hex -> String
func HexToString(s string) (string, error) {
	s = UpPadHex(s)
	bs, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

// StringToHex String -> Hex
func StringToHex(s string) (string, error) {
	return PadHex(hex.EncodeToString([]byte(s))), nil
}

// DecToHex Dec -> Hex
func DecToHex(s string) (string, error) {
	n := new(big.Int)
	var ok bool
	if n, ok = n.SetString(s, 10); !ok {
		return "", fmt.Errorf("Convert error")
	}
	return PadHex(n.Text(16)), nil
}

// HexToDec Hex -> Dec
func HexToDec(s string) (string, error) {
	s = UpPadHex(s)
	n := new(big.Int)
	var ok bool
	if n, ok = n.SetString(s, 16); !ok {
		return "", fmt.Errorf("Convert error")
	}
	return n.String(), nil
}

// HexToByteString Hex -> Bytes String
func HexToByteString(s string) (string, error) {
	s = UpPadHex(s)
	bs, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}
	res := make([]string, len(bs))
	for i, b := range bs {
		// 可见字符
		if 0x20 <= b && b <= 0x7E {
			res[i] = string(b)
		} else {
			res[i] = fmt.Sprintf("\\x%x", b)
		}
	}
	return fmt.Sprintf("b'%s'", strings.Join(res, "")), nil
}

// ByteStringToHex ByteString -> Hex
func ByteStringToHex(s string) (string, error) {
	var re = regexp.MustCompile(`^b["']([\S|(\\x\w{2})]*?)['"]$`)
	var matches = re.FindAllStringSubmatch(s, -1)
	if len(matches) == 0 || len(matches[0]) < 2 {
		return "", fmt.Errorf("Regexp Match error")
	}
	m := matches[0][1]
	res := ""
	for p := 0; p < len(m); {
		if m[p] == '\\' && m[p+1] == 'x' {
			res += m[p+2 : p+4]
			p += 4
			continue
		}
		res += hex.EncodeToString([]byte{m[p]})
		p++
	}
	return PadHex(res), nil
}

// ByteStringToString ByteString -> String
func ByteStringToString(s string) (string, error) {
	s, err := ByteStringToHex(s)
	if err != nil {
		return "", err
	}
	return HexToString(s)
}
