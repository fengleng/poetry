/*
@Time : 2019/9/26 14:37
@Author : zxr
@File : curl
@Software: GoLand
*/
package tools

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//http post request
func HttpPost(url string, postData string) (bytes []byte, err error) {
	var (
		request  *http.Request
		response *http.Response
	)
	client := &http.Client{
		Timeout: time.Duration(time.Second * 5),
	}
	if request, err = http.NewRequest(http.MethodPost, url, strings.NewReader(postData)); err != nil {
		return
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("referer", "https://www.gushiwen.org/")
	request.Header.Add("Host", "www.gushiwen.org")
	request.Header.Add("Origin", "https://www.gushiwen.org")

	if response, err = client.Do(request); err != nil {
		return
	}
	defer response.Body.Close()
	bytes, err = ioutil.ReadAll(response.Body)
	return
}
