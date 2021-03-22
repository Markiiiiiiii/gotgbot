package fetch

import (
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

func Fetch(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		zap.S().Errorf("Create a new request error: %s", err)
		return nil, err
	}
	//构建访问头
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36")
	req.Header.Add("Host", "api.dalipan.com")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Cache-Control", "max-age=0")

	//访问
	reps, err := client.Do(req)
	if err != nil {
		zap.S().Errorf("Don't get a response error: %s", err)
		return nil, err
	}
	defer reps.Body.Close()

	if reps.StatusCode != http.StatusOK {
		zap.S().Errorf("Don't get date from the URL: %d", reps.StatusCode)
		return nil, err
	} else {
		zap.S().Infof("Fetch response code: %d", reps.StatusCode)
		return ioutil.ReadAll(reps.Body)
	}

}
