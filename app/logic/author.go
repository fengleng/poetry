/*
@Time : 2019/9/20 19:10
@Author : zxr
@File : author
@Software: GoLand
*/
package logic

import "poetry/app/models"

type AuthorLogic struct {
	authorModel *models.Author
}

func NewAuthorLogic() *AuthorLogic {
	return &AuthorLogic{
		authorModel: models.NewAuthor(),
	}
}

//根据id查询作者信息
func (a *AuthorLogic) GetAuthorInfoByIds(ids []int64) (data []models.Author, err error) {
	data, err = a.authorModel.GetAuthorInfoByIds(ids)
	return
}
