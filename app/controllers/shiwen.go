/*
@Time : 2019/9/25 14:55
@Author : zxr
@File : nocdn
@Software: GoLand
*/
package controllers

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"poetry/app/bootstrap"
	"poetry/app/logic"
	"poetry/app/models"
	"poetry/config/define"
	templateHtml "poetry/libary/template"
	"poetry/tools"
	"sort"
	"strconv"
	"strings"
)

//诗文 控制器

//诗词分类列表页,根据分类名显示诗词列表页
func ShiWenList(w http.ResponseWriter, req *http.Request) {
	var (
		poetryList   []models.Content      //诗词列表
		authorIds    []int64               //作者ID集合
		authorData   map[int]models.Author //作者信息集合
		categoryData []models.Category     //分类列表
		contentData  define.ContentAll
		assign       map[string]interface{}
		cateName     string
		err          error
	)
	if cateName = req.FormValue("value"); len(cateName) < 2 {
		goto ErrorPage
	}
	if poetryList, err = logic.NewCategoryLogic().GetPoetryListByFilter(cateName, 0, 600); err != nil {
		goto ErrorPage
	}
	//获取诗文分类
	if categoryData, err = logic.NewCategoryLogic().GetCateByPositionLimit(define.PoetryShowPosition, 0, 95); err != nil {
		goto ErrorPage
	}
	sort.Slice(categoryData, func(i, j int) bool {
		return len(categoryData[i].CatName) > len(categoryData[j].CatName)
	})
	authorIds = logic.ExtractAuthorId(poetryList)
	//根据作者ID查询作者表数据
	if authorData, err = logic.NewAuthorLogic().GetAuthorInfoByIds(authorIds); err != nil {
		return
	}
	contentData = logic.NewContentLogic().ProcContentAuthorTagData(poetryList, authorData, nil)
	assign = make(map[string]interface{})
	assign["contentData"] = contentData.ContentArr
	assign["categoryData"] = categoryData
	assign["cateName"] = cateName
	assign["urlPath"] = define.PageShiWen
	assign["cdnDomain"] = bootstrap.G_Conf.CdnStaticDomain
	assign["webDomain"] = bootstrap.G_Conf.WebDomain
	assign["title"] = cateName
	assign["version"] = define.StaticVersion
	templateHtml.NewHtml(w).Display("shiwen/list.html", assign)
	logrus.Infof("%+v\n", poetryList)
	return
ErrorPage:
	if err == nil {
		err = errors.New("非法请求...请稍后重试")
	}
	templateHtml.NewHtml(w).DisplayErrorPage(err)
	return
}

//诗词详情页,根据诗词URL CRC值显示诗词详情页
func ShiWenDetail(w http.ResponseWriter, r *http.Request) {
	var (
		crcId          uint64
		poetryIdList   []int64
		html           *templateHtml.Html
		poetryRow      models.Content    //根据crcId查询的诗词ID信息
		poetryData     *define.Content   //发送给页面的诗词数据
		guessYouLike   []*define.Content //猜你喜欢
		guessYouLen    = 3               //猜你喜欢显示条数
		contentAll     define.ContentAll //诗词所有关联的信息
		notesData      []*models.Notes   //赏析和翻译信息
		creatBackData  *models.Notes     //创作背景
		err            error
		profileAddress string //作者头像
		assign         map[string]interface{}
		randomIdArr    []int64
	)
	html = templateHtml.NewHtml(w)
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
	if contentAll, err = contentLogic.GetPoetryContentAll(poetryIdList); err != nil {
		goto ShowErrorPage
	}
	for _, content := range contentAll.ContentArr {
		if content.PoetryInfo.Id == poetryRow.Id {
			poetryData = content
		} else {
			guessYouLike = append(guessYouLike, content)
		}
	}
	if poetryData == nil {
		goto ShowErrorPage
	}
	//获取翻译和赏析，创作背景数据
	if notesData, err = logic.NewShiWenLogic().GetAllNotesByPoetryId(poetryRow.Id, logic.NotesAll); err != nil {
		goto ShowErrorPage
	}
	if poetryData.PoetryInfo.CreatBackId > 0 {
		creatBackId := int(poetryData.PoetryInfo.CreatBackId)
		createBids := []int{creatBackId}
		if data, _ := logic.NewNotesLogic().GetNotesBytId(createBids); err == nil && len(data) > 0 {
			creatBackData = data[0]
			creatBackData.Content = tools.RemoveLinkHtml(creatBackData.Content)
		}
	}
	//头像地址
	profileAddress = logic.NewAuthorLogic().GetProfileAddress(poetryData.AuthorInfo)
	assign = make(map[string]interface{})
	assign["contentData"] = poetryData
	assign["guessYouLike"] = guessYouLike
	assign["notesList"] = notesData
	assign["creatBackData"] = creatBackData
	assign["authorProfileAddress"] = profileAddress
	assign["cdnDomain"] = bootstrap.G_Conf.CdnStaticDomain
	assign["webDomain"] = bootstrap.G_Conf.WebDomain
	assign["title"] = poetryData.PoetryInfo.Title
	assign["description"] = poetryData.PoetryInfo.Content
	assign["version"] = define.StaticVersion
	assign["urlPath"] = define.PageShiWen
	html.Display("shiwen/detail.html", assign)
	return
ShowErrorPage:
	html.DisplayErrorPage(err)
	return
}

// ajax 根据诗词URL crc32值获取注释和译文详情html
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

// ajax 根据赏析或译文id获取注释和译文详情html
func AjaxShiWenNotes(w http.ResponseWriter, r *http.Request) {
	var (
		idStr     string
		err       error
		id        int
		notesData []*models.Notes
		htmlStr   string
	)
	idStr = r.FormValue("id")
	if len(idStr) == 0 {
		goto OutPutEmptyStr
	}
	if id, err = strconv.Atoi(idStr); err != nil || id == 0 {
		goto OutPutEmptyStr
	}
	if notesData, err = logic.NewNotesLogic().GetNotesBytId([]int{id}); err != nil || len(notesData) == 0 {
		goto OutPutEmptyStr
	}
	htmlStr = logic.NewShiWenLogic().GetNotesDetailHtml(notesData[0])
	tools.OutputString(w, htmlStr)
	return
OutPutEmptyStr:
	tools.OutputString(w, "<div class='hr'></div><p>暂无内容</p>")
	return
}

//ajax根据赏析或译文id获取声音文件
func AjaxShiWenPlay(w http.ResponseWriter, r *http.Request) {
	var (
		idStr     string
		err       error
		id        int
		notesData []*models.Notes
		notes     *models.Notes
		songUrl   string
		htmlStr   string
	)
	idStr = r.FormValue("id")
	if len(idStr) == 0 {
		goto OutPutEmptyStr
	}
	if id, err = strconv.Atoi(idStr); err != nil || id == 0 {
		goto OutPutEmptyStr
	}
	if notesData, err = logic.NewNotesLogic().GetNotesBytId([]int{id}); err != nil || len(notesData) == 0 {
		goto OutPutEmptyStr
	}
	notes = notesData[0]
	songUrl = notes.PlayUrl
	if len(notes.FileName) > 0 {
		songUrl = define.CdnStoreDomain + "/" + notes.FileName
	}
	htmlStr = fmt.Sprintf(`<audio src="%s" autoplay></audio>`, songUrl)
	tools.OutputString(w, htmlStr)
	return
OutPutEmptyStr:
	tools.OutputString(w, `<audio src="ok.mp3" autoplay></audio>`)
	return
}
