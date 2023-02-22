package nettool

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/virzz/virzz/modules/crypto/basic"
)

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

// IP2Dec 127.0.0.1 -> 2130706433
func IP2Dec(s string) (string, error) {
	ip := net.ParseIP(s)
	return strconv.FormatInt(inet4aton(ip), 10), nil
}

// Dec2IP 2130706433 -> 127.0.0.1
func Dec2IP(s string) (string, error) {
	d, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return "", err
	}
	return inet4ntoa(d).String(), nil
}

// IP2Oct 127.0.0.1 -> 2130706433
func IP2Oct(s string) (string, error) {
	ip := net.ParseIP(s)
	return "0" + strconv.FormatInt(inet4aton(ip), 8), nil
}

// Oct2IP 017700000001 -> 127.0.0.1
func Oct2IP(s string) (string, error) {
	d, err := strconv.ParseInt(s, 8, 64)
	if err != nil {
		return "", err
	}
	return inet4ntoa(d).String(), nil
}

// IP2Hex 127.0.0.1 -> 0x7f000001
func IP2Hex(s string) (string, error) {
	s, _ = IP2Dec(s)
	return basic.DecToHex(s)
}

// Hex2IP 0x7f000001 -> 127.0.0.1
func Hex2IP(s string) (string, error) {
	h, err := basic.HexToDecStr(s)
	if err != nil {
		return "", err
	}
	return Dec2IP(h)
}

// IP2DotOct 127.0.0.1 -> 0177.0.0.01
func IP2DotOct(s string) (string, error) {
	rs := make([]string, 4)
	for i, p := range strings.Split(s, ".") {
		_i, _ := strconv.Atoi(p)
		rs[i] = fmt.Sprintf("%#[1]o", _i)
	}
	return strings.Join(rs, "."), nil
}

func DotOct2IP(s string) (string, error) {
	rs := make([]string, 4)
	for i, p := range strings.Split(s, ".") {
		_i, err := strconv.ParseInt(p, 8, 64)
		if err != nil {
			return "", err
		}
		rs[i] = fmt.Sprintf("%d", _i)
	}
	return strings.Join(rs, "."), nil
}

func DotHex2IP(s string) (string, error) {
	rs := make([]string, 4)
	for i, p := range strings.Split(s, ".") {
		_i, err := basic.HexToDec(p)
		if err != nil {
			return "", err
		}
		rs[i] = fmt.Sprintf("%d", _i.Int64())
	}
	return strings.Join(rs, "."), nil
}

// IP2DotHex 127.0.0.1 -> 0x7f.0.0.0x1
func IP2DotHex(s string) (string, error) {
	rs := make([]string, 4)
	for i, p := range strings.Split(s, ".") {
		_i, _ := strconv.Atoi(p)
		if _i == 0 {
			rs[i] = "0"
		} else {
			rs[i] = fmt.Sprintf("%#x", _i)
		}
	}
	return strings.Join(rs, "."), nil
}

// IP2All -
func IP2All(s string) (string, error) {
	rs := make([]string, 0)
	// oct
	if r, err := IP2Oct(s); err == nil {
		rs = append(rs, fmt.Sprintf("Oct   :    %s", r))
	}
	// dec
	if r, err := IP2Dec(s); err == nil {
		rs = append(rs, fmt.Sprintf("Dec   :    %s", r))
	}
	// hex
	if r, err := IP2Hex(s); err == nil {
		rs = append(rs, fmt.Sprintf("Hex   :    %s", r))
	}
	// dot oct
	if r, err := IP2DotOct(s); err == nil {
		rs = append(rs, fmt.Sprintf("DotOct: %s", r))
	}
	// dot hex
	if r, err := IP2DotHex(s); err == nil {
		rs = append(rs, fmt.Sprintf("DotHex: %s", r))
	}
	return strings.Join(rs, "\r\n"), nil
}

func AnyToIP(s string) (string, error) {
	if strings.HasPrefix(s, "0x") {
		if strings.Contains(s, ".") {
			return DotHex2IP(s)
		}
		return Hex2IP(s)
	}
	if strings.HasPrefix(s, "0") {
		if strings.Contains(s, ".") {
			return DotOct2IP(s)
		}
		return Oct2IP(s)
	}
	if strings.Contains(s, ".") {
		return s, nil
	}
	return Dec2IP(s)
}
