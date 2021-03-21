package bot

import (
	"fmt"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"go.uber.org/zap"
	tb "gopkg.in/tucnak/telebot.v2"
	"gotgbot.com/v0.01/config"
	"gotgbot.com/v0.01/fetch"
)

//按照传入的参数搜索
func searchByText(t *tb.Message) {
	zap.S().Infow("Handle func searchByText")
	if t.Text != "" {
		keyWord := t.Payload
		dates, _ := getAPIdate(keyWord, config.Page)
		var strs string
		for _, str := range dates {
			strs += str + "\n"
		}
		B.Send(t.Chat, strs)
		time.Sleep(time.Second * 10)
	}

}

func getAPIdate(keyWord string, pageNumber int) (date []string, err error) {
	zap.S().Infow("Keyword", keyWord,
		"Pagenumber", pageNumber)
	utmp := "https://api.dalipan.com/api/v1/pan/search?t=%d&kw=%s&page=%d&line=1&site=dalipan"
	url := fmt.Sprintf(utmp, time.Now().Unix(), keyWord, pageNumber)

	body, err := fetch.Fetch(url)
	if err != nil {
		zap.S().Errorw("Read response date error ", err)
		return nil, err
	}
	jsondate, err := simplejson.NewJson(body)
	if err != nil {
		zap.S().Errorw("SimpleJSON Umshall error", err)
		return nil, err
	}
	date = []string{}
	for i := 0; i < 29; i++ {
		tmpdate, _ := jsondate.Get("resources").GetIndex(i).Get("res").Get("filename").String()
		date = append(date, tmpdate)
	}
	return date, nil
}
