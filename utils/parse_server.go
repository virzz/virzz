package utils

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// Error
var (
	ErrHost = fmt.Errorf("cannot parse host")
)

func ipInc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// ParseHost 解析 Host 地址
func ParseHost(host string) (ips []string) {
	// 192.168.1.1
	// 192.168.1.1/24
	// 192.168.1.1-100
	// 192.168.1-20.100
	// 192.168.1-20.1-100
	ips = make([]string, 0)
	var re *regexp.Regexp
	for _, h := range strings.Split(host, ",") {
		// 1.2.4-6.8-10/12
		re = regexp.MustCompile(`(?m)^(\d+)\.(\d+)\.((\d+)(-(\d+))?)\.((\d+)(-(\d+))?)(/(\d+))?$`)
		match := re.FindAllStringSubmatch(h, -1)
		if len(match) > 0 && len(match[0]) > 0 {
			if match[0][6] != "" || match[0][10] != "" {
				ipJ := []int{}
				ipStart, _ := strconv.Atoi(match[0][4])
				ipEnd, _ := strconv.Atoi(match[0][6])
				if ipEnd > ipStart {
					for j := ipStart; j <= ipEnd && j < 256; j++ {
						ipJ = append(ipJ, j)
					}
				} else {
					ipJ = append(ipJ, ipStart)
				}
				// fmt.Println("ipJ", ipStart, ipEnd, ipJ)
				ipI := []int{}
				ipStart, _ = strconv.Atoi(match[0][8])
				ipEnd, _ = strconv.Atoi(match[0][10])
				if ipEnd > ipStart {
					for i := ipStart; i <= ipEnd && i < 256; i++ {
						ipI = append(ipI, i)
					}
				} else {
					ipI = append(ipI, ipStart)
				}
				// fmt.Println("ipI", ipStart, ipEnd, ipI)
				for _, j := range ipJ {
					for _, i := range ipI {
						ips = append(ips, fmt.Sprintf("%s.%s.%d.%d", match[0][1], match[0][2], j, i))
					}
				}
			} else if match[0][12] != "" {
				// CIDR - ip/mask
				if ip, ipnet, err := net.ParseCIDR(h); err == nil {
					for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); ipInc(ip) {
						ips = append(ips, ip.String())
					}
				}
			} else {
				if ip := net.ParseIP(h); ip != nil {
					ips = append(ips, ip.String())
				}
			}
		}
	}
	return
}

// ParsePort 解析端口
func ParsePort(portStr string) (ports []int) {
	ports = make([]int, 0)
	for _, s := range strings.Split(portStr, ",") {
		if p, err := strconv.Atoi(s); err != nil && p > 0 && p < 65536 {
			ports = append(ports, p)
			continue
		}
		pArr := strings.Split(s, "-")
		if len(pArr) == 2 {
			p1, _ := strconv.Atoi(pArr[0])
			p2, _ := strconv.Atoi(pArr[1])
			if p1 > 0 && p1 < 65536 && p2 > 0 && p2 < 65536 && p2 >= p1 {
				for p := p1; p <= p2; p++ {
					ports = append(ports, p)
				}
			}
		}
	}
	return
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
