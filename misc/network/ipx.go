package network

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/virink/virzz/misc/basic"
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

// IPToDec -
func IPToDec(s string) (string, error) {
	ip := net.ParseIP(s)
	if ip == nil {
		return "", fmt.Errorf("parse ip faild")
	}
	return strconv.FormatInt(inet4aton(ip), 10), nil
}

// IPToHex -
func IPToHex(s string) (string, error) {
	s, err := IPToDec(s)
	if err != nil {
		return "", err
	}
	return basic.DecToHex(s)
}

// DecToIP -
func DecToIP(s string) (string, error) {
	d, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return "", err
	}
	return inet4ntoa(d).String(), nil
}

// HexToIP -
func HexToIP(s string) (string, error) {
	h, err := basic.HexToDec(s)
	if err != nil {
		return "", err
	}
	return DecToIP(h)
}
