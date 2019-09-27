/*
@Time : 2019/9/20 19:10
@Author : zxr
@File : author
@Software: GoLand
*/
package logic

import (
	"poetry/app/models"
	"poetry/config/define"
)

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

//获取作者头像地址
func (a *AuthorLogic) GetProfileAddress(author models.Author) (profileAddress string) {
	profileAddress = author.PhotoUrl
	if len(author.PhotoFileName) > 0 {
		profileAddress = define.CdnStoreDomain + "/" + author.PhotoFileName
	}
	if len(profileAddress) == 0 {
		profileAddress = define.CdnStoreDomain + "/default.png"
	}
	return profileAddress
}
