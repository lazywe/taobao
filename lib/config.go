package lib

import (
	"log"
)

type config struct {
	Debug     bool
	AppKey    string
	SecretKey string
}

var c *config

// 设置配置 SetConfig
func SetConfig(appKey string, secretKey string, debug bool) {
	c = &config{
		Debug:     debug,
		AppKey:    appKey,
		SecretKey: secretKey,
	}
}

// 获取配置 GetConfig
func GetConfig() *config {
	if c.AppKey == "" || c.SecretKey == "" {
		log.Fatalf("未配置【AppKey】【SecretKey】")
	}
	return c
}
