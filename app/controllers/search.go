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
	"strconv"
	"strings"
)

//诗词搜索页
func ShiWenSearch(w http.ResponseWriter, req *http.Request) {
	var (
		dynastyData  []models.Dynasty  //朝代数据
		authorData   []models.Author   //作者数据
		categoryData []models.Category //分类数据
		poetryList   define.ContentAll //诗词列表信息
		assign       map[string]interface{}
		page         int //当前页数
		offset       = 0
		pOffset      = 0   //诗词列表偏移量
		pLimit       = 10  //诗词列表每页显示的条数
		limit        = 165 //作者，分类，查询的总条数
		limitOffset  = 100 //分割条数，右边显示 60条数据
		err          error
	)
	typeStr := req.FormValue("type") //搜索类型
	cstr := req.FormValue("cstr")    //搜索的具体值
	if pageStr := req.FormValue("page"); len(pageStr) > 0 {
		page, _ = strconv.Atoi(pageStr)
	}
	if page == 0 {
		page = 1
	}
	if len(cstr) > 1 {
		//走搜索获取取诗词列表
		cstr = strings.TrimSpace(cstr)
		pOffset = (page - 1) * pLimit
		poetryList, _ = logic.NewSearchLogic().GetSearchShiWenPoetryList(typeStr, cstr, pOffset, pLimit)
	}
	if len(cstr) == 0 {
		//默认取推荐表 获取诗词列表，从第30页取，避免与首页数据相同
		pOffset = (page + 30 - 1) * pLimit
		if poetryList, err = logic.NewRecommendLogic().GetSameDayRecommendPoetryData(pOffset, pLimit); err != nil {
			goto ErrorPage
		}
	}
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
	assign["poetryList"] = poetryList.ContentArr
	assign["cdnDomain"] = bootstrap.G_Conf.CdnStaticDomain
	assign["webDomain"] = bootstrap.G_Conf.WebDomain
	assign["title"] = define.WebTitle
	assign["version"] = define.StaticVersion
	template.NewHtml(w).Display("search/shiwen.html", assign)
	return
ErrorPage:
	if err == nil {
		err = errors.New("非法请求.... 请稍后重试")
	}
	template.NewHtml(w).DisplayErrorPage(err)
	return
}
