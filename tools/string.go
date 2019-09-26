/*
@Time : 2019/9/24 18:47
@Author : zxr
@File : string
@Software: GoLand
*/
package tools

import (
	"net/http"
	"poetry/config/define"
	"regexp"

	"strings"
)

//为字符串添加HTML标签
func AddHtmlLabel(content string) string {
	if len(content) == 0 {
		return ""
	}
	content = strings.TrimRight(content, "。")
	content = strings.Replace(content, "。", "。</p><p>", -1)
	return "<p>" + content
}

//输出string到ResponseWriter
func OutputString(w http.ResponseWriter, str string) {
	w.Write([]byte(str))
	return
}

//词典内容，替换路径
func ReplaceDictHtml(str string) string {
	str = strings.Replace(str, "https://song.gushiwen.org/dict", define.CdnStaticDomain+"/static", -1)
	str = strings.Replace(str, "imgs", "images", -1)
	compile := regexp.MustCompile("<img id=\"imgMp3\".*/>")
	str = compile.ReplaceAllString(str, "")
	return str
}
