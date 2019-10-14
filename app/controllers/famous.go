/*
@Time : 2019/10/11 19:16
@Author : zxr
@File : famou
@Software: GoLand
*/
package controllers

import (
	"errors"
	"github.com/sirupsen/logrus"
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
		topCategory    []models.Category //顶级分类
		cateNameInfo   models.Category   //当前要查询的顶级分类的分类信息
		subCateInfo    models.Category   //当前的二级分类的分类信息
		subCategory    []models.Category //当前顶级分类的所有二级分类
		allSubCategory []models.Category //所有二级分类 列表
		topName        string            //要查询的顶级分类名
		subName        string            //要查询的顶级分类名下的子分类
		err            error
		limit          = 30
		countNum       int //名句总数
		page           int //当前页数
		countPage      int //总页数
		catId          []int
		famousData     []define.Famous //名句列表
		assign         map[string]interface{}
	)
	topName = strings.TrimSpace(req.FormValue("c"))
	subName = strings.TrimSpace(req.FormValue("t"))
	pageStr := req.FormValue("page")
	cateLogic := logic.NewCategoryLogic()
	//所有顶级分类
	if topCategory, err = cateLogic.GetCateByPositionLimit(define.FamousShowPosition, 0, 20); err != nil {
		goto ErrorPage
	}
	//查询所有子分类
	if allSubCategory, err = cateLogic.GetAllSubCateData(0, define.FamousShowPosition, 0, 1000); err != nil {
		goto ErrorPage
	}
	if len(topName) > 0 {
		if cateNameInfo = cateLogic.FindCateListByName(topCategory, topName); cateNameInfo.Id == 0 {
			err = errors.New("分类不存在")
			goto ErrorPage
		}
		//查二级分类
		catId, subCategory = cateLogic.FindSubCateByPid(allSubCategory, cateNameInfo.Id)
	}
	//根据子分类ID查询名句列表
	if len(topName) > 0 && len(subName) > 0 {
		if subCateInfo = cateLogic.FindCateListByName(subCategory, subName); subCateInfo.Id == 0 {
			err = errors.New("分类不存在")
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
	assign["allSubList"] = cateLogic.ProcSubCategory(topCategory, allSubCategory)
	assign["cateName"] = topName
	assign["tName"] = subName
	assign["page"] = page
	assign["countPage"] = countPage
	assign["cdnDomain"] = bootstrap.G_Conf.CdnStaticDomain
	assign["webDomain"] = bootstrap.G_Conf.WebDomain
	assign["nextPage"] = page + 1
	assign["prevPage"] = page - 1
	assign["pageUrl"] = tools.GetPageUrl(req.URL.String())
	assign["urlPath"] = define.PageFamous
	template.NewHtml(w).Display("famous/index.html", assign)
	return
ErrorPage:
	logrus.Infoln("err:", err)
	template.NewHtml(w).DisplayErrorPage(err)
	return
}
