/*
@Time : 2019/9/19 17:47
@Author : zxr
@File : recommend
@Software: GoLand
*/
package logic

import (
	"github.com/sirupsen/logrus"
	"poetry/app/models"
)

type RecommendLogic struct {
}

func NewRecommendLogic() *RecommendLogic {
	return &RecommendLogic{}
}

//获取当前推荐的诗词数据
func (r *RecommendLogic) GetSameDayPoetryData(offset, limit int) (err error) {
	var (
		recommendData []models.Recommend
		poetryIdList  []int64
	)
	if recommendData, err = models.NewRecommendModel().GetSameDayData(offset, limit); err != nil {
		return
	}
	poetryIdList = make([]int64, len(recommendData))
	for k, recommend := range recommendData {
		poetryIdList[k] = recommend.PoetryId
	}

	//根据诗词ID查询诗词表数据
	logrus.Infof("%+v\n", poetryIdList)
	logrus.Infof("%+v", recommendData)
	return
}
