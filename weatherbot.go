package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	/*
		国内主机需要使用http代理
		 proxyUrl, err := url.Parse("http://127.0.0.1:8001")                                  //设置代理http或sock5
		myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}} //使用代理
		bot, err := tgbotapi.NewBotAPIWithClient("1479970937:AAFKPdrFUmUCsnF4waRr4dZmLieY9Qj76-U", myClient)
	*/

	// 创建一个新的BOT连接
	bot, err := tgbotapi.NewBotAPI("1479970937:AAFKPdrFUmUCsnF4waRr4dZmLieY9Qj76-U")
	if err != nil {
		log.Panic(err)
	}
	// 开启debug
	bot.Debug = true
	//日志输出
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			// ignore any non-Message Updates
			continue
		}
		weather := getinfo(update.Message.Text)
		fmt.Println(weather)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, weather)
		msg.ReplyToMessageID = update.Message.MessageID
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

// 获取城市天气
func getinfo(cityname string) string {
	URL := "https://way.jd.com/he/freeweather?appkey=79d6586692c7cb8127acf38815f8f6c8&city=" + cityname
	queryUrl := fmt.Sprintf("%s", URL)
	resp, err := http.Get(queryUrl)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}
	res, err := simplejson.NewJson(body)
	qlty, _ := res.Get("result").Get("HeWeather5").GetIndex(0).Get("aqi").Get("city").Get("qlty").String()
	return qlty
}
