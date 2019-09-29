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

//去掉字符串右边某些字符
func TrimRightHtml(str string) string {
	if len(str) == 0 {
		return ""
	}
	str = strings.TrimRight(str, "”</p>")
	return str
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

//诗词详情页赏析和翻译数据和诗文详情替换所有的超链接，和字符中的ID
func DealWithNotes(content, idStr string) string {
	compile := regexp.MustCompile(`javascript:shangxiClose\(\d+\)`)
	content = compile.ReplaceAllString(content, "javascript:shangxiClose("+idStr+")")

	compile = regexp.MustCompile(`javascript:fanyiClose\(\d+\)`)
	content = compile.ReplaceAllString(content, "javascript:fanyiClose("+idStr+")")

	compile = regexp.MustCompile(`<a href="https://so.gushiwen.org/\w+\.\w+"\s*target="_blank">`)
	findAllString := compile.FindAllString(content, -1)
	content = compile.ReplaceAllString(content, "")

	compile = regexp.MustCompile(`<a href="https://so.gushiwen.org/\w+\.\w+">`)
	allString := compile.FindAllString(content, -1)
	content = compile.ReplaceAllString(content, "")
	content = strings.Replace(content, "</a>", "", len(findAllString)+len(allString))

	return content
}

//替换https://so.gushiwen.org超链接
func RemoveLinkHtml(content string) string {
	compile := regexp.MustCompile(`<a href="https://so.gushiwen.org/\w+\.\w+"\s*target="_blank">`)
	content = compile.ReplaceAllString(content, "")
	compile = regexp.MustCompile(`</a>`)
	content = compile.ReplaceAllString(content, "")
	return content
}
