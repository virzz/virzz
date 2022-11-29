package gopher

import "strings"

func replaceFastCGIPayload(p string) string {
	p = strings.ReplaceAll(p, "+", "%20")
	p = strings.ReplaceAll(p, "%2F", "/")
	return p
}

func replaceRedisPayload(p string) string {
	p = replaceFastCGIPayload(p)
	p = strings.ReplaceAll(p, "%25", "%")
	p = strings.ReplaceAll(p, "%3A", ":")
	return p
}
