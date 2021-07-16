package id

import (
	"strconv"

	"github.com/go-resty/resty/v2"
)

// 生成ID
func NewID() (int64, error) {
	c := resty.New()
	res, err := c.R().Get("http://id.wcxst.com")
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(string(res.Body()), 10, 64)
}
