/*
@Time : 2019/10/11 19:16
@Author : zxr
@File : famou
@Software: GoLand
*/
package controllers

import (
	"math"
	"net/http"
	"poetry/app/bootstrap"
	"poetry/app/logic"
	"poetry/app/models"
	"poetry/config/define"
	"poetry/libary/template"
	"poetry/tools"
	"strconv"
	"strings"
)

//名句 控制器

//名句 首页
func FamousIndex(w http.ResponseWriter, req *http.Request) {
	/*
			1.查询名句分类-》一级分类和二级分类
			2.如果有分类名，则根据分类名查名句列表，如果没有分类名，则查默认的名句列表
		    3.根据名句的URL去查询诗词表或古籍表对应的ID，生成链接地址
	*/
	var (
		topCategory  []models.Category //顶级分类
		cateNameInfo models.Category   //当前要查询的顶级分类的分类信息
		subCateInfo  models.Category   //当前要查询的二级分类的分类信息
		subCategory  []models.Category //二级分类
		cateName     string            //要查询的顶级分类名
		tName        string            //要查询的顶级分类名下的子分类
		err          error
		limit        = 10
		countNum     int //名句总数
		page         int //当前页数
		countPage    int //总页数
		catId        []int
		famousData   []define.Famous //名句列表
		assign       map[string]interface{}
	)
	cateName = strings.TrimSpace(req.FormValue("c"))
	tName = strings.TrimSpace(req.FormValue("t"))
	pageStr := req.FormValue("page")
	cateLogic := logic.NewCategoryLogic()
	if topCategory, err = cateLogic.GetCateByPositionLimit(define.FamousShowPosition, 0, 20); err != nil {
		goto ErrorPage
	}
	//查询所有的子分类，根据顶级分类ID查询名句列表
	if len(cateName) > 0 {
		if cateNameInfo, err = cateLogic.GetCateInfoByName(cateName, define.FamousShowPosition); err != nil || cateNameInfo.Id == 0 {
			goto ErrorPage
		}
		//查二级分类
		subCategory, err = cateLogic.GetSubCategoryData(cateNameInfo.Id, define.FamousShowPosition, 0, 100)
		catId = make([]int, len(subCategory))
		for k, subCat := range subCategory {
			catId[k] = subCat.Id
		}
	}
	//根据子分类ID查询名句列表
	if len(cateName) > 0 && len(tName) > 0 {
		if subCateInfo, err = cateLogic.GetCateInfoByNameAndPid(cateNameInfo.Id, tName); err != nil || subCateInfo.Id == 0 {
			goto ErrorPage
		}
		catId = []int{subCateInfo.Id}
	}
	// 查名句列表
	countNum = logic.NewFamousLogic().GetCountByCatIds(catId)
	countPage = int(math.Ceil(float64(countNum) / float64(limit)))
	page, _ = strconv.Atoi(pageStr)
	if page <= 0 || page > countPage {
		page = 1
	}
	if famousData, err = logic.NewFamousLogic().GetListByCatId(catId, (page-1)*limit, limit); err != nil {
		goto ErrorPage
	}
	assign = make(map[string]interface{})
	assign["famousData"] = famousData
	assign["topCategory"] = topCategory
	assign["subCategory"] = subCategory
	assign["cateName"] = cateName
	assign["tName"] = tName
	assign["page"] = page
	assign["countPage"] = countPage
	assign["cdnDomain"] = bootstrap.G_Conf.CdnStaticDomain
	assign["webDomain"] = bootstrap.G_Conf.WebDomain
	assign["nextPage"] = page + 1
	assign["prevPage"] = page - 1
	assign["pageUrl"] = tools.GetPageUrl(req.URL.String())
	assign["urlPath"] = ""
	template.NewHtml(w).Display("famous/index.html", assign)
	return
ErrorPage:
	template.NewHtml(w).DisplayErrorPage(err)
	return
}
