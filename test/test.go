package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"gotgbot.com/v0.01/model"
)

func main() {
	client := &http.Client{}
	url := "https://api.dalipan.com/api/v1/pan/search?t=" + strconv.FormatInt(time.Now().Unix(), 10) + "&kw=" + url.QueryEscape("windows") + "&page=1&line=1&site=dalipan"
	req, _ := http.NewRequest("GET", url, nil)
	//构建访问头
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36")
	req.Header.Add("Host", "api.dalipan.com")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Cache-Control", "max-age=0")

	//访问api
	reps, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer reps.Body.Close()

	tmp, _ := ioutil.ReadAll(reps.Body)

	var dates model.NetDiskDate
	json.Unmarshal(tmp, &dates)
	fmt.Println(dates)
}
