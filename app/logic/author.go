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
	"poetry/tools"
	"strings"
)

type AuthorLogic struct {
	authorModel *models.Author
}

func NewAuthorLogic() *AuthorLogic {
	return &AuthorLogic{
		authorModel: models.NewAuthor(),
	}
}

//根据id查询作者信息基础信息
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

//根据朝代ID查询作者列表
func (a *AuthorLogic) GetListByDynastyId(dynastyId int, offset, limit int) (data []models.Author, err error) {
	return models.NewAuthor().GetListByDynastyId(dynastyId, offset, limit)
}

//根据朝代ID查询作者总数
func (a *AuthorLogic) GetCountByDynastyId(dynastyId int) (count int64, err error) {
	return models.NewAuthor().GetCountByDynastyId(dynastyId)
}

//根据作者名字获取作者基础信息
func (a *AuthorLogic) GetAuthorInfoByName(name string) (data models.Author, err error) {
	name = strings.TrimSpace(name)
	if data, err = a.authorModel.GetAuthorInfoByName(name); err != nil {
		return
	}
	data.AuthorIntro = tools.TrimRightHtml(data.AuthorIntro)
	return
}

//根据作者ID查询作者详情资料信息
func (a *AuthorLogic) GetAuthorDetailDataListById(id int) (authorNotes []define.AuthorNotes, err error) {
	var (
		authorData   []models.AuthorData
		notesData    []models.Notes
		notesIds     []int
		notesDataIds map[int]int
	)
	if authorData, err = models.NewAuthorData().GetAuthorDetailDataById(id); err != nil || len(authorData) == 0 {
		return
	}
	notesIds = make([]int, len(authorData))
	notesDataIds = make(map[int]int)
	for k, data := range authorData {
		notesIds[k] = data.NotesId
		notesDataIds[data.NotesId] = data.DataId
	}
	if notesData, err = models.NewNotes().GetNotesByIds(notesIds); err != nil || len(notesData) == 0 {
		return
	}
	authorNotes = make([]define.AuthorNotes, len(notesData))
	for k, notes := range notesData {
		authorNotes[k].DataId = notesDataIds[notes.Id]
		authorNotes[k].Notes = notes
	}
	return
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

//根据作者名字搜索诗词列表
func (a *AuthorLogic) GetPoetryListByFilter(authorName string, offset, limit int) (data []models.Content, err error) {
	var authorInfo models.Author
	if authorInfo, err = a.authorModel.GetAuthorInfoByName(authorName); err != nil || authorInfo.Id == 0 {
		return
	}
	data, err = models.NewContent().GetContentListByAuthorId(authorInfo.Id, offset, limit, "id")
	return
}

//根据作者名字查询诗词总数
func (a *AuthorLogic) GetPoetryCountByFilter(authorName string) (count int, err error) {
	var authorInfo models.Author
	if authorInfo, err = a.authorModel.GetAuthorInfoByName(authorName); err != nil || authorInfo.Id == 0 {
		return
	}
	num, err := models.NewContent().GetContentCountByAuthorId(authorInfo.Id)
	count = int(num)
	return
}
