package dns

import (
	"net"
	"regexp"
	"strings"

	"github.com/miekg/dns"
	"github.com/virzz/logger"

	"github.com/virzz/virzz/modules/tools/netool"
	"github.com/virzz/virzz/services/server/models"
)

// DOMAIN_PREFIX_*
const (
	DOMAIN_PREFIX_CUSTOM    = "c"
	DOMAIN_PREFIX_REBINDING = "r"
	DOMAIN_PREFIX_TXT       = "t"
	DOMAIN_PREFIX_IP        = "ip"
	DOMAIN_PREFIX_VIRZZ     = "v"
)

/*
parseQName

	[data/r].token.domain
	normal
		data xxx.aaa.bbb.ccc.token.[xxx.com]
	special
		rebinding token.r.[xxx.com] r.token.[xxx.com]
		custom ssss.c.[xxx.com]
		txt token.t.[xxx.com]
		ip 127.0.0.1.ip.[xxx.com]
		virzz v.[xxx.com]
*/
func parseQName(qName string) (data, t string) {
	qName = strings.Trim(qName, dotDomain)
	reg := regexp.MustCompile(`(?m)([\w\.-]+)\.([\w]+)`)
	res := reg.FindAllStringSubmatch(qName, -1)
	logger.DebugF("qName = '%s' , parseQName = '%s'", qName, res)
	if len(res) > 0 && len(res[0]) > 2 {
		data = res[0][1]
		t = res[0][2]
	}
	return
}

func dnslog(name, remoteIP string) (ttl int, ip net.IP, resp string) {
	// unFqdn
	if dns.IsFqdn(name) {
		name = name[:len(name)-1]
	}

	data, t := parseQName(name)
	logger.DebugF("data = '%s' , t = '%s'", data, t)
	if data == "" && t == "" {
		// return error by ttl == -1
		ttl = -1
		return
	}

	ttl = confTTL
	ip = confIP

	switch t {
	case DOMAIN_PREFIX_REBINDING:
		// [data=token].r.domain
		// ip1-ip2.r.domain
		ttl = 0
		var ips []string
		if strings.Contains(data, "-") {
			ips = strings.Split(data, "-")
		} else {
			m, err := models.FindRecordByToken(data, t)
			if err != nil {
				resp = "error=not set rebinding ip"
				logger.Error(err, resp)
			}
			ips = strings.Split(m.Record, ",")
		}
		c, err := cachePool.IncrementInt(remoteIP, 1)
		logger.Debug(c)
		if err != nil {
			cachePool.SetDefault(remoteIP, c)
			logger.Warn(err)
		}
		ip = net.ParseIP(ips[c%2])
	case DOMAIN_PREFIX_CUSTOM:
		// [data=randstr].c.domain
		m, err := models.FindRecordByToken(data, t)
		if err == nil {
			ip = net.ParseIP(m.Record)
		}
	case DOMAIN_PREFIX_TXT:
		// [data=token].t.domain
		m, err := models.FindRecordByToken(data, t)
		if err == nil {
			resp = m.Record
		}
	case DOMAIN_PREFIX_IP:
		// [data=token].ip.domain
		// 017700000001,2130706433,0x7f000001,127.0.0.1
		// TODO: auto convert ip
		r, err := netool.AnyToIP(data)
		if err != nil {
			logger.Error(err)
			break
		}
		ip = net.ParseIP(r)
	default:
		logger.DebugF(`New DNS Recode => [%s:%s] -> "%s"`, t, remoteIP, data)
		_, err := models.NewLog(t, data, remoteIP)
		if err != nil {
			logger.Error(err)
		}
	}
	if ip == nil {
		ip = net.ParseIP("127.0.0.1")
	}
	return
}
