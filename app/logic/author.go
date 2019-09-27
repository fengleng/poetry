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
func (a *AuthorLogic) GetAuthorInfoByIds(ids []int64) (authorList map[int]models.Author, err error) {
	var authorData []models.Author
	authorList = make(map[int]models.Author)
	if authorData, err = a.authorModel.GetAuthorInfoByIds(ids); err != nil {
		return
	}
	for _, author := range authorData {
		authorList[int(author.Id)] = author
	}
	return authorList, err
}

//根据诗词总数倒序查询作者列表
func (a *AuthorLogic) GetListByOrderCountDesc(offset, limit int) (data []models.Author, err error) {
	return models.NewAuthor().GetListByOrderCountDesc(offset, limit)
}
