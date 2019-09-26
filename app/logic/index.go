/*
@Time : 2019/9/26 11:24
@Author : zxr
@File : index
@Software: GoLand
*/
package logic

import "poetry/config/define"

type IndexLogic struct {
}

func NewIndexLogic() *IndexLogic {
	return &IndexLogic{}
}

//获取首页推荐数据
func (i *IndexLogic) GetSameDayRecommendPoetryData(offset, limit int) (contentData define.ContentAll, err error) {
	recommendLogic := NewRecommendLogic()
	contentData, err = recommendLogic.GetRecommendData(offset, limit)
	return
}
