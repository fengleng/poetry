/*
@Time : 2019/10/9 19:05
@Author : zxr
@File : search
@Software: GoLand
*/
package controllers

import (
	"github.com/pkg/errors"
	"net/http"
	"poetry/app/bootstrap"
	"poetry/app/logic"
	"poetry/app/models"
	"poetry/config/define"
	"poetry/libary/template"
	"poetry/tools"
	"sort"
)

//诗词搜索页
func ShiWenSearch(w http.ResponseWriter, req *http.Request) {
	var (
		dynastyData  []models.Dynasty  //朝代数据
		authorData   []models.Author   //作者数据
		categoryData []models.Category //分类数据
		assign       map[string]interface{}
		offset       = 0
		limit        = 165 //作者，分类，查询的总条数
		limitOffset  = 100 //分割条数，右边显示 60条数据
		err          error
	)
	typeStr := req.FormValue("type") //搜索类型
	cstr := req.FormValue("cstr")    //搜索的具体值
	//pageStr := req.FormValue("page") //当前页数

	//todo....明天继续
	logic.NewSearchLogic().GetSearchShiWenPoetryList(typeStr, cstr)

	//查询100个诗词分类 offset随机
	offset = tools.RandInt(830)
	if categoryData, err = logic.NewCategoryLogic().GetCateByPositionLimit(define.PoetryShowPosition, offset, limit); err != nil {
		goto ErrorPage
	}
	sort.Slice(categoryData, func(i, j int) bool {
		return len(categoryData[i].CatName) < len(categoryData[j].CatName)
	})
	//查询100个作者 offset随机
	offset = tools.RandInt(900)
	if authorData, err = logic.NewAuthorLogic().GetListByOrderCountDesc(offset, limit); err != nil {
		goto ErrorPage
	}
	sort.Slice(authorData, func(i, j int) bool {
		return len(authorData[i].Author) < len(authorData[j].Author)
	})
	//查询所有朝代
	if dynastyData, err = logic.NewDynastyLogic().GetAll(0, limit); err != nil {
		goto ErrorPage
	}
	assign = make(map[string]interface{})
	assign["categoryData"] = categoryData[:limitOffset]
	assign["rightCategoryData"] = categoryData[limitOffset:]
	assign["authorData"] = authorData[:limitOffset]
	assign["rightAuthorData"] = authorData[limitOffset:]
	assign["dynastyData"] = dynastyData
	assign["cdnDomain"] = bootstrap.G_Conf.CdnStaticDomain
	assign["webDomain"] = bootstrap.G_Conf.WebDomain
	assign["title"] = define.WebTitle
	assign["version"] = define.StaticVersion
	//显示HTML
	template.NewHtml(w).Display("search/shiwen.html", assign)
	return
ErrorPage:
	if err == nil {
		err = errors.New("非法请求.... 请稍后重试")
	}
	template.NewHtml(w).DisplayErrorPage(err)
	return
}
