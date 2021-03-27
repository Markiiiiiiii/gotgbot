package bot

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"go.uber.org/zap"
	tb "gopkg.in/tucnak/telebot.v2"
	"gotgbot.com/v0.01/config"
	"gotgbot.com/v0.01/fetch"
	"gotgbot.com/v0.01/model"
	"gotgbot.com/v0.01/util"
)

//按照传入的参数搜索
func searchByText(t *tb.Message) {
	if t.Text != "" {
		Kw <- t.Payload
		Nm <- config.Page
		go getApiDate()
		// dates, err := getApiDate(t.Payload, config.Page)
		// if err != nil {
		// 	zap.S().Errorf("Get dates error: %s ", err)
		// 	return
		// }
		// fmt.Printf("%#v", dates.Resours)
		dates := <-RespDates
		if dates.Resours == nil {
			B.Reply(t, "您搜索的关键词被屏蔽或相关的资源不存在")
		} else {
			context := ""
			url := "http://www.dalipan.com"
			for _, v := range dates.Resours {
				context = context + fmt.Sprintf("[%s](%s)-[%s]\n", v.Ress.Filename, url, util.FileSize(v.Ress.Size))
			}
			pages := PageCount(dates.Total)

			// B.Send(t.Chat, context, &tb.SendOptions{
			// 	DisableWebPagePreview: true,
			// 	ParseMode:             tb.ModeMarkdown,
			// 	ReplyMarkup:           pagesNumber(pages, t.Payload),
			// })
			B.Reply(t, context, &tb.SendOptions{
				DisableWebPagePreview: true,
				ParseMode:             tb.ModeMarkdown,
				ReplyMarkup:           pagesNumber(pages, t.Payload),
			})
		}
	}
}

func getApiDate() {
	var dates model.NetDiskDate
	kw := <-Kw
	num := <-Nm
	fmt.Println(kw)
	zap.S().Infof("Search by Keyword: %s , Pagenumber: %d ", kw, num)
	url := fmt.Sprintf("https://api.dalipan.com/api/v1/pan/search?t=%d&kw=%s&page=%d&line=1&site=dalipan", time.Now().Unix(), url.QueryEscape(kw), num)
	//请求数据
	body, err := fetch.Fetch(url)
	if err != nil {
		zap.S().Errorf("Read response date error: %s ", err)
	}
	if string(body[:]) != "privacy" {
		//序列化json数据
		err := json.Unmarshal(body, &dates)
		if err != nil {
			zap.S().Errorf("Unmarshal date error: %s ", err)
		}
	}
	RespDates <- dates
}

//计算分页数
func PageCount(number int) (pages int) {
	if number/30 < 2 && number%30 != 0 {
		pages = 2
	} else {
		pages = number / 30
	}
	return
}
