/*
@Time : 2019/9/25 15:55
@Author : zxr
@File : nocde
@Software: GoLand
*/
package logic

import "poetry/app/models"

//诗词译文，注释服务
type nocdnLogic struct {
	contentLogic *contentLogic
}

func NewNocdnLogic() *nocdnLogic {
	return &nocdnLogic{
		contentLogic: NewContentLogic(),
	}
}

//查询诗词赏析信息，注释信息
func (n *nocdnLogic) GetPoetryTextTrans(crc32Id uint32, typeStr string) (err error) {
	var (
		poetryData models.Content
	)
	if poetryData, err = n.contentLogic.GetContentByCrc32Id(crc32Id); err != nil || poetryData.Id == 0 {
		return
	}

	return
}
