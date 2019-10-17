/*
@Time : 2019/9/27 12:19
@Author : zxr
@File : ancientBook
@Software: GoLand
*/
package logic

import (
	"poetry/app/models"
	"poetry/config/define"
	"strings"
)

type AncientClassifyLogic struct {
}

func NewAncientClassify() *AncientClassifyLogic {
	return &AncientClassifyLogic{}
}

//根据PID查询分类数据并分页
func (a *AncientClassifyLogic) GetDataLimitByPid(pid int, offset, limit int) (data []models.AncientClassify, err error) {
	return models.NewAncientClassify().GetDataLimitByPid(pid, offset, limit)
}

//根据分类名称查询数据
func (a *AncientClassifyLogic) GetCategoryDataByName(catName string) (data models.AncientClassify, err error) {
	return models.NewAncientClassify().GetCategoryDataByName(catName)
}

//获取古文所有分类数据
func (a *AncientClassifyLogic) GetAllClassify(subLimit int) (allClassify []*define.GuWenClassify, err error) {
	var classifyData []models.AncientClassify

	if classifyData, err = a.GetDataLimitByPid(0, 0, 10); err != nil || len(classifyData) == 0 {
		return
	}
	allClassify = make([]*define.GuWenClassify, len(classifyData))
	for i, classify := range classifyData {
		allClassify[i] = &define.GuWenClassify{
			Id:           classify.Id,
			ClassifyName: strings.Trim(classify.CatName, "："),
			Pid:          int(classify.Pid),
			Sort:         classify.Sort,
		}
	}
	for i, classify := range allClassify {
		var classData []models.AncientClassify
		if classData, err = a.GetDataLimitByPid(classify.Id, 0, subLimit); err != nil {
			continue
		}
		allClassify[i].SubClassify = classData
	}
	return
}

//根据分类名在所有分类中查找信息，返回具体子分类ID
func (a *AncientClassifyLogic) FindCatIdListByCateName(allClassify []*define.GuWenClassify, classStr string) (subClassIdList []int) {
	if len(allClassify) == 0 || len(classStr) == 0 {
		return
	}
	for _, classify := range allClassify {
		if strings.EqualFold(classify.ClassifyName, classStr) {
			subClassIdList = a.FindCatIdListByPid(allClassify, classify.Id)
			break
		}
		for _, subClass := range classify.SubClassify {
			if strings.EqualFold(subClass.CatName, classStr) {
				subClassIdList = []int{subClass.Id}
				break
			}
		}
	}
	return
}

//在所有分类中，根据pid查找所有子分类Id
func (a *AncientClassifyLogic) FindCatIdListByPid(allClassify []*define.GuWenClassify, pid int) (subClassIdList []int) {
	if pid == 0 {
		return
	}
	for _, classify := range allClassify {
		for _, subClass := range classify.SubClassify {
			if int(subClass.Pid) == pid {
				subClassIdList = append(subClassIdList, subClass.Id)
			}
		}
	}
	return
}
