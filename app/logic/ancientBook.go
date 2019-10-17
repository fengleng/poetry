/*
@Time : 2019/9/27 12:19
@Author : zxr
@File : ancientBook
@Software: GoLand
*/
package logic

import (
	"poetry/app/models"
	"poetry/config/define"
)

type AncientBookLogic struct {
}
type bookMp map[int]models.AncientBook

func NewAncientBook() *AncientBookLogic {
	return &AncientBookLogic{}
}

//根据偏移量查询古籍-书名列表
func (a *AncientBookLogic) GetBookListByLimit(offset, limit int) (data []models.AncientBook, err error) {
	return models.NewAncientBook().GetBookListByLimit(offset, limit)
}

//根据分类ID查询书籍列表
func (a *AncientBookLogic) GetBookListLimitByCatId(catId []int, offset, limit int) (data []models.AncientBook, err error) {
	if data, err = models.NewAncientBook().GetBookListLimitByCatId(catId, offset, limit); err != nil {
		return
	}
	for k, book := range data {
		book.CoverChart = a.GetBookCoverImage(book)
		data[k] = book
	}
	return
}

//获取封面图地址
func (a *AncientBookLogic) GetBookCoverImage(book models.AncientBook) (coverImage string) {
	if len(book.CoverChartPath) > 0 {
		coverImage = define.CdnStoreDomain + "/" + book.CoverChartPath
	}
	if coverImage == "" && len(book.CoverChart) > 0 {
		coverImage = book.CoverChart
	}
	if coverImage == "" {
		coverImage = define.CdnStoreDomain + "/default.png"
	}
	return
}

//根据分类ID查询书籍总数
func (a *AncientBookLogic) GetBookCountByCatId(catId []int) (num int64, err error) {
	return models.NewAncientBook().GetBookCountByCatId(catId)
}

//根据ID查询书名列表
func (a *AncientBookLogic) GetBookListByIds(bookIds []int) (result bookMp, err error) {
	var bookData []models.AncientBook
	if bookData, err = models.NewAncientBook().GetBookListByIds(bookIds); err != nil {
		return
	}
	result = make(bookMp, len(bookData))
	for _, book := range bookData {
		result[int(book.Id)] = book
	}
	return
}
