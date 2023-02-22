package nettool

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

const hostRegex = `(?m)^((\d+)(-(\d+))?)\.((\d+)(-(\d+))?)\.((\d+)(-(\d+))?)\.((\d+)(-(\d+))?)(/(\d+))?`

func ipInc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// ParseHost 解析 Host 地址
// 192.168.1.1
// 192.168.1.1/24
// 192.168.1.1-8
// 192.168.1-20.8
// 192.168.1-20.1-8
// 192.168.3.1-5,192.168.1-20.1-12
func ParseHost(host string) (ips []string) {
	ips = make([]string, 0)
	var re *regexp.Regexp
	for _, h := range strings.Split(host, ",") {
		re = regexp.MustCompile(hostRegex)
		match := re.FindAllStringSubmatch(h, -1)
		// 2-4.6-8.10-12.14-16/18
		if len(match) == 0 || len(match[0]) == 0 {
			continue
		}
		m := make([]int, 9)
		for i := 1; i < 10; i++ {
			j, err := strconv.Atoi(match[0][i*2])
			if err != nil {
				m[i-1] = 0
			} else {
				m[i-1] = j
			}
		}
		// 0-1.2-3.4-5.6-7/8
		// CIDR - ip/mask 192.168.1.1/24
		if m[8] != 0 {
			if ip, ipnet, err := net.ParseCIDR(h); err == nil {
				for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); ipInc(ip) {
					ips = append(ips, ip.String())
				}
			}
		} else {
			// 1   2   3   4
			// 0-1.2-3.4-5.6-7
			ipPart := make([][]int, 4)
			for i := 1; i <= 4; i++ {
				// start = i*2-2  end = i*2-1
				if m[i*2-1] != 0 {
					partLen := m[i*2-1] - m[i*2-2] + 1 // end - start + 1
					part := make([]int, partLen)
					for j := 0; j < partLen; j++ {
						part[j] = m[i*2-2] + j
					}
					ipPart[i-1] = part
				} else {
					ipPart[i-1] = []int{m[i*2-2]}
				}
			}
			total := len(ipPart[0]) * len(ipPart[1]) * len(ipPart[2]) * len(ipPart[3])
			_ips := make([]string, 0, total)
			for _, a := range ipPart[0] {
				for _, b := range ipPart[1] {
					for _, c := range ipPart[2] {
						for _, d := range ipPart[3] {
							_ips = append(_ips, fmt.Sprintf("%d.%d.%d.%d", a, b, c, d))
						}
					}
				}
			}
			ips = append(ips, _ips...)
		}
	}
	return
}

// ParsePort 解析端口
// port,port1-port2,-port,...
// 80,8080,8000-8010,443
func ParsePort(portStr string) (ports []int) {
	psMap := make(map[int]struct{})
	waitDel := make(map[int]struct{})
	for _, s := range strings.Split(portStr, ",") {
		_s := strings.Split(s, "-")
		if len(_s) == 1 {
			// single port
			j, err := strconv.Atoi(_s[0])
			if err != nil {
				continue
			}
			if j <= 0 || j > 65535 {
				continue
			}
			psMap[j] = struct{}{}
		} else if _s[0] != "" {
			// port range
			j, err := strconv.Atoi(_s[0])
			if err != nil {
				continue
			}
			if j <= 0 || j > 65535 {
				continue
			}
			k, err := strconv.Atoi(_s[1])
			if err != nil {
				continue
			}
			if k <= 0 || k > 65535 {
				continue
			}
			for i := j; i <= k; i++ {
				psMap[i] = struct{}{}
			}
		} else {
			// -port: delete port
			k, err := strconv.Atoi(_s[1])
			if err != nil {
				continue
			}
			if k <= 0 || k > 65535 {
				continue
			}
			waitDel[k] = struct{}{}
		}
	}
	// Remove Duplication
	for k := range waitDel {
		delete(psMap, k)
	}
	ports = make([]int, 0, len(psMap))
	for k := range psMap {
		ports = append(ports, k)
	}
	return ports
}

// CheckPort -
func CheckPort(port int) error {
	if port < 1 || port > 65535 {
		return fmt.Errorf("port should be a number and the range is [1,65536)")
	}
	return nil
}

// ParseAddr -
func ParseAddr(addr string) (string, int, error) {
	ipAndPort := strings.Split(addr, ":")
	if len(ipAndPort) != 2 {
		return "", 0, fmt.Errorf("address should be a string like [ip:port]")
	}
	ip := net.ParseIP(ipAndPort[0])
	if ip == nil {
		return "", 0, fmt.Errorf("parse ip faild")
	}
	port, err := strconv.ParseInt(ipAndPort[1], 10, 64)
	if err != nil {
		return "", 0, err
	}
	return ip.String(), int(port), nil
}

// ParseURLToHostAndURI [http/s://]domain/ip[:port]/uri -> host,uri,err
func ParseURLToHostAndURI(u string) (string, string, error) {
	if !strings.HasPrefix(u, "http") {
		u = "http://" + u
	}
	us, err := url.Parse(u)
	if err != nil {
		return "", "", err
	}
	host := us.Host
	if !strings.Contains(host, ":") {
		host = host + ":80"
	}
	return host, us.RequestURI(), nil
}
