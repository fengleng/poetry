/*
@Time : 2019/9/19 16:46
@Author : zxr
@File : index
@Software: GoLand
*/
package controllers

import (
	"net/http"
	"poetry/app/logic"
	"poetry/config/define"
	templateHtml "poetry/libary/template"
)

//首页
func Index(writer http.ResponseWriter, request *http.Request) {
	/**
	          1.推荐数据，分页
		      2.诗词分类
	*/
	var (
		err         error
		contentData define.ContentAll
		html        *templateHtml.Html
	)
	html = templateHtml.NewHtml(writer, request)
	if contentData, err = logic.NewRecommendLogic().GetSameDayPoetryData(0, 10); err != nil {
		html.DisplayErrorPage(err)
		return
	}
	html.Display("index.html", contentData)
}
