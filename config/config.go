package config

import (
	"time"

	"github.com/spf13/viper"
)

var (
	//请求地址
	BotToken   string
	HttpProxy  string
	Url        string = "https://api.dalipan.com/api/v1/pan/search?t=%s&kw=%s&page=%s&line=0&site=dalipan"
	Page       int    = 1
	SearchTime int64  = time.Now().Unix()
)

func GetString(key string) string {
	var value string
	if viper.IsSet(key) {
		value = viper.GetString(key)
	}
	return value
}
