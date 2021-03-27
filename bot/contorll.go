package bot

import (
	"fmt"
	"strconv"
	"strings"

	tb "gopkg.in/tucnak/telebot.v2"
	"gotgbot.com/v0.01/util"
)

func getCallBack(t *tb.Callback) {
	callbackdate := strings.Split(t.Data, "|")
	Kw <- callbackdate[0]
	pagenumber, _ := strconv.Atoi(callbackdate[1])
	Nm <- pagenumber
	go getApiDate()
	dates := <-RespDates
	context := ""
	url := "http://www.dalipan.com"
	for _, v := range dates.Resours {
		context = context + fmt.Sprintf("[%s](%s)-[%s]\n", v.Ress.Filename, url, util.FileSize(v.Ress.Size))
	}
	pages := PageCount(dates.Total)
	B.Edit(t.Message, context, &tb.SendOptions{
		DisableWebPagePreview: true,
		ParseMode:             tb.ModeMarkdown,
		ReplyMarkup:           pagesNumber(pages, callbackdate[0]),
	})
}
func pagesNumber(n int, keyword string) *tb.ReplyMarkup {
	var (
		selector = &tb.ReplyMarkup{}
	)
	//一行最多添加八个按钮
	if n > 8 {
		var tmp []tb.Btn
		for i := 1; i < 9; i++ {
			tmp = append(tmp, selector.Data(strconv.Itoa(i), "next_page", fmt.Sprintf("%s|%d", keyword, i)))
		}
		selector.Inline(selector.Row(tmp...))
	} else {
		var tmp []tb.Btn
		for i := 1; i < n+1; i++ {
			tmp = append(tmp, selector.Data(strconv.Itoa(i), "next_page", fmt.Sprintf("%s|%d", keyword, i)))
		}
		selector.Inline(selector.Row(tmp...))
	}

	return selector
}
