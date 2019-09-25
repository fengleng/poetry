/*
@Time : 2019/9/25 14:55
@Author : zxr
@File : nocdn
@Software: GoLand
*/
package controllers

import (
	"net/http"
	"poetry/app/logic"
	"poetry/tools"
	"strconv"
)

//注释和译文控制器

//ajax获取注释和译文
func AjaxShiWenCont(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		crcId uint64
	)
	id, value := r.FormValue("id"), r.FormValue("value")
	if len(id) == 0 || len(value) == 0 {
		goto OutPutEmptyStr
	}
	if crcId, err = strconv.ParseUint(id, 10, 32); err != nil {
		goto OutPutEmptyStr
	}
	if err = logic.NewNocdnLogic().GetPoetryTextTrans(uint32(crcId), value); err != nil {
		goto OutPutEmptyStr
	}
	return
OutPutEmptyStr:
	tools.OutputString(w, "")
	return
}
