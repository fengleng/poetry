/*
@Time : 2019/9/19 17:47
@Author : zxr
@File : recommend
@Software: GoLand
*/
package logic

import (
	"poetry/app/models"
	"poetry/config/define"
)

//推荐数据服务
type RecommendLogic struct {
	recommendData []models.Recommend //查询的推荐数据
	contentLogic  *contentLogic
	authorLogic   *AuthorLogic
}

func NewRecommendLogic() *RecommendLogic {
	return &RecommendLogic{
		contentLogic: NewContentLogic(),
		authorLogic:  NewAuthorLogic(),
	}
}

//根据offset,limit 偏移量查询推荐数据
func (r *RecommendLogic) GetRecommendByOffset(offset, limit int) (recommendData []models.Recommend, err error) {
	recommendData, err = models.NewRecommendModel().GetRecommendByOffset(offset, limit)
	return
}

//获取当前推荐的诗词数据,按日期倒序排列
func (r *RecommendLogic) GetSameDayRecommendPoetryData(offset, limit int) (contentData define.ContentAll, err error) {
	var poetryIdList []int64 //诗词ID集合
	defer func() {
		r.contentLogic = nil
		r.authorLogic = nil
		r.recommendData = nil
	}()
	if r.recommendData, err = r.GetRecommendByOffset(offset, limit); err != nil || len(r.recommendData) == 0 {
		return
	}
	poetryIdList = r.extractPoetryId()
	contentData, err = r.contentLogic.GetPoetryContentAll(poetryIdList)
	return
}

//获取推荐总数
func (r *RecommendLogic) GetRecommendCount() int {
	var (
		count int64
		err   error
	)
	defer func() {
		r.contentLogic = nil
		r.authorLogic = nil
	}()
	if count, err = models.NewRecommendModel().GetCount(); err != nil {
		return 0
	}
	return int(count)
}

//从推荐数据中提取诗词ID
func (r *RecommendLogic) extractPoetryId() (poetryIdList []int64) {
	poetryIdList = make([]int64, len(r.recommendData))
	for k, recommend := range r.recommendData {
		poetryIdList[k] = recommend.PoetryId
	}
	return
}
