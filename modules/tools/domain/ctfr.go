package domain

import (
	"fmt"
	"time"

	"github.com/mozhu1024/virzz/common"
	"github.com/mozhu1024/virzz/modules/httpreq"
)

type CrtResp struct {
	IssuerCaID     int    `json:"issuer_ca_id"`
	IssuerName     string `json:"issuer_name"`
	CommonName     string `json:"common_name"`
	NameValue      string `json:"name_value"`
	ID             int64  `json:"id"`
	EntryTimestamp string `json:"entry_timestamp"`
	NotBefore      string `json:"not_before"`
	NotAfter       string `json:"not_after"`
	SerialNumber   string `json:"serial_number"`
}

func ctfr(domain string) (string, error) {
	var result []CrtResp
	_, err := httpreq.New().SetTimeout(60*time.Second).
		R().
		SetResult(&result).
		SetHeader("Accept", "application/json").
		SetQueryParams(map[string]string{
			"q":      domain,
			"output": "json",
		}).Get("https://crt.sh/")
	if err != nil {
		return "", err
	}
	domainMap := make(map[string]bool, 0)
	for _, ret := range result {
		domainMap[ret.NameValue] = true
	}
	data := make([]map[int]string, 0, len(domainMap))
	for key := range domainMap {
		data = append(data, map[int]string{2: key})
	}
	return common.TableOutput(
		data,
		[]string{fmt.Sprintf("Domains - total: %d", len(domainMap))}), nil
}
