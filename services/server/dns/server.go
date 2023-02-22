package dns

import (
	"fmt"
	"net"
	"runtime"
	"time"

	"github.com/miekg/dns"
	"github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
	"github.com/virzz/logger"
)

func handle(w dns.ResponseWriter, req *dns.Msg) {
	q := req.Question[0]
	if q.Qclass != dns.ClassINET {
		dns.HandleFailed(w, req)
		return
	}
	m := new(dns.Msg)
	m.SetReply(req)

	// DNS Name
	// name := unFqdn(q.Name)
	logger.Debug("[D] q.Name = ", q.Name)
	remoteAddr := w.RemoteAddr().(*net.UDPAddr)

	ttl, ip, resp := dnslog(q.Name, remoteAddr.IP.String())
	if ttl == -1 {
		dns.HandleFailed(w, req)
		return
	}

	// Query Type
	rrHeader := dns.RR_Header{
		Name:   q.Name,
		Rrtype: q.Qtype,
		Class:  dns.ClassINET,
		Ttl:    uint32(ttl),
	}
	switch q.Qtype {
	case dns.TypeA:
		m.Answer = append(m.Answer, &dns.A{Hdr: rrHeader, A: ip})
	case dns.TypeAAAA:
		m.Answer = append(m.Answer, &dns.AAAA{Hdr: rrHeader, AAAA: ip})
	case dns.TypeTXT:
		m.Answer = append(m.Answer, &dns.TXT{Hdr: rrHeader, Txt: []string{resp}})
	default:
		dns.HandleFailed(w, req)
		return
	}
	// NOTE: 难搞啊 不同的 answer 会破坏响应包
	// if q.Qtype != dns.TypeTXT && resp != "" {
	// 	m.Answer = append(m.Answer, &dns.TXT{Hdr: rrHeader, Txt: []string{resp}})
	// }
	w.WriteMsg(m)
}

var (
	dotDomain string
	confTTL   int
	confIP    net.IP

	cachePool *cache.Cache
)

// NewServer New DNS Server
func NewServer() *dns.Server {

	cachePool = cache.New(2*time.Minute, 5*time.Minute)

	dotDomain = fmt.Sprintf(".%s", viper.GetString("dns.domain"))
	confTTL = viper.GetInt("dns.ttl")
	confIP = net.ParseIP(viper.GetString("dns.host"))

	port := viper.GetInt("dns.port")
	if runtime.GOOS == "windows" {
		port += 10000
	}
	timeout := time.Duration(viper.GetInt("dns.timeout")) * time.Second

	handler := dns.NewServeMux()
	handler.HandleFunc(".", handle)

	server := &dns.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", port),
		Net:          "udp",
		Handler:      handler,
		UDPSize:      65535,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}
	return server
}
