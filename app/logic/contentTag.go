/*
@Time : 2019/9/24 16:51
@Author : zxr
@File : contentTag
@Software: GoLand
*/
package logic

import (
	"poetry/app/models"
)

type TagMp map[int][]*models.Category

type contentTagLogic struct {
	contentTagModel *models.ContentTag
}

func NewContentTagLogic() *contentTagLogic {
	return &contentTagLogic{
		contentTagModel: models.NewContentTag(),
	}
}

//根据诗词ID集合查询tag数据
func (c *contentTagLogic) GetDataByPoetryId(poetryIds []int64) (data TagMp, err error) {
	var (
		tags         []models.ContentTag
		tagIds       []int
		categoryData MpCategory
	)
	if tags, err = c.contentTagModel.GetDataByPoetryId(poetryIds); err != nil {
		return
	}
	data = make(TagMp)
	for _, tag := range tags {
		tagIds = append(tagIds, tag.CategoryId)
	}
	if categoryData, err = NewCategoryLogic().GetDataByIds(tagIds); err != nil {
		return
	}
	for _, tag := range tags {
		if category, ok := categoryData[tag.CategoryId]; ok {
			data[tag.PoetryId] = append(data[tag.PoetryId], category)
		}
	}
	return
}
