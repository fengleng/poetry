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
	recommendData []models.Recommend //查询的推荐数据
	contentLogic  *contentLogic
}

func NewRecommendLogic() *RecommendLogic {
	return &RecommendLogic{
		contentLogic: NewContentLogic(),
	}
}

//获取当前推荐的诗词数据
func (r *RecommendLogic) GetSameDayPoetryData(offset, limit int) (err error) {
	var (
		contentList  []models.Content //根据诗词ID查询出来的诗词数据
		poetryIdList []int64          //诗词ID集合
		authorIds    []int64          //作者ID集合
		authorData   []models.Author  //作者信息集合
	)
	defer func() {
		r.contentLogic = nil
		r.recommendData = nil
	}()

	if r.recommendData, err = models.NewRecommendModel().GetSameDayData(offset, limit); err != nil || len(r.recommendData) == 0 {
		return
	}
	poetryIdList = r.extractPoetryId()
	//根据诗词ID查询诗词表数据
	if contentList, err = r.contentLogic.GetContentByIdList(poetryIdList); err != nil || len(contentList) == 0 {
		return
	}
	authorIds = r.contentLogic.extractAuthorId(contentList)
	//根据诗词ID查询作者表数据
	if authorData, err = NewAuthorLogic().GetAuthorInfoByIds(authorIds); err != nil {
		return
	}

	//根据诗词ID查询诗词表数据
	logrus.Infoln("authorData:")
	logrus.Infof("%+v\n", authorData)
	logrus.Infoln("contentList:")
	logrus.Infof("%+v\n", contentList)
	logrus.Infoln("recommendData:")
	logrus.Infof("%+v", r.recommendData)
	return
}

//从推荐数据中提取诗词ID
func (r *RecommendLogic) extractPoetryId() (poetryIdList []int64) {
	poetryIdList = make([]int64, len(r.recommendData))
	for k, recommend := range r.recommendData {
		poetryIdList[k] = recommend.PoetryId
	}
	return
}
