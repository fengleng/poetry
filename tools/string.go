/*
@Time : 2019/9/24 18:47
@Author : zxr
@File : string
@Software: GoLand
*/
package tools

import (
	"net/http"
	"net/url"
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

//诗词详情页赏析和翻译数据和诗文详情替换所有的超链接，和字符中的ID，（AJAX获取赏析和翻译详情时用到了）
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
	compile = regexp.MustCompile(`<a href="https://so.gushiwen.org/\w+\.\w+"\s*>`)
	content = compile.ReplaceAllString(content, "")
	return content
}

//替换 <div class="cankao">HTML内容
func ReplaceCanKaoHtml(content string) string {
	compile := regexp.MustCompile(`(?msU)<div class="cankao">.*</div>`)
	str := compile.ReplaceAllString(content, "")
	return str
}

//替换诗词注释，翻译简介信息，诗词详情页展示用到 了
func PreContentHtml(content string) string {
	compile := regexp.MustCompile("<a.*</a>")
	content = compile.ReplaceAllString(content, "")
	content = ReplaceCanKaoHtml(content)
	content = strings.TrimPrefix(content, "<div>")
	content = strings.TrimSuffix(content, "</div>")
	if content[0:3] != "<p>" {
		content = "<p>" + content + "</p>"
	}
	return content
}

//获取当前URL
func GetPageUrl(urlStr string) (pageUrl string) {
	urlPath, _ := url.Parse(urlStr)
	if len(urlPath.Query()) == 0 {
		return urlStr + "?time=t88299332"
	}
	pageUrl = regexp.MustCompile("[&|?]page=\\d").ReplaceAllString(urlStr, "")
	if strings.Contains(pageUrl, "?") == false {
		pageUrl = strings.Replace(pageUrl, "&", "?", 1)
	}
	return
}
