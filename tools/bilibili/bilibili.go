package bilibili

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/virink/virzz/common"
	"github.com/virink/virzz/tools/downloader"
)

// PlayInfo -
type PlayInfo struct {
	Data struct {
		Dash struct {
			Duration int64 `json:"duration"`
			Audio    []struct {
				Bandwidth int64  `json:"bandwidth"`
				BaseURL   string `json:"base_url"`
				MimeType  string `json:"mime_type"`
			} `json:"audio"`
			Video []struct {
				Bandwidth int64  `json:"bandwidth"`
				BaseURL   string `json:"base_url"`
				Height    int64  `json:"height"`
				Width     int64  `json:"width"`
				MimeType  string `json:"mime_type"`
			} `json:"video"`
		} `json:"dash"`
	} `json:"data"`
}

const biliUserAgent = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/538.36 (KHTML, like Gecko) Chrome/85.0.4247.125 Safari/547.36 Edg/85.0.521.69`

// ParseVideoInfo -
func ParseVideoInfo(u string) (pi PlayInfo, err error) {
	c := colly.NewCollector(
		colly.UserAgent(biliUserAgent),
	)
	c.OnHTML("script", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "__playinfo__") {
			script := strings.ReplaceAll(strings.TrimSpace(e.Text), "window.__playinfo__=", "")
			if err = json.Unmarshal([]byte(script), &pi); err != nil {
				return
			}
		}
	})
	c.OnHTML("title", func(e *colly.HTMLElement) {
		common.Logger.Normal("Title: %s", strings.ReplaceAll(strings.TrimSpace(e.Text), "_哔哩哔哩 (゜-゜)つロ 干杯~-bilibili", ""))
	})
	c.Visit(u)
	return pi, err
}

// Bilibilis -
func Bilibilis(s ...string) error {
	for _, i := range s {
		Bilibili(i)
	}
	return nil
}

// Bilibili -
func Bilibili(s string) error {
	// BV / av
	var re = regexp.MustCompile(`(?i)(?m)\b((av\d+)|(bv1\w+))`)
	m := re.FindAllStringSubmatch(s, -1)
	if len(m) == 0 || len(m[0]) == 0 {
		return fmt.Errorf("not found av/bv")
	}
	bv := m[0][0]
	common.Logger.Normal("ID(bv1/av): %s", bv)
	bvURL := fmt.Sprintf("https://www.bilibili.com/video/%s", bv)
	// Parse
	pi, err := ParseVideoInfo(bvURL)
	if err != nil {
		return err
	}
	common.Logger.Normal("Duration: %d s", pi.Data.Dash.Duration)

	// Download
	dc := downloader.NewDownloader().
		SetHeader("user-agent", biliUserAgent).
		SetHeader("referer", bvURL).
		SetTimeout(0)

	if len(pi.Data.Dash.Audio) > 0 && len(pi.Data.Dash.Video) > 0 {
		u := pi.Data.Dash.Audio[0].BaseURL
		ux, _ := url.Parse(u)
		dc.SetHeader("authority", ux.Hostname())
		dc.AddTask(u, filepath.Join(bv, "virzz.mp4a"))
		dc.AddTask(pi.Data.Dash.Video[0].BaseURL, filepath.Join(bv, "virzz.mp4"))
	}
	dc.Start()
	dc.PrintResults()
	return nil
}
