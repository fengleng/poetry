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
	"poetry/app/models"
	"poetry/tools"
	"strconv"
)

//注释和译文控制器

//ajax获取注释和译文
func AjaxShiWenCont(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		crcId      uint64
		notesData  *models.Notes
		poetryData models.Content
		swLogic    *logic.ShiWenLogic
		htmlStr    string
	)
	id, value := r.FormValue("id"), r.FormValue("value")
	if len(id) == 0 || len(value) == 0 {
		goto OutPutEmptyStr
	}
	if crcId, err = strconv.ParseUint(id, 10, 32); err != nil {
		goto OutPutEmptyStr
	}
	if poetryData, err = logic.NewContentLogic().GetContentByCrc32Id(uint32(crcId)); err != nil || poetryData.Id == 0 {
		goto OutPutEmptyStr
	}
	swLogic = logic.NewShiWenLogic()
	if notesData, err = swLogic.GetNotesByPoetryCrcId(poetryData.Id, value); err != nil {
		goto OutPutEmptyStr
	}
	htmlStr = swLogic.GetNotesContentHtml(notesData, poetryData)
	tools.OutputString(w, htmlStr)
	return
OutPutEmptyStr:
	tools.OutputString(w, "")
	return
}
