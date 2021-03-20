package bot

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"log"

	simplejson "github.com/bitly/go-simplejson"
	tb "gopkg.in/tucnak/telebot.v2"
)

func searchByText(t *tb.Message) {
	if t.Text != "" {
		keyWord := t.Payload
		dates, _ := getAPIdate(keyWord, 1)
		B.Send(t.Chat, dates)

	}

}

func getAPIdate(keyword string, pageNumber int) (date []string, err error) {
	if pageNumber <= 0 {
		pageNumber = 1
	}
	utmp := "https://api.dalipan.com/api/v1/pan/search?t=%d&kw=%s&page=%d&line=1&site=dalipan"
	url := fmt.Sprintf(utmp, time.Now().Unix(), keyword, pageNumber)
	fmt.Println(url)

	resp, err := http.Get(url)

	if err != nil {
		log.Printf("Don't get the url,error:%s\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read response date error ,error:%s\n", err)
		return nil, err
	}
	jsondate, err := simplejson.NewJson(body)
	date = []string{}
	for i := 0; i < 29; i++ {
		tmpdate, _ := jsondate.Get("resources").GetIndex(i).Get("res").Get("filename").String()
		date = append(date, tmpdate)
	}
	return date, nil
}
