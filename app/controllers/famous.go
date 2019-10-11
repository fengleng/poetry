/*
@Time : 2019/10/11 19:16
@Author : zxr
@File : famou
@Software: GoLand
*/
package controllers

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"poetry/app/logic"
	"poetry/app/models"
	"poetry/config/define"
	"poetry/libary/template"
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
		subCategory  []models.Category //二级分类
		cateName     string            //要查询的顶级分类名
		tName        string            //要查询的顶级分类名下的子分类
		err          error
	)
	cateName = req.FormValue("c")
	tName = req.FormValue("t")
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
	}
	//根据子分类ID查询名句列表
	if len(cateName) > 0 && len(tName) > 0 {

	}

	logrus.Infof("topCategory %+v:", topCategory)
	logrus.Infof("subCategory %+v\n\n", subCategory)

	return
ErrorPage:
	template.NewHtml(w).DisplayErrorPage(err)
	return
}
