/*
@Time : 2019/10/15 19:06
@Author : zxr
@File : guwen
@Software: GoLand
*/
package controllers

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"math"
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

//古籍首页
func GuWenIndex(w http.ResponseWriter, req *http.Request) {
	var (
		allClassify   []*define.GuWenClassify //所有分类数据
		bookList      []models.AncientBook    //古籍列表
		rRemBookList  []models.AncientBook    //右侧古籍列表
		classifyLogic *logic.AncientClassifyLogic
		bookLogic     *logic.AncientBookLogic
		catIds        []int
		countPage     int
		page          int
		limit         = 6
		count         int64
		assign        map[string]interface{}
		err           error
	)
	classifyLogic = logic.NewAncientClassify()
	bookLogic = logic.NewAncientBook()
	typeStr := strings.TrimSpace(req.FormValue("type"))
	if pageStr := req.FormValue("page"); len(pageStr) > 0 {
		page, _ = strconv.Atoi(pageStr)
	}
	//分类列表
	if allClassify, err = classifyLogic.GetAllClassify(20); err != nil {
		goto ErrorPage
	}
	//搜索分类ID
	catIds = classifyLogic.FindCatIdListByCateName(allClassify, typeStr)
	if (len(typeStr) > 0 && len(catIds) > 0) || (len(typeStr) == 0) {
		count, _ = bookLogic.GetBookCountByCatId(catIds)
		if count > 0 {
			countPage = int(math.Ceil(float64(count) / float64(limit)))
		}
		if page > countPage || page == 0 {
			page = 1
		}
		if bookList, err = bookLogic.GetBookListLimitByCatId(catIds, (page-1)*limit, limit); err != nil {
			goto ErrorPage
		}
	}
	//右侧古籍列表
	rRemBookList, _ = bookLogic.GetBookListByLimit(0, 100)
	sort.Slice(rRemBookList, func(i, j int) bool {
		return len(rRemBookList[i].BookTitle) < len(rRemBookList[j].BookTitle)
	})
	assign = make(map[string]interface{})
	assign["allClassify"] = allClassify
	assign["bookList"] = bookList
	assign["rRemBookList"] = rRemBookList
	assign["count"] = count
	assign["countPage"] = countPage
	assign["page"] = page
	assign["nextPage"] = page + 1
	assign["prevPage"] = page - 1
	assign["typeStr"] = typeStr
	assign["cdnDomain"] = bootstrap.G_Conf.CdnStaticDomain
	assign["webDomain"] = bootstrap.G_Conf.WebDomain
	assign["pageUrl"] = tools.GetPageUrl(req.URL.String())
	assign["urlPath"] = define.PageGuWen
	templateHtml.NewHtml(w).Display("guwen/index.html", assign)
	return
ErrorPage:
	templateHtml.NewHtml(w).DisplayErrorPage(err)
	return
}

//ajax 根据bookID获取声音文件并返回播放声音的html
func BookPlay(w http.ResponseWriter, r *http.Request) {
	var (
		id       int
		err      error
		htmlStr  string
		songUrl  string
		bookData map[int]models.AncientBook
	)
	if idStr := r.FormValue("id"); len(idStr) > 0 {
		id, err = strconv.Atoi(idStr)
	}
	if id == 0 || err != nil {
		goto NoneMp3
	}
	if bookData, err = logic.NewAncientBook().GetBookListByIds([]int{id}); err != nil || len(bookData) == 0 {
		goto NoneMp3
	}
	songUrl = bookData[id].SongUrl
	if len(bookData[id].SongFilePath) > 0 {
		songUrl = define.CdnStoreDomain + "/" + bookData[id].SongFilePath
	}
	htmlStr = fmt.Sprintf(`<audio src="%s" autoplay></audio>`, songUrl)
	tools.OutputString(w, htmlStr)
	return
NoneMp3:
	tools.OutputString(w, `<audio src="ok.mp3" autoplay></audio>`)
	return
}

//古文详情页
func GuWenDetail(w http.ResponseWriter, req *http.Request) {
	var (
		bookId      uint64
		bookData    map[int]models.AncientBook
		catalogList []define.GuWenCatalogList
		assign      map[string]interface{}
		err         error
	)
	if bookId, err = logic.NewShiWenLogic().GetCrcIdByUrlPath(req.URL.Path); err != nil {
		goto ErrorPage
	}
	if bookData, err = logic.NewAncientBook().GetBookListByIds([]int{int(bookId)}); err != nil || len(bookData) == 0 {
		goto ErrorPage
	}
	//查所有目录列表
	if catalogList, err = logic.NewAncientCatalogueLogic().GetAllCatalogByBookId(int(bookId)); err != nil {
		goto ErrorPage
	}
	//todo  明天继续古籍详情页
	logrus.Infof("%+v", bookData)
	logrus.Infof("catalogList:%+v", catalogList)
	logrus.Infoln("err:", err)

	assign = make(map[string]interface{})
	assign["bookData"] = bookData[int(bookId)]
	assign["catalogList"] = catalogList
	assign["urlPath"] = define.PageGuWen
	assign["cdnDomain"] = bootstrap.G_Conf.CdnStaticDomain
	assign["webDomain"] = bootstrap.G_Conf.WebDomain
	templateHtml.NewHtml(w).Display("guwen/detail.html", assign)
	return
ErrorPage:
	templateHtml.NewHtml(w).DisplayErrorPage(err)
	return
}

//目录详情页
func GuWenBook(w http.ResponseWriter, req *http.Request) {
	var (
		dirId uint64
		err   error
	)
	if dirId, err = logic.NewShiWenLogic().GetCrcIdByUrlPath(req.URL.Path); err != nil || dirId == 0 {
		goto ErrorPage
	}

ErrorPage:
	templateHtml.NewHtml(w).DisplayErrorPage(err)
	return
}
