package bot

import (
	"fmt"
	"strconv"

	tb "gopkg.in/tucnak/telebot.v2"
)

func getCallBack(t *tb.Callback) {
	fmt.Println(t)
}
func pagesNumber(n int) *tb.ReplyMarkup {
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
			tmp = append(tmp, selector.Data(strconv.Itoa(i), "next_page", strconv.Itoa(i)))
		}
		selector.Inline(selector.Row(tmp...))
	}

	return selector
}
