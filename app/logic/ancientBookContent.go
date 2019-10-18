/*
@Time : 2019/9/27 12:19
@Author : zxr
@File : ancientBook
@Software: GoLand
*/
package logic

import "poetry/app/models"

type AncientBookContentLogic struct {
}

func NewAncientBookContentLogic() *AncientBookContentLogic {
	return &AncientBookContentLogic{}
}

//根据目录ID查询正文内容
func (a *AncientBookContentLogic) GetBookContentByCataLogId(id int) (data models.AncientBookContent, err error) {
	return models.NewAncientBookContent().GetBookContentByCataLogId(id)
}

//根据ID查询正文内容
func (a *AncientBookContentLogic) GetBookContentById(id int) (data models.AncientBookContent, err error) {
	return models.NewAncientBookContent().GetBookContentById(id)
}
