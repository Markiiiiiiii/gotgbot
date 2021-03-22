package bot

import (
	"fmt"
	"net/url"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"go.uber.org/zap"
	tb "gopkg.in/tucnak/telebot.v2"
	"gotgbot.com/v0.01/config"
	"gotgbot.com/v0.01/fetch"
)

//按照传入的参数搜索
func searchByText(t *tb.Message) {
	if t.Text != "" {
		keyWord := t.Payload
		dates, _ := getApiDate(keyWord, config.Page)
		var strs string
		for _, str := range dates {
			strs += str + "\n"
		}
		B.Send(t.Chat, strs)
	}

}

func getApiDate(keyWord string, pageNumber int) (date []string, err error) {
	zap.S().Infof("Search by Keyword: %s , Pagenumber: %d ", keyWord, pageNumber)
	utmp := "https://api.dalipan.com/api/v1/pan/search?t=%d&kw=%s&page=%d&line=1&site=dalipan"
	url := fmt.Sprintf(utmp, time.Now().Unix(), url.QueryEscape(keyWord), pageNumber)
	body, err := fetch.Fetch(url)
	if err != nil {
		zap.S().Errorf("Read response date error: %s ", err)
		return nil, err
	}
	date = []string{}
	jsondate, err := simplejson.NewJson(body)
	if err != nil {
		zap.S().Errorf("SimpleJSON Umshall error: %s", err)
		date = append(date, "关键词:< "+keyWord+" >可能被屏蔽了")
		return date, err
	}

	for i := 0; i < 29; i++ {
		tmpdate, _ := jsondate.Get("resources").GetIndex(i).Get("res").Get("filename").String()
		date = append(date, tmpdate)
	}
	return date, nil
}
