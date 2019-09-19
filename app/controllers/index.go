/*
@Time : 2019/9/19 16:46
@Author : zxr
@File : index
@Software: GoLand
*/
package controllers

import (
	"net/http"
	"poetry/app/logic"
)

//首页
func Index(writer http.ResponseWriter, request *http.Request) {
	/**
	          1.推荐数据，分页
		      2.诗词分类
	*/
	logic.NewRecommendLogic().GetSameDayPoetryData(0, 10)
}
