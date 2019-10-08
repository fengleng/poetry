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
	templateHtml "poetry/libary/template"
	"poetry/tools"
	"strings"
)

//作者详情页
func AuthorDetail(w http.ResponseWriter, req *http.Request) {
	var (
		authorLogic *logic.AuthorLogic
		authorName  string
		authorInfo  models.Author
		notesData   []models.Notes
		assign      map[string]interface{}
		html        *templateHtml.Html
		err         error
	)
	if authorName = req.FormValue("value"); len(authorName) == 0 {
		goto ErrorPage
	}
	authorName = strings.TrimSpace(authorName)
	authorLogic = logic.NewAuthorLogic()
	if authorInfo, err = authorLogic.GetAuthorInfoByName(authorName); err != nil {
		goto ErrorPage
	}
	authorInfo.AuthorIntro = tools.TrimRightHtml(authorInfo.AuthorIntro)
	if notesData, err = authorLogic.GetAuthorDetailDataListById(int(authorInfo.Id)); err != nil {
		goto ErrorPage
	}
	html = templateHtml.NewHtml(w)
	assign = make(map[string]interface{})
	assign["profileAddr"] = authorLogic.GetProfileAddress(authorInfo)
	assign["authorInfo"] = authorInfo
	assign["notesData"] = notesData
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
