/*
@Time : 2019/10/8 15:51
@Author : zxr
@File : author
@Software: GoLand
*/
package controllers

import (
	"errors"
	"math"
	"net/http"
	"poetry/app/bootstrap"
	"poetry/app/logic"
	"poetry/app/models"
	"poetry/config/define"
	templateHtml "poetry/libary/template"
	"poetry/tools"
	"strconv"
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
	if authorInfo, err = authorLogic.GetAuthorInfoByName(authorName); err != nil || authorInfo.Id == 0 {
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
	assign["version"] = define.StaticVersion
	assign["urlPath"] = define.PageShiWen
	html.Display("author/detail.html", assign)
	return
ErrorPage:
	if err == nil {
		err = errors.New("非法请求...请稍后重试")
	}
	templateHtml.NewHtml(w).DisplayErrorPage(err)
	return
}

//作者诗词列表页
func PoetryList(w http.ResponseWriter, req *http.Request) {
	var (
		authorName     string                 //作者名字
		authorInfo     models.Author          //作者基础信息
		poetryListData define.ContentAll      //诗词列表
		assign         map[string]interface{} //发给模板的变量
		err            error                  //错误信息
		page           int                    //当前页
		countNum       int                    //作者的诗词总数
		countPage      int                    //总页数
		offset         int                    //当前偏移量
		limit          = 10                   //每页显示的诗词条数
	)
	contentLogic := logic.NewContentLogic()
	if authorName = req.FormValue("value"); len(authorName) == 0 {
		goto ErrorPage
	}
	if pageStr := req.FormValue("page"); len(pageStr) > 0 {
		page, _ = strconv.Atoi(pageStr)
	}
	if authorInfo, err = logic.NewAuthorLogic().GetAuthorInfoByName(authorName); err != nil || authorInfo.Id == 0 {
		goto ErrorPage
	}
	if countNum, err = contentLogic.GetContentCountByAuthorId(authorInfo.Id); err != nil {
		goto ErrorPage
	}
	countPage = int(math.Ceil(float64(countNum / limit)))
	if page <= 0 || page > countPage {
		page = 1
	}
	offset = (page - 1) * limit
	if poetryListData, err = contentLogic.GetPoetryListByAuthorId(authorInfo, offset, limit, "id"); err != nil {
		goto ErrorPage
	}
	assign = make(map[string]interface{})
	assign["profileAddr"] = logic.NewAuthorLogic().GetProfileAddress(authorInfo)
	assign["poetryList"] = poetryListData.ContentArr
	assign["authorInfo"] = authorInfo
	assign["page"] = page
	assign["countPage"] = countPage
	assign["nextPage"] = page + 1
	assign["prevPage"] = page - 1
	assign["cdnDomain"] = bootstrap.G_Conf.CdnStaticDomain
	assign["webDomain"] = bootstrap.G_Conf.WebDomain
	assign["title"] = authorInfo.Author + "的诗词全集_诗集、词集"
	assign["description"] = authorInfo.AuthorIntro
	assign["pageUrl"] = bootstrap.G_Conf.WebDomain + "/author/poetryList?value=" + authorName
	assign["version"] = define.StaticVersion
	assign["urlPath"] = define.PageShiWen
	templateHtml.NewHtml(w).Display("author/poetryList.html", assign)
	return
ErrorPage:
	if err == nil {
		err = errors.New("非法请求...请稍后重试")
	}
	templateHtml.NewHtml(w).DisplayErrorPage(err)
	return
}
