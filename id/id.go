package id

import (
	"crypto/tls"

	"github.com/go-resty/resty/v2"
)

// NewID 生成ID
func NewID() string {
	c := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	res, err := c.R().Get("https://id.wcxst.com")
	if err != nil {
		return ""
	}
	return string(res.Body())
}
