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

type MpCategory map[int]*models.Category

type categoryLogic struct {
	categoryModel *models.Category
}

func NewCategoryLogic() *categoryLogic {
	return &categoryLogic{
		categoryModel: models.NewCategory(),
	}
}

//根据诗词ID集合查询数据
func (c *categoryLogic) GetDataByIds(ids []int) (data MpCategory, err error) {
	var (
		categoryData []models.Category
	)
	if categoryData, err = c.categoryModel.GetDataByIds(ids); err != nil {
		return
	}
	data = make(MpCategory, len(categoryData))
	for _, category := range categoryData {
		tempCate := category
		data[category.Id] = &tempCate
	}
	return
}

//根据show_position查询所有分类
func (c *categoryLogic) GetCateByPositionLimit(showPosition, offset, limit int) (data []models.Category, err error) {
	return models.NewCategory().GetCateByPositionLimit(showPosition, offset, limit)
}

//根据分类名字搜索诗词列表
func (c *categoryLogic) GetPoetryListByFilter(categoryName string, offset, limit int) (data []models.Content, err error) {
	var categoryInfo models.Category
	//查询分类信息
	if categoryInfo, err = c.categoryModel.GetCategoryInfoByCateName(categoryName); err != nil || categoryInfo.Id == 0 {
		return
	}
	//写SQL 根据分类ID查询诗词列表
	data, err = models.NewContent().GetContentListByCategoryId(categoryInfo.Id, offset, limit)
	return
}

//根据分类名字查询诗词总数
func (c *categoryLogic) GetPoetryCountByFilter(categoryName string) (count int, err error) {
	var categoryInfo models.Category
	//查询分类信息
	if categoryInfo, err = c.categoryModel.GetCategoryInfoByCateName(categoryName); err != nil || categoryInfo.Id == 0 {
		return
	}
	count, err = models.NewContentTag().GetCountByCategoryId(categoryInfo.Id)
	return
}
