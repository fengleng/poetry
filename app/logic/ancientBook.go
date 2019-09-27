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

func NewAncientBook() *AncientBookLogic {
	return &AncientBookLogic{}
}

//根据偏移量查询古籍-书名列表
func (a *AncientBookLogic) GetBookListByLimit(offset, limit int) (data []models.AncientBook, err error) {
	return models.NewAncientBook().GetBookListByLimit(offset, limit)
}
