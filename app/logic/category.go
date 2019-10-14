/*
@Time : 2019/9/24 16:51
@Author : zxr
@File : contentTag
@Software: GoLand
*/
package logic

import (
	"poetry/app/models"
	"poetry/config/define"
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
	if categoryInfo, err = c.categoryModel.GetCategoryInfoByCateName(categoryName, define.PoetryShowPosition); err != nil || categoryInfo.Id == 0 {
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
	if categoryInfo, err = c.categoryModel.GetCategoryInfoByCateName(categoryName, define.PoetryShowPosition); err != nil || categoryInfo.Id == 0 {
		return
	}
	count, err = models.NewContentTag().GetCountByCategoryId(categoryInfo.Id)
	return
}

//根据分类名字查询分类信息
func (c *categoryLogic) GetCateInfoByName(categoryName string, showPosition int) (data models.Category, err error) {
	data, err = models.NewCategory().GetCategoryInfoByCateName(categoryName, showPosition)
	return
}

//根据Pid查询子分类,如果pid为0，则查询所有子分类
func (c *categoryLogic) GetAllSubCateData(pid, showPosition, offset, limit int) (data []models.Category, err error) {
	return models.NewCategory().GetAllSubCateData(pid, showPosition, offset, limit)
}

//根据分类名字和PID查询分类信息
func (c *categoryLogic) GetCateInfoByNameAndPid(pid int, cateName string) (data models.Category, err error) {
	return models.NewCategory().GetCateInfoByNameAndPid(pid, cateName)
}

//根据 category,pid 查询子分类 查二级分类
func (c *categoryLogic) FindSubCateByPid(cateMap []models.Category, pid int) (catId []int, subCategory []models.Category) {
	catId = make([]int, 0)
	for _, cate := range cateMap {
		if cate.Pid == pid {
			subCategory = append(subCategory, cate)
			catId = append(catId, cate.Id)
		}
	}
	return
}

//在[]models.Category中根据名字找出对应的分类信息
func (c *categoryLogic) FindCateListByName(cateList []models.Category, name string) (cateInfo models.Category) {
	for _, cate := range cateList {
		if cate.CatName == name {
			cateInfo = cate
			break
		}
	}
	return
}

//处理子分类，把父分类名字加在子分类上,每个父分类显示 10个子分类
func (c *categoryLogic) ProcSubCategory(topCategory []models.Category, subAllCate []models.Category) (subList []*define.PCategory) {
	topCateMap := make(map[int]models.Category, len(topCategory))
	for _, category := range topCategory {
		topCateMap[category.Id] = category
	}
	topCateArrMap := make(map[int][]*define.PCategory)
	limit := 10
	for _, category := range subAllCate {
		cate := &define.PCategory{
			Model: category,
			PName: topCateMap[category.Pid].CatName,
			Pid:   category.Pid,
		}
		if len(topCateArrMap[category.Pid]) > limit {
			continue
		}
		topCateArrMap[category.Pid] = append(topCateArrMap[category.Pid], cate)
	}
	for _, category := range topCateArrMap {
		subList = append(subList, category...)
	}
	return
}
