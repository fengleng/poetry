/*
@Time : 2019/9/24 18:47
@Author : zxr
@File : string
@Software: GoLand
*/
package tools

import (
	"strings"
)

//为字符串添加HTML标签
func AddHtmlLabel(content string) string {
	if len(content) == 0 {
		return ""
	}
	content = strings.Replace(content, "。", "。</p><p>", -1)
	return "<p>" + content
}
