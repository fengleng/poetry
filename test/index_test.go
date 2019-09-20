/*
@Time : 2019/9/20 17:26
@Author : zxr
@File : index_test
@Software: GoLand
*/
package test

import (
	"poetry/app/logic"
	"testing"
)

func TestRecommend(t *testing.T) {
	logic.NewRecommendLogic().GetSameDayPoetryData(0, 10)
}
