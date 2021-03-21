package bot

import (
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"
	tb "gopkg.in/tucnak/telebot.v2"
	"gotgbot.com/v0.01/config"
)

var (
	B *tb.Bot
)

//初始化
func init() {
	//轮询间隔5秒
	poller := &tb.LongPoller{Timeout: time.Second * 5}
	//使用默认的bot注册api地址注册本bot
	apiurl := tb.DefaultApiURL
	// 使用http proxy创建连接
	proxyurl, err := url.Parse(config.HttpProxy)
	if err != nil {
		zap.S().Errorw("Read http proxy faild", err)
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyurl),
		},
	}
	//输出INFO到log
	zap.S().Infow("Init bot token",
		"token", config.BotToken,
		"endpoint", apiurl)
	//创建一个bot实例
	B, err = tb.NewBot(tb.Settings{
		URL:    apiurl,
		Token:  config.BotToken,
		Poller: poller,
		Client: client,
	})
	if err != nil {
		zap.S().Errorw("Creat bot faild", err)
		return
	}
}

//Start BOT启动函数
func Start() {

	B.Handle("/test", func(t *tb.Message) {
		B.Send(t.Sender, "Hello World!")
	})
	//注册/search命令
	B.Handle("/search", searchByText)
	B.Start()
}
