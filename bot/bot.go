package bot

import (
	"log"
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

func init() {
	poller := &tb.LongPoller{Timeout: time.Second * 5}
	apiurl := tb.DefaultApiURL
	proxyurl, err := url.Parse(config.HttpProxy)
	if err != nil {
		log.Printf("%s", err)
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyurl),
		},
	}

	zap.S().Infow("Init bot token",
		"token", config.BotToken,
		"endpoint", apiurl)
	B, err = tb.NewBot(tb.Settings{
		URL:    apiurl,
		Token:  config.BotToken,
		Poller: poller,
		Client: client,
	})
	if err != nil {
		log.Printf("%s", err)
		return
	}
}

func Start() {
	B.Handle("/test", func(t *tb.Message) {
		B.Send(t.Sender, "Hello World!")
	})
	B.Handle("/search", searchByText)
	B.Start()
}
