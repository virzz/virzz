package netool

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func parseIP(ip string) (net.IP, error) {
	r := net.ParseIP(ip)
	if r == nil {
		return nil, fmt.Errorf("parse ip faild")
	}
	return r, nil
}

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
