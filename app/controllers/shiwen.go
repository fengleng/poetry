/*
@Time : 2019/9/25 14:55
@Author : zxr
@File : nocdn
@Software: GoLand
*/
package controllers

import (
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

//诗文 控制器

//诗词详情页
func ShiWenIndex(w http.ResponseWriter, r *http.Request) {
	var (
		crcId          uint64
		poetryIdList   []int64
		html           *template.Html
		poetryRow      models.Content    //根据crcId查询的诗词ID信息
		poetryData     *define.Content   //发送给页面的诗词数据
		guessYouLike   []*define.Content //猜你喜欢
		guessYouLen    = 3               //猜你喜欢显示条数
		contentAll     define.ContentAll //诗词所有关联的信息
		notesList      []*models.Notes   //赏析和翻译信息
		creatBackData  []*models.Notes   //创作背景
		err            error
		profileAddress string //作者头像
		assign         map[string]interface{}
		randomIdArr    []int64
	)
	html = template.NewHtml(w)
	contentLogic := logic.NewContentLogic()
	if crcId, err = logic.NewShiWenLogic().GetCrcIdByUrlPath(r.URL.Path); err != nil {
		goto ShowErrorPage
	}
	//获取诗词ID
	if poetryRow, err = contentLogic.GetContentByCrc32Id(uint32(crcId)); err != nil {
		goto ShowErrorPage
	}
	//随机生成3个随机ID，用于获取猜你喜欢诗词
	randomIdArr = tools.RandInt64Slice(guessYouLen, define.MaxIdNumber)
	//获取诗词详情信息和猜你喜欢
	poetryIdList = append([]int64{int64(poetryRow.Id)}, randomIdArr...)
	if contentAll, err = contentLogic.GetPoetryContentAll(poetryIdList); err != nil || len(contentAll.ContentArr) == 0 {
		goto ShowErrorPage
	}
	for _, content := range contentAll.ContentArr {
		if content.PoetryInfo.Id == poetryRow.Id {
			poetryData = content
		} else {
			guessYouLike = append(guessYouLike, content)
		}
	}
	//获取翻译和赏析，创作背景数据
	if notesList, err = logic.NewShiWenLogic().GetAllNotesByPoetryId(poetryRow.Id, logic.NotesAll); err != nil {
		goto ShowErrorPage
	}
	if poetryData.PoetryInfo.CreatBackId > 0 {
		creatBackId := int(poetryData.PoetryInfo.CreatBackId)
		createBids := []int{creatBackId}
		if creatBackData, err = logic.NewNotesLogic().GetNotesBytId(createBids); err == nil {
			notesList = append(notesList, creatBackData...)
		}
	}
	//头像地址
	profileAddress = logic.NewAuthorLogic().GetProfileAddress(poetryData.AuthorInfo)
	assign = make(map[string]interface{})
	assign["contentData"] = poetryData
	assign["guessYouLike"] = guessYouLike
	assign["notesList"] = notesList
	assign["authorProfileAddress"] = profileAddress
	assign["cdnDomain"] = bootstrap.G_Conf.CdnStaticDomain
	assign["webDomain"] = bootstrap.G_Conf.WebDomain
	assign["title"] = poetryData.PoetryInfo.Title
	assign["description"] = poetryData.PoetryInfo.Content
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
	if notesData, err = swLogic.GetOneNotesDetailByCrcId(uint32(crcId), value); err != nil || notesData == nil || len(notesData.Content) == 0 {
		goto OutPutEmptyStr
	}
	htmlStr = swLogic.GetNotesContentHtml(notesData, value)
	tools.OutputString(w, htmlStr)
	return
OutPutEmptyStr:
	tools.OutputString(w, "<div class='hr'></div><p>暂无内容</p>")
	return
}
