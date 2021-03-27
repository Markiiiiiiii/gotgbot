package bot

import (
	"fmt"
	"strconv"
	"strings"

	"go.uber.org/zap"
	tb "gopkg.in/tucnak/telebot.v2"
	"gotgbot.com/v0.01/util"
)

func getCallBack(t *tb.Callback) {
	callbackdate := strings.Split(t.Data, "|")
	keyword := callbackdate[0]
	pagenumber, _ := strconv.Atoi(callbackdate[1])
	dates, err := getApiDate(keyword, pagenumber)
	if err != nil {
		zap.S().Errorf("Get dates error: %s ", err)
		return
	}
	context := ""
	url := "http://www.dalipan.com"
	for _, v := range dates.Resours {
		context = context + fmt.Sprintf("[%s](%s)-[%s]\n", v.Ress.Filename, url, util.FileSize(v.Ress.Size))
	}
	pages := PageCount(dates.Total)
	B.Edit(t.Message, context, &tb.SendOptions{
		DisableWebPagePreview: true,
		ParseMode:             tb.ModeMarkdown,
		ReplyMarkup:           pagesNumber(pages, keyword),
	})
}
func pagesNumber(n int, keyword string) *tb.ReplyMarkup {
	var (
		// SetBtn   = [][]tb.InlineButton{}
		selector = &tb.ReplyMarkup{}
	)
	// if n > 5 {
	// 	for i := 0; i < 5; i++ {
	// 		SetBtn = append(SetBtn, []tb.InlineButton{
	// 			tb.InlineButton{
	// 				Unique: "next_page",
	// 				Text:   strconv.Itoa(i + 1),
	// 				Data:   "page_num:" + strconv.Itoa(i)},
	// 		})
	// 	}
	// } else {
	// 	for i := 0; i < n; i++ {
	// 		SetBtn = append(SetBtn, []tb.InlineButton{
	// 			tb.InlineButton{
	// 				Unique: "next_page",
	// 				Text:   strconv.Itoa(i + 1),
	// 				Data:   "page_num:" + strconv.Itoa(i)},
	// 		})
	// 	}
	// }
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
