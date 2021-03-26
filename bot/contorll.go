package bot

import (
	"fmt"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	SetBtn = [][]tb.InlineButton{}
)

func getCallBack(t *tb.Callback) {
	fmt.Println(t)
}
