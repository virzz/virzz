package utils

import (
	"fmt"
	"testing"
)

func TestParseHost(t *testing.T) {
	// ip := net.ParseIP("192.168.0.1")
	// fmt.Println(len(ip), []byte(ip))
	// ip = net.ParseIP("192.168.0.2")
	// fmt.Println(len(ip), []byte(ip))
	ParseHost("192.168.1.100-999/30,192.168.0.1/30")
	// ParseHost("192.168.1.1/24")
	// ParseHost("192.168.1.1-100")
	return
}

func TestParsePort(t *testing.T) {
	ps := ParsePort("80,8080,8000-8010,443,8443,81-81,101,02")
	fmt.Println(ps)
	return
}
