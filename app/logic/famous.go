/*
@Time : 2019/10/12 15:20
@Author : zxr
@File : famous
@Software: GoLand
*/
package logic

import (
	"net/url"
	"poetry/app/bootstrap"
	"poetry/app/models"
	"poetry/config/define"
	"strconv"
	"strings"
)

type FamousLogic struct {
}

type famUrlCrcMp map[uint32]models.Famous  //名句url crc32 map
type famUrlpathMp map[string]models.Famous //名句 URL path map

func NewFamousLogic() *FamousLogic {
	return &FamousLogic{}
}

//根据分类ID查询名句列表
func (f *FamousLogic) GetListByCatId(catIds []int, offset, limit int) (result []define.Famous, err error) {
	var famousData []models.Famous
	if famousData, err = models.NewFamous().GetListByCatId(catIds, offset, limit); err != nil || len(famousData) == 0 {
		return
	}
	result, err = f.GetFamousRetByFamousData(famousData)
	return
}

//根据分类ID查询名句总数
func (f *FamousLogic) GetCountByCatIds(catIds []int) (count int) {
	num, _ := models.NewFamous().GetCountByCatIds(catIds)
	return int(num)
}

//从表里查出名句数据，整理作者，诗词，古籍信息，合成名句列表集
func (f *FamousLogic) GetFamousRetByFamousData(famousData []models.Famous) (result []define.Famous, err error) {
	var (
		famCrcMap       famUrlCrcMp
		famPathMap      famUrlpathMp
		poetryCrcIds    []uint32 //诗词的SourceCrc32集合
		poetryListData  []models.Content
		ancientPaths    []string //古文的url path集合
		catalogListData []models.AncientCatalogue
		authorData      map[int]models.Author //作者信息集合
		bookData        bookMp
	)
	defer func() {
		famCrcMap = nil
		famPathMap = nil
		catalogListData = nil
		bookData = nil
		authorData = nil
	}()
	famCrcMap = make(famUrlCrcMp, len(famousData))
	famPathMap = make(famUrlpathMp, len(famousData))
	for _, famous := range famousData {
		if strings.Contains(famous.SourceUrl, "guwen") {
			parse, _ := url.Parse(famous.SourceUrl)
			ancientPaths = append(ancientPaths, parse.Path)
			famPathMap[parse.Path] = famous
		} else {
			poetryCrcIds = append(poetryCrcIds, famous.SourceCrc32)
			famCrcMap[famous.SourceCrc32] = famous
		}
	}
	if len(poetryCrcIds) == 0 && len(ancientPaths) == 0 {
		return
	}
	//查诗词表
	if len(poetryCrcIds) > 0 {
		if poetryListData, err = models.NewContent().GetContentByCrc32IdArr(poetryCrcIds); err != nil {
			return
		}
		authorIds := ExtractAuthorId(poetryListData)
		//根据作者ID查询作者表数据
		if authorData, err = NewAuthorLogic().GetAuthorInfoByIds(authorIds); err != nil {
			return
		}
	}
	//查古籍表
	if len(ancientPaths) > 0 {
		if catalogListData, err = models.NewAncientCatalogue().GetCatalogListByPaths(ancientPaths); err != nil {
			return
		}
		bookIds := ExtractBookId(catalogListData)
		bookData, _ = NewAncientBook().GetBookListByIds(bookIds)
	}
	for _, poetry := range poetryListData {
		var famousStr define.Famous
		if famous, ok := famCrcMap[poetry.SourceUrlCrc32]; ok {
			author := authorData[int(poetry.AuthorId)]
			famousStr.Content = famous.Content
			famousStr.PoetryTitle = poetry.Title
			famousStr.Ftype = 1
			famousStr.LinkUrl = bootstrap.G_Conf.WebDomain + "/shiwen/" + strconv.FormatUint(uint64(poetry.SourceUrlCrc32), 10) + define.UrlSuffix
			famousStr.AncientTitle = author.Author + "《" + poetry.Title + "》"
			famousStr.AuthorName = author.Author
			result = append(result, famousStr)
		}
	}
	for _, catLog := range catalogListData {
		if famous, ok := famPathMap[catLog.LinkUrl]; ok {
			var famousStr define.Famous
			book := bookData[int(catLog.BookId)]
			famousStr.Content = famous.Content
			famousStr.AncientTitle = catLog.CatalogTitle
			famousStr.Ftype = 2
			famousStr.LinkUrl = bootstrap.G_Conf.WebDomain + "/guwen/book/" + strconv.FormatUint(uint64(catLog.Id), 10) + define.UrlSuffix
			famousStr.AncientTitle = "《" + book.BookTitle + "." + catLog.CatalogTitle + "》"
			result = append(result, famousStr)
		}
	}
	return result, err
}
