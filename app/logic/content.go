/*
@Time : 2019/9/20 15:42
@Author : zxr
@File : content
@Software: GoLand
*/
package logic

import "poetry/app/models"

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
func (c *contentLogic) extractAuthorId(contentList []models.Content) (authorIds []int64) {
	authorIds = make([]int64, len(contentList))
	for k, content := range contentList {
		authorIds[k] = content.AuthorId
	}
	return
}

//根据诗词数据获取诗词ID
func (c *contentLogic) extractPoetryId(contentList []models.Content) (poetryIds []int) {
	poetryIds = make([]int, len(contentList))
	for k, content := range contentList {
		poetryIds[k] = content.Id
	}
	return
}

////根据诗词ID数组查询诗词所有详情信息
//func (c *contentLogic) GetPoetryDataByIdList(ids []int64) (data []models.Content, err error) {
//	if len(ids) == 0 {
//		return
//	}
//	var (
//		contentData []models.Content
//	)
//	//根据诗词ID查询诗词表数据
//	if contentData, err = c.contentModel.GetContentByIdList(ids); err != nil || len(contentData) == 0 {
//		return
//	}
//	//根据诗词ID查询作者表数据
//
//	//根据诗词ID查询诗词详情表数据
//
//	return
//}
