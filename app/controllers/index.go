/*
@Time : 2019/9/19 16:46
@Author : zxr
@File : index
@Software: GoLand
*/
package controllers

import (
	"net/http"
	"poetry/app/bootstrap"
	"poetry/app/logic"
	"poetry/config/define"
	templateHtml "poetry/libary/template"
	"strconv"
)

//首页
func Index(writer http.ResponseWriter, request *http.Request) {
	/**
		          1.推荐数据，分页
			      2.诗词分类
		明日继续：
	 分类导航没做，名句导航没做，古籍导航没做
	*/
	var (
		err         error
		contentData define.ContentAll
		html        *templateHtml.Html
		assign      map[string]interface{}
		currPage    int    //当前页数
		offset      int    //偏移量
		limit       = 10   //显示多少条
		countPage   = 50   //总页数，先写死
		pageStr     string //URL传过来的当前页数
	)
	if pageStr = request.FormValue("page"); len(pageStr) == 0 {
		pageStr = "1"
	}
	currPage, _ = strconv.Atoi(pageStr)
	offset = (currPage - 1) * limit
	//count = logic.NewIndexLogic().GetRecommendCount()
	//countPage = int(math.Ceil(float64(count / limit)))
	if currPage > countPage {
		currPage = 1
		offset = 0
	}
	//获取推荐数据
	if contentData, err = logic.NewIndexLogic().GetSameDayRecommendPoetryData(offset, limit); err != nil {
		html.DisplayErrorPage(err)
		return
	}
	assign = make(map[string]interface{})
	assign["contentData"] = contentData
	assign["cdnDomain"] = bootstrap.G_Conf.CdnStaticDomain
	assign["currPage"] = currPage
	assign["nextPage"] = currPage + 1
	assign["prevPage"] = currPage - 1
	assign["countPage"] = countPage
	templateHtml.NewHtml(writer).Display("index.html", assign)
	return
}
