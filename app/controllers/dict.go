/*
@Time : 2019/9/26 14:28
@Author : zxr
@File : dict
@Software: GoLand
*/
package controllers

import (
	"net/http"
	url2 "net/url"
	"poetry/tools"
)

//词典反查
func FanCha(w http.ResponseWriter, r *http.Request) {
	var (
		bytes   []byte
		err     error
		htmlStr string
	)
	zStr := r.FormValue("z")
	urlVal := url2.Values{}
	if len(zStr) == 0 {
		goto OutPutEmptyStr
	}
	urlVal.Add("z", zStr)
	urlVal.Add("url", "www.gushiwen.org/")
	if bytes, err = tools.HttpPost("https://www.gushiwen.org/dict/fancha.aspx", urlVal.Encode()); err != nil {
		goto OutPutEmptyStr
	}
	htmlStr = tools.ReplaceDictHtml(string(bytes))
	tools.OutputString(w, htmlStr)
	return
OutPutEmptyStr:
	tools.OutputString(w, "")
	return
}
