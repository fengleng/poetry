/*
@Time : 2019/9/19 17:47
@Author : zxr
@File : recommend
@Software: GoLand
*/
package logic

import (
	"poetry/app/models"
	"poetry/config/define"
)

type RecommendLogic struct {
	recommendData []models.Recommend //查询的推荐数据
	contentLogic  *contentLogic
	authorLogic   *AuthorLogic
}

func NewRecommendLogic() *RecommendLogic {
	return &RecommendLogic{
		contentLogic: NewContentLogic(),
		authorLogic:  NewAuthorLogic(),
	}
}

//获取当前推荐的诗词数据
func (r *RecommendLogic) GetSameDayPoetryData(offset, limit int) (contentData define.ContentAll, err error) {
	var (
		contentList  []models.Content      //根据诗词ID查询出来的诗词数据
		poetryIdList []int64               //诗词ID集合
		authorIds    []int64               //作者ID集合
		authorData   map[int]models.Author //作者信息集合
	)
	defer func() {
		r.contentLogic = nil
		r.authorLogic = nil
		r.recommendData = nil
	}()

	if r.recommendData, err = models.NewRecommendModel().GetSameDayData(offset, limit); err != nil || len(r.recommendData) == 0 {
		return
	}
	poetryIdList = r.extractPoetryId()
	//根据诗词ID查询诗词表数据
	if contentList, err = r.contentLogic.GetContentByIdList(poetryIdList); err != nil || len(contentList) == 0 {
		return
	}
	authorIds = r.contentLogic.extractAuthorId(contentList)
	//根据诗词ID查询作者表数据
	if authorData, err = r.authorLogic.GetAuthorInfoByIds(authorIds); err != nil {
		return
	}
	contentData = r.ProcContentAuthorData(contentList, authorData)
	return contentData, nil
}

//从推荐数据中提取诗词ID
func (r *RecommendLogic) extractPoetryId() (poetryIdList []int64) {
	poetryIdList = make([]int64, len(r.recommendData))
	for k, recommend := range r.recommendData {
		poetryIdList[k] = recommend.PoetryId
	}
	return
}

/**
将诗词数据，作者数据，朝代数据整合一起
*/
func (r *RecommendLogic) ProcContentAuthorData(contentList []models.Content, authorData map[int]models.Author) (contentData define.ContentAll) {
	dynastyList := NewDynastyLogic().GetDynastyIds(authorData)
	contentData.ContentArr = make([]*define.Content, len(contentList))
	defer func() {
		dynastyList = nil
		contentList = nil
		authorData = nil
	}()
	for k, poetryText := range contentList {
		var (
			text         define.PoetryText
			defineAuthor define.Author
			author       models.Author
			content      define.Content
		)
		author, _ = authorData[int(poetryText.AuthorId)]
		text.Id = poetryText.Id
		text.SourceUrl = poetryText.SourceUrl
		text.Sort = poetryText.Sort
		text.AuthorId = poetryText.AuthorId
		text.Content = poetryText.Content
		text.Title = poetryText.Content
		text.CreatBackId = poetryText.CreatBackId
		text.GenreId = poetryText.GenreId
		defineAuthor.Id = text.AuthorId
		defineAuthor.Author = author.Author
		defineAuthor.SourceUrl = author.SourceUrl
		defineAuthor.DynastyId = author.DynastyId
		defineAuthor.AuthorIntro = author.AuthorIntro
		defineAuthor.PhotoUrl = author.PhotoUrl
		defineAuthor.WorksUrl = author.WorksUrl
		defineAuthor.Pinyin = author.Pinyin
		defineAuthor.Acronym = author.Acronym
		defineAuthor.IsRecommend = author.IsRecommend
		defineAuthor.AuthorsId = author.AuthorsId
		defineAuthor.AuthorTitle = author.AuthorTitle
		defineAuthor.PhotoFileName = author.PhotoFileName
		defineAuthor.PoetryCount = author.PoetryCount
		defineAuthor.DynastyName = dynastyList[author.DynastyId]
		content.PoetryText = text
		content.Author = defineAuthor
		contentData.ContentArr[k] = &content
	}
	return
}
