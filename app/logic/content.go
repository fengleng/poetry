/*
@Time : 2019/9/20 15:42
@Author : zxr
@File : content
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

type contentLogic struct {
	contentModel *models.Content
}

func NewContentLogic() *contentLogic {
	return &contentLogic{
		contentModel: models.NewContent(),
	}
}

//根据诗词ID数组查询诗词表数据
func (c *contentLogic) GetContentByIdList(ids []int64) (data []models.Content, err error) {
	return c.contentModel.GetContentByIdList(ids)
}

//根据诗词数据获取作者ID
//func (c *contentLogic) extractAuthorId(contentList []models.Content) (authorIds []int64) {
//	authorIds = make([]int64, len(contentList))
//	for k, content := range contentList {
//		authorIds[k] = content.AuthorId
//	}
//	return
//}

//根据诗词数据获取诗词ID
//func (c *contentLogic) extractPoetryId(contentList []models.Content) (poetryIds []int) {
//	poetryIds = make([]int, len(contentList))
//	for k, content := range contentList {
//		poetryIds[k] = content.Id
//	}
//	return
//}

//根据sourceurl_crc32 查询正文数据
func (c *contentLogic) GetContentByCrc32Id(crc32Id uint32) (data models.Content, err error) {
	if crc32Id == 0 {
		return
	}
	return c.contentModel.GetContentByCrc32Id(crc32Id)
}

//根据诗词ID获取诗词的作者，分类，诗词具体内容等信息
func (c *contentLogic) GetPoetryContentAll(poetryIdList []int64) (contentData define.ContentAll, err error) {
	var (
		contentList []models.Content      //根据诗词ID查询出来的诗词数据
		authorIds   []int64               //作者ID集合
		authorData  map[int]models.Author //作者信息集合
		tags        TagMp                 //诗词的分类标签信息
	)
	//根据诗词ID查询诗词表数据
	if contentList, err = c.GetContentByIdList(poetryIdList); err != nil || len(contentList) == 0 {
		return
	}
	authorIds = ExtractAuthorId(contentList)
	//根据作者ID查询作者表数据
	if authorData, err = NewAuthorLogic().GetAuthorInfoByIds(authorIds); err != nil {
		return
	}
	//根据诗词ID查询分类标签表数据
	tags, _ = NewContentTagLogic().GetDataByPoetryId(poetryIdList)
	//将诗词数据，作者数据，朝代数据,分类整合一起
	contentData = c.ProcContentAuthorTagData(contentList, authorData, tags)
	return
}

//将诗词数据，作者数据，朝代数据,分类整合一起
func (c *contentLogic) ProcContentAuthorTagData(contentList []models.Content, authorData map[int]models.Author, tags TagMp) (contentData define.ContentAll) {
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
		text.LinkUrl = bootstrap.G_Conf.WebDomain + "/shiwen/" + strconv.FormatUint(uint64(poetryText.SourceUrlCrc32), 10) + define.UrlSuffix
		author, _ = authorData[int(poetryText.AuthorId)]
		author.Id = poetryText.AuthorId
		author.AuthorIntro = tools.TrimRightHtml(author.AuthorIntro)
		defineAuthor.AuthorInfo = author
		defineAuthor.DynastyName = dynastyList[author.DynastyId]
		defineAuthor.AuthorLinkUrl = bootstrap.G_Conf.WebDomain + "/author/poetryList?value=" + author.Author
		defineAuthor.AuthorDetailUrl = bootstrap.G_Conf.WebDomain + "/author/detail?value=" + author.Author
		defineAuthor.DynastyLinkUrl = bootstrap.G_Conf.WebDomain + "/search/shiwen/?type=dynasty&cstr=" + defineAuthor.DynastyName
		content.PoetryText = text
		content.PoetryAuthor = defineAuthor
		if tags == nil || len(tags) == 0 {
			content.Tags = nil
		} else {
			content.Tags = tags[poetryText.Id]
		}
		contentData.ContentArr[k] = &content
	}
	return
}

//根据作者信息查询作者诗词列表
func (c *contentLogic) GetPoetryListByAuthorId(authorInfo models.Author, offset, limit int, orderFiled string) (poetryListData define.ContentAll, err error) {
	var (
		tags         TagMp //诗词的分类标签信息
		poetryIdList []int64
		authorData   map[int]models.Author
		contentList  []models.Content
	)
	//诗词列表信息
	if contentList, err = c.contentModel.GetContentListByAuthorId(authorInfo.Id, offset, limit, orderFiled); err != nil {
		return
	}
	poetryIdList = make([]int64, len(contentList))
	for i, content := range contentList {
		poetryIdList[i] = int64(content.Id)
	}
	//根据诗词ID查询分类标签表数据
	tags, _ = NewContentTagLogic().GetDataByPoetryId(poetryIdList)
	//将诗词数据，朝代数据,分类整合一起
	authorData = make(map[int]models.Author, 1)
	authorData[int(authorInfo.Id)] = authorInfo
	poetryListData = c.ProcContentAuthorTagData(contentList, authorData, tags)
	return
}

//根据作者ID查询作者诗词总数
func (c *contentLogic) GetContentCountByAuthorId(authorId int64) (count int, err error) {
	var cNum int64
	if cNum, err = c.contentModel.GetContentCountByAuthorId(authorId); err != nil {
		return
	}
	count = int(cNum)
	return
}
