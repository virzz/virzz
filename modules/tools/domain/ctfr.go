package domain

import (
	"fmt"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/virzz/virzz/utils"
	"github.com/virzz/virzz/utils/httpreq"
)

type crtResp struct {
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

// Ctfr 滥用证书透明记录 Certificate Search By https://crt.sh
func Ctfr(domain string) (string, error) {
	var result []crtResp
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
	// uniq
	domainMap := make(map[string]bool, 0)
	for _, ret := range result {
		domainMap[ret.NameValue] = true
	}
	tableCells := make([][]*simpletable.Cell, 0)
	for name := range domainMap {
		tableCells = append(tableCells, []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: name},
		})
	}
	return utils.TableOutput(
		tableCells,
		// Header
		[]*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "crt.sh | Certificate Search"},
		},
		// Footer
		[]*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("Domains - total: %d", len(domainMap))},
		},
	), nil
}
