/*
@Time : 2019/10/9 19:33
@Author : zxr
@File : search
@Software: GoLand
*/
package logic

import (
	"fmt"
	"poetry/app/models"
)

type Searcher interface {
	GetPoetryListByFilter(cstr string) []models.Content
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
	var searchMod Searcher
	switch typeStr {
	case searchAuthor:
		searchMod = NewDynastyLogic() //todo
	case searchTag:
		searchMod = NewDynastyLogic() //todo
	case searchDynasty:
		searchMod = NewDynastyLogic()
	}
	val := searchMod.GetPoetryListByFilter(cstr)
	//统一处理返回的诗词列表格式，
	fmt.Println(val)
}
