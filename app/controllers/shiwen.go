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
	"poetry/config/define"
	"poetry/libary/template"
	"poetry/tools"
	"strconv"
	"strings"
)

//诗文 控制器

//诗词详情页
func ShiWenIndex(w http.ResponseWriter, r *http.Request) {
	var (
		crcId        uint64
		poetryIdList []int64
		html         *template.Html
		content      models.Content    //诗词信息
		contentAll   define.ContentAll //诗词所有关联的信息
		err          error
		assign       map[string]interface{}
	)
	html = template.NewHtml(w)
	contentLogic := logic.NewContentLogic()
	if crcId, err = logic.NewShiWenLogic().GetCrcIdByUrlPath(r.URL.Path); err != nil {
		goto ShowErrorPage
	}
	if content, err = contentLogic.GetContentByCrc32Id(uint32(crcId)); err != nil {
		goto ShowErrorPage
	}
	poetryIdList = []int64{int64(content.Id)}
	if contentAll, err = contentLogic.GetPoetryContentAll(poetryIdList); err != nil {
		goto ShowErrorPage
	}
	assign = make(map[string]interface{})
	assign["contentData"] = contentAll
	html.Display("sw_detail.html", assign)
	return
ShowErrorPage:
	html.DisplayErrorPage(err)
	return
}

//ajax获取注释和译文详情html
func AjaxShiWenCont(w http.ResponseWriter, r *http.Request) {
	var (
		err       error
		crcId     uint64
		notesData *models.Notes
		swLogic   *logic.ShiWenLogic
		htmlStr   string
	)
	defer func() {
		notesData = nil
		swLogic = nil
	}()
	id, value := r.FormValue("id"), r.FormValue("value")
	if len(id) == 0 || len(value) == 0 {
		goto OutPutEmptyStr
	}
	if crcId, err = strconv.ParseUint(id, 10, 32); err != nil {
		goto OutPutEmptyStr
	}
	swLogic = logic.NewShiWenLogic()
	value = strings.ToLower(value)
	if notesData, err = swLogic.GetOneNotesDetailByCrcId(uint32(crcId), value); err != nil || len(notesData.Content) == 0 {
		goto OutPutEmptyStr
	}
	htmlStr = swLogic.GetNotesContentHtml(notesData, value)
	tools.OutputString(w, htmlStr)
	return
OutPutEmptyStr:
	tools.OutputString(w, "<div class='hr'></div><p>暂无内容</p>")
	return
}
