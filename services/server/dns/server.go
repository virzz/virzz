package dns

import (
	"fmt"
	"net"
	"runtime"
	"time"

	"github.com/miekg/dns"
	"github.com/virzz/virzz/common"
	"github.com/virzz/virzz/logger"
)

var conf common.DNSConfig

func do(w dns.ResponseWriter, req *dns.Msg) {
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

	ttl, ip, resp := TTT(q.Name, remoteAddr.IP.String())
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

// NewDNSServer New DNS Server
func NewDNSServer() *dns.Server {
	conf = common.GetConfig().DNS
	handler := dns.NewServeMux()
	if runtime.GOOS == "windows" {
		conf.Port += 10000
	}
	timeout := time.Duration(conf.Timeout) * time.Second
	handler.HandleFunc(".", do)

	server := &dns.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", conf.Port),
		Net:          "udp",
		Handler:      handler,
		UDPSize:      65535,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}
	return server
}
