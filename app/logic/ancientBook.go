/*
@Time : 2019/9/27 12:19
@Author : zxr
@File : ancientBook
@Software: GoLand
*/
package logic

import "poetry/app/models"

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
