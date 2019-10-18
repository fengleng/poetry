/*
@Time : 2019/10/18 19:03
@Author : zxr
@File : ancientAuthor
@Software: GoLand
*/
package logic

import "poetry/app/models"

type AncientAuthorLogic struct {
}

func NewAncientAuthorLogic() *AncientAuthorLogic {
	return &AncientAuthorLogic{}
}

//根据作者ID查询作者信息
func (a *AncientAuthorLogic) GetAuthorById(id int) (author models.AncientAuthor, err error) {
	return models.NewAncientAuthor().GetAuthorById(id)
}
