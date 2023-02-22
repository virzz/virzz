package nettool

import (
	"fmt"
	"net"
	"strings"

	"github.com/virzz/virzz/modules/crypto/basic"
)

// MacToDec -
func Mac2Dec(s string) (string, error) {
	_, err := net.ParseMAC(s)
	if err != nil {
		return "", err
	}
	for _, c := range []string{":", "-", "."} {
		s = strings.ReplaceAll(s, c, "")
	}
	return basic.HexToDecStr(strings.Trim(s, ":-."))
}

// Dec2MAC -
func Dec2MAC(s string) (string, error) {
	h, err := basic.DecToHex(s)
	if err != nil {
		return "", err
	}
	h = h[2:]
	for len(h) < 12 {
		h = "0" + h
	}
	l := len(h)
	if l == 12 || l == 16 || l == 40 {
		r := make([]string, l/2)
		for i := range r {
			r[i] = h[i*2 : (i+1)*2]
		}
		mac := strings.Join(r, ":")
		_, err := net.ParseMAC(mac)
		if err != nil {
			return "", err
		}
		return mac, nil
	}
	return "", fmt.Errorf("invalid MAC address")
}
