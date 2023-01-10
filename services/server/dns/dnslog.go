package dns

import (
	"fmt"
	"math/rand"
	"net"
	"regexp"
	"strings"

	"github.com/miekg/dns"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/services/server/models"
)

const (
	DOMAIN_PREFIX_CUSTOM    = "c"
	DOMAIN_PREFIX_REBINDING = "r"
	DOMAIN_PREFIX_TXT       = "t"
)

/*
parseQName

	[data/r].token.domain
	normal
		data xxx.aaa.bbb.ccc.token.[xxx.com]
	special
		rebinding token.r.[xxx.com] r.token.[xxx.com]
		custom ssss.s.[xxx.com]
		txt token.t.[xxx.com]
		ip 127.0.0.1.ip.[xxx.com]
		ip 2130706433.ip.[xxx.com]
		single a.[xxx.com]
*/
func parseQName(qName string) (data, t string) {
	reg := regexp.MustCompile(`(?m)([\w\.]+)\.([\w]+)`)
	res := reg.FindAllStringSubmatch(qName, -1)
	logger.DebugF("qName = '%s' , parseQName = '%s'", qName, res)
	if len(res) > 0 && len(res[0]) > 2 {
		data = res[0][1]
		t = res[0][2]
	}
	return
}

func TTT(name, remoteIP string) (ttl int, ip net.IP, resp string) {
	// unFqdn
	if dns.IsFqdn(name) {
		name = name[:len(name)-1]
	}
	// trim Domain
	name = strings.Trim(name, fmt.Sprintf(".%s", conf.Domain))

	data, t := parseQName(name)
	logger.DebugF("data = '%s' , t = '%s'", data, t)
	if data == "" && t == "" {
		// return error by ttl == -1
		ttl = -1
		return
	}

	ttl = conf.TTL
	ip = net.ParseIP(conf.Host)

	switch t {
	case DOMAIN_PREFIX_REBINDING:
		// [data=token].r.domain
		// ip1-ip2.r.domain
		ttl = 0
		if strings.Contains(data, "-") {
			ip = net.ParseIP(strings.Split(data, "-")[rand.Intn(2)])
		} else {
			// 重绑定只允许 2 个 IP，且随机切换
			m, err := models.FindRecordByToken(data, t)
			if err == nil {
				ip = net.ParseIP(strings.Split(m.Record, ",")[rand.Intn(2)])
			} else {
				resp = "error=not set rebinding ip"
				ip = net.ParseIP("127.0.0.1")
			}
		}
	case DOMAIN_PREFIX_CUSTOM:
		// [randstr].c.domain
		m, err := models.FindRecordByToken(data, t)
		if err == nil {
			ip = net.ParseIP(m.Record)
		}
	case DOMAIN_PREFIX_TXT:
		// [token].t.domain
		m, err := models.FindRecordByToken(data, t)
		if err == nil {
			resp = m.Record
		}
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
