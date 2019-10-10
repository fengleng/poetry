/*
@Time : 2019/10/9 19:33
@Author : zxr
@File : search
@Software: GoLand
*/
package logic

import (
	"github.com/sirupsen/logrus"
	"poetry/app/models"
	"strings"
)

type Searcher interface {
	GetPoetryListByFilter(cstr string, offset, limit int) ([]models.Content, error)
}

type SearchLogic struct {
}

func NewSearchLogic() *SearchLogic {
	return &SearchLogic{}
}

const (
	searchTag     = "tag"     //按tag标签搜索
	searchDynasty = "dynasty" //按朝代搜索
	searchAuthor  = "author"  //按作者搜索
)

//诗文搜索
func (s *SearchLogic) GetSearchShiWenPoetryList(typeStr, cstr string) {
	var (
		searchMod  Searcher
		poetryList []models.Content
		err        error
	)
	switch typeStr {
	case searchAuthor:
		searchMod = NewAuthorLogic()
	case searchTag:
		searchMod = NewCategoryLogic()
	case searchDynasty:
		searchMod = NewDynastyLogic()
	}
	cstr = strings.TrimSpace(cstr)
	if poetryList, err = searchMod.GetPoetryListByFilter(cstr, 0, 10); err != nil {
		return
	}
	//统一处理返回的诗词列表格式，
	logrus.Infof("%+v\n", poetryList)
}
