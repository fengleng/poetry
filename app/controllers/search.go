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
)

//诗词搜索页
func ShiWenSearch(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello.....")
}
