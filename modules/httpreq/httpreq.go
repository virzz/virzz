package httpreq

import (
	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
)

var std = resty.New()

func init() {
	initJson(std)
}

func initJson(c *resty.Client) {
	c.JSONMarshal = json.Marshal
	c.JSONUnmarshal = json.Unmarshal
}

func R() *resty.Request {
	return std.R()
}

func New() *resty.Client {
	c := resty.New()
	initJson(c)
	return c
}
