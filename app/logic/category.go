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
		data[category.Id] = &category
	}
	return
}
