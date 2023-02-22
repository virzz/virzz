package nettool

import (
	"fmt"
	"testing"

	"github.com/virzz/logger"
)

var hosts = map[string]int{
	"192.168.1.18":                    1,
	"192.168.1.1/24":                  256,
	"192.168.1.1-8":                   8,
	"192.168.1-12.8":                  12,
	"192.168.1-18.1-8":                18 * 8,
	"192.168.3.1-5,192.168.1-20.1-12": 5 + 20*12,
}

func init() {
	logger.SetDebug(true)
}

func TestParseHost(t *testing.T) {
	for host, hostLen := range hosts {
		ps := ParseHost(host)
		if len(ps) != hostLen {
			t.Errorf("%s: %d!= %d", host, len(ps), hostLen)
		}
	}
}

func TestParseHostMask(t *testing.T) {
	ps := ParseHost("192.168.1.1/24")
	if len(ps) != 256 || ps[0] != "192.168.1.0" {
		t.Fail()
	}
}

func TestParseHostRange(t *testing.T) {
	ps := ParseHost("1-2.3-4.5-6.7-8")
	if len(ps) != 16 || ps[0] != "1.3.5.7" {
		t.Fail()
	}
}
func BenchmarkParseHostRange(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		ParseHost("1-2.3-4.5-6.7-8")
	}
}

func TestParsePort(t *testing.T) {
	ps := ParsePort("80,8080,8000-8010,443,8443,81-81,101,02")
	fmt.Println(ps)
}

var ports = map[string]int{
	",71,72,70-80,-77,81-":                    10,
	"80,8080,8000-8010,443,8443,81-81,101,02": 18,
	"80,8080,8000-8010,443":                   14,
}

func TestPort(t *testing.T) {
	for pStr := range ports {
		p := ParsePort(pStr)
		if len(p) != ports[pStr] {
			t.Errorf("%s: %v", pStr, p)
		}
	}
}
