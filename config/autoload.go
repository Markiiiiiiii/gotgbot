package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

//自动初始化
func init() {
	//config文件读取设置
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		zap.S().Fatalf("Read config file faild: %s ", err)
		return
	}
	//读取token和proxy信息
	BotToken = viper.GetString("bot_token")
	HttpProxy = viper.GetString("http_proxy")
}
