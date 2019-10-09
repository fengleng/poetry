/*
@Time : 2019/10/8 15:51
@Author : zxr
@File : author
@Software: GoLand
*/
package controllers

import (
	"errors"
	"net/http"
	"poetry/app/bootstrap"
	"poetry/app/logic"
	"poetry/app/models"
	"poetry/config/define"
	templateHtml "poetry/libary/template"
	"poetry/tools"
)

//作者详情页
func AuthorDetail(w http.ResponseWriter, req *http.Request) {
	var (
		authorLogic    *logic.AuthorLogic
		authorName     string
		authorInfo     models.Author
		notesData      []define.AuthorNotes
		assign         map[string]interface{}
		html           *templateHtml.Html
		poetryListData define.ContentAll
		err            error
		randN          int
		orderFiled     = "id"
	)
	if authorName = req.FormValue("value"); len(authorName) == 0 {
		goto ErrorPage
	}
	authorLogic = logic.NewAuthorLogic()
	if authorInfo, err = authorLogic.GetAuthorInfoByName(authorName); err != nil {
		goto ErrorPage
	}
	if notesData, err = authorLogic.GetAuthorDetailDataListById(int(authorInfo.Id)); err != nil {
		goto ErrorPage
	}
	//根据作者ID获取3条诗词列表,这里生成随机数，0为id升序，大于0为id降序
	if randN = tools.RandInt(2); randN > 0 {
		orderFiled = "-id"
	}
	if poetryListData, err = logic.NewContentLogic().GetPoetryListByAuthorId(authorInfo, 0, 3, orderFiled); err != nil {
		goto ErrorPage
	}
	html = templateHtml.NewHtml(w)
	assign = make(map[string]interface{})
	assign["profileAddr"] = authorLogic.GetProfileAddress(authorInfo)
	assign["authorInfo"] = authorInfo
	assign["notesData"] = notesData
	assign["poetryListData"] = poetryListData.ContentArr
	assign["cdnDomain"] = bootstrap.G_Conf.CdnStaticDomain
	assign["webDomain"] = bootstrap.G_Conf.WebDomain
	assign["title"] = authorInfo.Author + "简介"
	assign["description"] = authorInfo.AuthorIntro
	html.Display("author/detail.html", assign)
	return
ErrorPage:
	if err == nil {
		err = errors.New("非法请求...请稍后重试")
	}
	templateHtml.NewHtml(w).DisplayErrorPage(err)
	return
}
