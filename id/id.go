package id

import (
	"github.com/go-resty/resty/v2"
)

// 生成ID
func NewID() string {
	c := resty.New()
	res, err := c.R().Get("https://id.wcxst.com")
	if err != nil {
		return ""
	}
	return string(res.Body())
}
