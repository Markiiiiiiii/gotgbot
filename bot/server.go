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
)

//按照传入的参数搜索
func searchByText(t *tb.Message) {
	//创建行内按钮，注册一个唯一的名称
	setBtn := [][]tb.InlineButton{}
	setBtn = append(setBtn, []tb.InlineButton{
		tb.InlineButton{
			Unique: "next_btn",
			Text:   "下一页",
			Data:   "getCallBack:next",
		},
	})
	setBtn = append(setBtn, []tb.InlineButton{
		tb.InlineButton{
			Unique: "pre_btn",
			Text:   "上一页",
			Data:   "getCallBack:prevg",
		},
	})
	if t.Text != "" {
		dates, err := getApiDate(t.Payload, config.Page)
		if err != nil {
			zap.S().Errorf("Get dates error: %s ", err)
			return
		}
		// fmt.Printf("%#v", dates.Resours)
		if dates.Resours == nil {
			B.Send(t.Chat, "关键词被屏蔽或资源不存在！")
		} else {
			context := ""
			url := "http://www.dalipan.com"
			for _, v := range dates.Resours {
				context = context + fmt.Sprintf("[%s](%s)-[%d]\n", v.Ress.Filename, url, v.Ress.Size)
			}
			B.Send(t.Chat, context, &tb.SendOptions{
				DisableWebPagePreview: true,
				ParseMode:             tb.ModeMarkdown,
			}, &tb.ReplyMarkup{
				InlineKeyboard: setBtn,
			})

		}
	}
}

//调用API使用model.NetDiskDate结构体反序列化
func getApiDate(kw string, num int) (dates model.NetDiskDate, err error) {
	zap.S().Infof("Search by Keyword: %s , Pagenumber: %d ", kw, num)
	url := fmt.Sprintf("https://api.dalipan.com/api/v1/pan/search?t=%d&kw=%s&page=%d&line=1&site=dalipan", time.Now().Unix(), url.QueryEscape(kw), num)
	//请求数据
	body, err := fetch.Fetch(url)
	if err != nil {
		zap.S().Errorf("Read response date error: %s ", err)
		return dates, err
	}
	if string(body[:]) == "privacy" {
		return dates, err
	} else {
		//序列化json数据
		err := json.Unmarshal(body, &dates)
		if err != nil {
			zap.S().Errorf("Unmarshal date error: %s ", err)
			return dates, nil
		}
	}
	return dates, nil
}

func getCallBack(t *tb.Callback) {
	fmt.Println(t)
}
