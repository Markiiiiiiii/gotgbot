package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		zap.S().Fatal(err)
		return
	}
	BotToken = viper.GetString("bot_token")
	HttpProxy = viper.GetString("http_proxy")
}
