/*
@Time : 2019/9/23 16:07
@Author : zxr
@File : dynasty
@Software: GoLand
*/
package logic

import (
	"poetry/app/models"
)

type dynastyLogic struct {
	model *models.Dynasty
}

func NewDynastyLogic() *dynastyLogic {
	return &dynastyLogic{
		model: models.NewDynasty(),
	}
}

func (d *dynastyLogic) GetDataById(id int) (data models.Dynasty, err error) {
	return d.model.GetDataById(id)
}

func (d *dynastyLogic) GetDataByIdArr(id []int) (data []models.Dynasty, err error) {
	if len(id) == 0 {
		return
	}
	return d.model.GetDataByIdArr(id)
}

//查询所有朝代列表
func (d *dynastyLogic) GetAll(offset, limit int) (data []models.Dynasty, err error) {
	return d.model.GetAll(offset, limit)
}

//根据朝代名字查询朝代信息
func (d *dynastyLogic) GetDynastyDataByName(dynastyName string) (dynastyRow models.Dynasty, err error) {
	return d.model.GetDynastyDataByName(dynastyName)
}

//根据作者数据获取朝代ID
func (d *dynastyLogic) GetDynastyIds(authorData map[int]models.Author) (dynastyList map[int]string) {
	dynastyList = make(map[int]string)
	var dynastyIds []int
	for _, author := range authorData {
		dynastyIds = append(dynastyIds, author.DynastyId)
	}
	if dynastyData, err := NewDynastyLogic().GetDataByIdArr(dynastyIds); err == nil {
		for _, dynastyVal := range dynastyData {
			dynastyList[dynastyVal.Id] = dynastyVal.DynastyName
		}
	}
	return dynastyList
}

//根据朝代名字搜索诗词列表
func (d *dynastyLogic) GetPoetryListByFilter(dynastyName string, offset, limit int) (data []models.Content, err error) {
	var (
		dynastyRow models.Dynasty
	)
	//1.先查出朝代ID
	if dynastyRow, err = d.model.GetDynastyDataByName(dynastyName); err != nil || dynastyRow.Id == 0 {
		return
	}
	//2.根据朝代ID写SQL关联诗词列表
	if data, err = models.NewContent().GetContentListByDynastyId(dynastyRow.Id, offset, limit); err != nil {
		return
	}
	return
}

//根据朝代名字查询诗词总数
func (d *dynastyLogic) GetPoetryCountByFilter(dynastyName string) (count int, err error) {
	var dynastyRow models.Dynasty
	if dynastyRow, err = d.model.GetDynastyDataByName(dynastyName); err != nil || dynastyRow.Id == 0 {
		return
	}
	count, err = models.NewContent().GetContentCountByDynastyId(dynastyRow.Id)
	return
}
