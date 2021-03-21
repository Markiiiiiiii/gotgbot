package config

import (
	"github.com/spf13/viper"
)

var (
	//BOT的TOKEN请求地址
	BotToken string
	//http代理地址
	HttpProxy string
	//API地址
	Url string = "https://api.dalipan.com/api/v1/pan/search?t=%s&kw=%s&page=%s&line=0&site=dalipan"
	//默认请求页码
	Page int = 1
)

//GetString 获取配置文件的内容
func GetString(key string) string {
	var value string
	if viper.IsSet(key) {
		value = viper.GetString(key)
	}
	return value
}
