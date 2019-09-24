/*
@Time : 2019/9/20 17:26
@Author : zxr
@File : index_test
@Software: GoLand
*/
package test

import (
	"github.com/sirupsen/logrus"
	"poetry/app/logic"
	"testing"
)

func TestRecommend(t *testing.T) {
	contentData, _ := logic.NewRecommendLogic().GetSameDayPoetryData(0, 10)
	for _, c := range contentData.ContentArr {
		for _, tag := range c.Tags {
			logrus.Infof("%+v\n\n", tag)
		}
		//	logrus.Infof("%+v\n\n", c.Tags)

	}

	//根据诗词ID查询诗词表数据
	//logrus.Infoln("contentData:")
	//logrus.Infof("%+v\n\n\n", contentData)
	//logrus.Infoln("authorData:")
	//logrus.Infof("%+v\n", authorData)
	//logrus.Infoln("contentList:")
	//logrus.Infof("%+v\n", contentList)
	//logrus.Infoln("recommendData:")
	//logrus.Infof("%+v", r.recommendData)
	return

}
