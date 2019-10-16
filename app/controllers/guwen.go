/*
@Time : 2019/10/15 19:06
@Author : zxr
@File : guwen
@Software: GoLand
*/
package controllers

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"poetry/app/logic"
	"poetry/app/models"
	"poetry/config/define"
	templateHtml "poetry/libary/template"
	"strconv"
	"strings"
)

//古籍首页
func GuWenIndex(w http.ResponseWriter, req *http.Request) {
	var (
		allClassify   []*define.GuWenClassify //所有分类数据
		bookList      []models.AncientBook    //书籍列表
		classifyLogic *logic.AncientClassifyLogic
		bookLogic     *logic.AncientBookLogic
		catIds        []int
		typeStr       string
		typeClass     models.AncientClassify
		page          int
		limit         = 10
		count         int64
		assign        map[string]interface{}
		err           error
	)
	classifyLogic = logic.NewAncientClassify()
	bookLogic = logic.NewAncientBook()
	typeStr = strings.TrimSpace(req.FormValue("type"))
	first := req.FormValue("first")
	if pageStr := req.FormValue("page"); len(pageStr) > 0 {
		page, _ = strconv.Atoi(pageStr)
	}
	//分类列表
	if allClassify, err = classifyLogic.GetAllClassify(20); err != nil {
		goto ErrorPage
	}
	//搜索子分类
	if len(typeStr) > 0 && first == "0" {
		typeClass = classifyLogic.FindClassifyListByCateName(allClassify, typeStr)
		if typeClass.Id > 0 {
			catIds = []int{typeClass.Id}
		}
	}
	//搜索顶级分类
	if len(typeStr) > 0 && first == "1" {
		if typeClass, err = classifyLogic.GetCategoryDataByName(typeStr); err != nil {
			goto ErrorPage
		}
		catIds = classifyLogic.FindPidByClassifyData(allClassify, typeClass.Id)
	}
	if len(typeStr) > 0 && len(catIds) == 0 {
		goto ClassPage
	}
	count, _ = bookLogic.GetBookCountByCatId(catIds)
	if bookList, err = bookLogic.GetBookListLimitByCatId(catIds, (page-1)*limit, limit); err != nil {
		goto ErrorPage
	}
	//todo 明天继续古文列表
	logrus.Infoln("%+v", bookList)
	assign = make(map[string]interface{})
	assign["allClassify"] = allClassify
	assign["bookList"] = bookList
	assign["count"] = count
	assign["page"] = page
	assign["typeStr"] = typeStr
	assign["urlPath"] = define.PageGuWen
	templateHtml.NewHtml(w).Display("guwen/index.html", assign)
	return
ClassPage:
	assign = make(map[string]interface{})
	assign["allClassify"] = allClassify
	assign["typeStr"] = typeStr
	assign["urlPath"] = define.PageGuWen
	assign["bookList"] = bookList
	templateHtml.NewHtml(w).Display("guwen/index.html", assign)
	return
ErrorPage:
	templateHtml.NewHtml(w).DisplayErrorPage(err)
	return
}

//古文详情页
func GuWenDetail(w http.ResponseWriter, req *http.Request) {

}
