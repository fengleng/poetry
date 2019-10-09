/*
@Time : 2019/10/9 19:05
@Author : zxr
@File : search
@Software: GoLand
*/
package controllers

import (
	"fmt"
	"net/http"
	"poetry/app/logic"
)

//诗词搜索页
func ShiWenSearch(w http.ResponseWriter, req *http.Request) {
	typeStr := req.FormValue("type") //搜索类型
	cstr := req.FormValue("cstr")    //搜索的具体值
	pageStr := req.FormValue("page") //当前页数
	fmt.Println(pageStr)
	logic.NewSearchLogic().GetSearchShiWenPoetryList(typeStr, cstr)
}
