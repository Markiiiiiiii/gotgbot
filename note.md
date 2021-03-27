```
	//创建行内按钮，注册一个唯一的名称

	// SetBtn = append(SetBtn, []tb.InlineButton{
	// 	tb.InlineButton{
	// 		Unique: "next_btn",
	// 		Text:   "下一页",
	// 		Data:   "getCallBack:next",
	// 	},
	// })
	// SetBtn = append(SetBtn, []tb.InlineButton{
	// 	tb.InlineButton{
	// 		Unique: "pre_btn",
	// 		Text:   "上一页",
	// 		Data:   "getCallBack:prevg",
	// 	},
	// })
```
```
/发送行内按钮
			// B.Send(t.Chat, context, &tb.SendOptions{
			// 	DisableWebPagePreview: true,
			// 	ParseMode:             tb.ModeMarkdown,
			// }, &tb.ReplyMarkup{
			// 	InlineKeyboard: pagesNumber(pages),
			// })
```
单线程
```
//调用API使用model.NetDiskDate结构体反序列化
// func getApiDate(kw string, num int) (dates model.NetDiskDate, err error) {
// 	zap.S().Infof("Search by Keyword: %s , Pagenumber: %d ", kw, num)
// 	url := fmt.Sprintf("https://api.dalipan.com/api/v1/pan/search?t=%d&kw=%s&page=%d&line=1&site=dalipan", time.Now().Unix(), url.QueryEscape(kw), num)
// 	//请求数据
// 	body, err := fetch.Fetch(url)
// 	if err != nil {
// 		zap.S().Errorf("Read response date error: %s ", err)
// 		return dates, err
// 	}
// 	if string(body[:]) == "privacy" {
// 		return dates, err
// 	} else {
// 		//序列化json数据
// 		err := json.Unmarshal(body, &dates)
// 		if err != nil {
// 			zap.S().Errorf("Unmarshal date error: %s ", err)
// 			return dates, nil
// 		}
// 	}
// 	return dates, nil
// }
```
```
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

```