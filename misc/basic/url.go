package basic

import "net/url"

// URLEncode -
func URLEncode(s string) (string, error) {
	return url.QueryEscape(s), nil
}

// URLDecode -
func URLDecode(s string) (string, error) {
	return url.QueryUnescape(s)
}

// // URLEncodePlus -
// func URLEncodePlus(s string) (string, error) {
// 	return url.QueryEscape(s), nil
// }

// // URLDecodePlus -
// func URLDecodePlus(s string) (string, error) {
// 	return url.QueryUnescape(s)
// }
