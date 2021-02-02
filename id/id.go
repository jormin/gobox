package id

import (
	"fmt"

	"gitlab.wcxst.com/jormin/goid/pkg/id"
)

var client *id.Client

// 配置
type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// 初始化
func Init(cfg *Config) error {
	var err error
	client, err = id.NewClient(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	return err
}

// 生成ID
func NewID() (int64, error) {
	return client.NewID()
}
