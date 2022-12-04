package network

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/mozhu1024/virzz/modules/misc/basic"
)

// https://github.com/projectdiscovery/ipranger

var errParseIP = fmt.Errorf("parse ip faild")

func inet4aton(ipnr net.IP) int64 {
	var sum int64
	bit := 24
	for _, p := range strings.Split(ipnr.String(), ".") {
		i, _ := strconv.Atoi(p)
		sum += int64(i) << bit
		bit -= 8
	}
	return sum
}
func inet4ntoa(ipnr int64) net.IP {
	return net.IPv4(
		byte((ipnr>>24)&0xFF),
		byte((ipnr>>16)&0xFF),
		byte((ipnr>>8)&0xFF),
		byte(ipnr&0xFF),
	)
}

// ipToDec 127.0.0.1 -> 2130706433
func ipToDec(s string) (string, error) {
	ip := net.ParseIP(s)
	if ip == nil {
		return "", errParseIP
	}
	return strconv.FormatInt(inet4aton(ip), 10), nil
}

// ipToOct 127.0.0.1 -> 2130706433
func ipToOct(s string) (string, error) {
	ip := net.ParseIP(s)
	if ip == nil {
		return "", errParseIP
	}
	return "0" + strconv.FormatInt(inet4aton(ip), 8), nil
}

// ipToHex 127.0.0.1 -> 0x7f000001
func ipToHex(s string) (string, error) {
	s, err := ipToDec(s)
	if err != nil {
		return "", err
	}
	return basic.DecToHex(s)
}

// ipToDotOct 127.0.0.1 -> 0x7f.0.0.0x1
func ipToDotOct(s string) (string, error) {
	ip := net.ParseIP(s)
	if ip == nil {
		return "", errParseIP
	}
	rs := make([]string, 0)
	for _, p := range strings.Split(s, ".") {
		i, _ := strconv.Atoi(p)
		if i > 0 {
			if i > 7 {
				rs = append(rs, "0"+strconv.FormatInt(int64(i), 8))
			} else {
				rs = append(rs, strconv.FormatInt(int64(i), 8))
			}
		} else {
			rs = append(rs, "0")
		}
	}
	return strings.Join(rs, "."), nil
}

// ipToDotHex 127.0.0.1 -> 0x7f.0.0.0x1
func ipToDotHex(s string) (string, error) {
	ip := net.ParseIP(s)
	if ip == nil {
		return "", errParseIP
	}
	rs := make([]string, 0)
	for _, p := range strings.Split(s, ".") {
		i, _ := strconv.Atoi(p)
		if i > 0 {
			if i > 9 {
				rs = append(rs, "0x"+strconv.FormatInt(int64(i), 16))
			} else {
				rs = append(rs, strconv.FormatInt(int64(i), 16))
			}
		} else {
			rs = append(rs, "0")
		}
	}
	return strings.Join(rs, "."), nil
}

// ipToAll -
func ipToAll(s string) (string, error) {
	rs := make([]string, 0)
	// oct
	if r, err := ipToOct(s); err == nil {
		rs = append(rs, fmt.Sprintf("Oct:    %s", r))
	}
	// dec
	if r, err := ipToDec(s); err == nil {
		rs = append(rs, fmt.Sprintf("Dec:    %s", r))
	}
	// hex
	if r, err := ipToHex(s); err == nil {
		rs = append(rs, fmt.Sprintf("Hex:    %s", r))
	}
	// dot oct
	if r, err := ipToDotOct(s); err == nil {
		rs = append(rs, fmt.Sprintf("DotOct: %s", r))
	}
	// dot hex
	if r, err := ipToDotHex(s); err == nil {
		rs = append(rs, fmt.Sprintf("DotHex: %s", r))
	}
	return strings.Join(rs, "\r\n"), nil
}

// OctToIP -
func octToIP(s string) (string, error) {
	d, err := strconv.ParseInt(s, 8, 64)
	if err != nil {
		return "", err
	}
	return inet4ntoa(d).String(), nil
}

// DecToIP -
func decToIP(s string) (string, error) {
	d, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return "", err
	}
	return inet4ntoa(d).String(), nil
}

// HexToIP -
func hexToIP(s string) (string, error) {
	h, err := basic.HexToDec(s)
	if err != nil {
		return "", err
	}
	return decToIP(h)
}

// MACToDec -
func macToDec(s string) (string, error) {
	_, err := net.ParseMAC(s)
	if err != nil {
		return "", err
	}
	for _, c := range []string{":", "-", "."} {
		s = strings.ReplaceAll(s, c, "")
	}
	return basic.HexToDec(strings.Trim(s, ":-."))
}

// DecToMAC -
func decToMAC(s string) (string, error) {
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
