package bot

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"time"

	"go.uber.org/zap"
	tb "gopkg.in/tucnak/telebot.v2"
	"gotgbot.com/v0.01/config"
	"gotgbot.com/v0.01/fetch"
)

//定义返回数据类型
// type NetDiskDate struct {
// 	Title string
// 	Size  int64
// }
// type ResultDate struct{
// 	NetDiskDates []NetDiskDate
// }

//按照传入的参数搜索
func searchByText(t *tb.Message) {
	if t.Text != "" {
		dates, _ := getApiDates(t.Payload,config.Page)
		//判断返回值是否有被屏蔽的值
		if _, ok := dates.(string); ok {
			B.Send(t.Chat, dates)
		}
		//判断返回值是否是一个NetDiskDate类型的切片
		if _, ok := dates.(NetDiskDate); ok {
			dateList := reflect.TypeOf(dates)
			strs := ""
			url := "http://www.dalipan.com"
			for j := 0; j < dateList.Len(); j++ {
					fmt.Println(dateList.FieldByName("Title"))

					// field := dateList.Field(i)
					// strs = strs + fmt.Sprintf("[%s](%s)-[%d]\n", field.Value, url)
				

			}

		// 	B.Send(t.Chat, strs, &tb.SendOptions{
		// 		DisableWebPagePreview: true,
		// 		ParseMode:             tb.ModeMarkdown,
		// 	})
		// }
	}

}

// func sendToTelgram (t *tb.Message,date interface{}){

// }
//调用API
func getApiDate(keyWord string,pageNumber int)(ResultDate,error){
	var dates []NetDiskDate
	zap.S().Infof("Search by Keyword: %s , Pagenumber: %d ", keyWord, pageNumber)
	url := fmt.Sprintf("https://api.dalipan.com/api/v1/pan/search?t=%d&kw=%s&page=%d&line=1&site=dalipan", time.Now().Unix(), url.QueryEscape(keyWord), pageNumber)
	//请求数据
	body, err := fetch.Fetch(url)
	if err != nil {
		zap.S().Errorf("Read response date error: %s ", err)
		return nil,err
	}
	// 如果请求返回的数据一个私有类型则存在关键词被屏蔽返回信息
	if string(body[:]) == "privacy" {
		return  nil,err
	} else {
		//序列化json数据
		jsondate, err := simplejson.NewJson(body)
		if err != nil {
			zap.S().Errorf("SimpleJSON Unmarshal error: %s", err)
			return nil,err
		}

		for i := 0; i < 29; i++ {
			&dates[i].Title, _ = jsondate.Get("resources").GetIndex(i).Get("res").Get("filename").String()
			&dates[i].Size, _ = jsondate.Get("resources").GetIndex(i).Get("res").Get("size").Int64()
		}
		return dates,nil
	}
}

func getApiDates(k string,num int)(baiduNetDiskDate,error){
	var dates baiduNetDiskDate
	zap.S().Infof("Search by Keyword: %s , Pagenumber: %d ", k, num)
	url := fmt.Sprintf("https://api.dalipan.com/api/v1/pan/search?t=%d&kw=%s&page=%d&line=1&site=dalipan", time.Now().Unix(), url.QueryEscape(k), num)
	//请求数据
	body, err := fetch.Fetch(url)
	if err != nil {
		zap.S().Errorf("Read response date error: %s ", err)
		return nil,err
	}
	if string(body[:]) == "privacy" {
		return  nil,err
	} else {
				//序列化json数据
		err :=json.Unmarshal(body,&dates)
		if err!=nil {
			zap.S().Errorf("Unmarshal date error: %s ", err)
			return nil ,nil 
		}
	}
return dates,nil
}