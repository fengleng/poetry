/*
@Time : 2019/9/19 17:47
@Author : zxr
@File : recommend
@Software: GoLand
*/
package logic

import (
	"poetry/app/bootstrap"
	"poetry/app/models"
	"poetry/config/define"
	"poetry/tools"
	"strconv"
)

//推荐数据服务
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
		tags         TagMp                 //诗词的分类标签信息
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
	//根据诗词ID查询分类标签表数据
	tags, _ = NewContentTagLogic().GetDataByPoetryId(poetryIdList)
	//将诗词数据，作者数据，朝代数据,分类整合一起
	contentData = r.ProcContentAuthorTagData(contentList, authorData, tags)
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
将诗词数据，作者数据，朝代数据,分类整合一起
*/
func (r *RecommendLogic) ProcContentAuthorTagData(contentList []models.Content, authorData map[int]models.Author, tags TagMp) (contentData define.ContentAll) {
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
			defineAuthor define.PoetryAuthor
			author       models.Author
			content      define.Content
		)
		oriContent := poetryText.Content
		poetryText.Content = tools.AddHtmlLabel(poetryText.Content)
		text.OriContent = oriContent
		text.PoetryInfo = poetryText
		text.LinkUrl = bootstrap.G_Conf.WebDomain + "/shiwen/" + strconv.FormatUint(uint64(poetryText.SourceUrlCrc32), 10) + ".html"
		author, _ = authorData[int(poetryText.AuthorId)]
		author.Id = poetryText.AuthorId
		defineAuthor.AuthorInfo = author
		defineAuthor.DynastyName = dynastyList[author.DynastyId]
		defineAuthor.AuthorLinkUrl = bootstrap.G_Conf.WebDomain + "/author/?type=author&value=" + author.Author
		defineAuthor.DynastyLinkUrl = bootstrap.G_Conf.WebDomain + "/search/?type=dynasty&cstr=" + defineAuthor.DynastyName
		content.PoetryText = text
		content.PoetryAuthor = defineAuthor
		content.Tags = tags[poetryText.Id]
		contentData.ContentArr[k] = &content
	}
	return
}
