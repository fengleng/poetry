/*
@Time : 2019/10/17 19:29
@Author : zxr
@File : ancientCatalogue
@Software: GoLand
*/
package logic

import (
	"poetry/app/models"
	"poetry/config/define"
)

type AncientCatalogueLogic struct {
}

func NewAncientCatalogueLogic() *AncientCatalogueLogic {
	return &AncientCatalogueLogic{}
}

//根据bookId获取所有目录
func (a *AncientCatalogueLogic) GetAllCatalogByBookId(bookId int) (data []define.GuWenCatalogList, err error) {
	var (
		classList   []models.AnCatalogClass
		classIdList []int
		catalogList []models.AncientCatalogue
	)
	classList, err = models.NewAnCatalogClass().GetClassListByBookId(bookId)
	for _, class := range classList {
		classIdList = append(classIdList, int(class.Id))
	}
	if catalogList, err = models.NewAncientCatalogue().GetCatalogListByBookIdCids(bookId, classIdList); err != nil {
		return
	}
	if len(classList) == 0 {
		data = append(data, define.GuWenCatalogList{
			SubCatalog: catalogList,
		})
		return
	}
	for _, class := range classList {
		log := define.GuWenCatalogList{
			Id:        int(class.Id),
			ClassName: class.CatName,
			Sort:      class.Sort,
		}
		for _, catalog := range catalogList {
			if catalog.CatalogCatgoryId == class.Id {
				log.SubCatalog = append(log.SubCatalog, catalog)
			}
		}
		data = append(data, log)
	}
	return
}

//根据目录ID查询目录信息
func (a *AncientCatalogueLogic) GetDataById(logId int) (data models.AncientCatalogue, err error) {
	return models.NewAncientCatalogue().GetDataById(logId)
}

//根据bookId和id查询比id小的目录信息，用于获取上一章的内容
func (a *AncientCatalogueLogic) GetLogLtIdByBookId(bookId, id int64, offset, limit int) (data []models.AncientCatalogue, err error) {
	return models.NewAncientCatalogue().GetLogLtIdByBookId(bookId, id, offset, limit)
}

//根据bookId和id查询比id大的目录信息，用于获取下一章的内容
func (a *AncientCatalogueLogic) GetLogGtIdByBookId(bookId, id int64, offset, limit int) (data []models.AncientCatalogue, err error) {
	return models.NewAncientCatalogue().GetLogGtIdByBookId(bookId, id, offset, limit)
}

//根据目录分类ID查询目录分类信息
func (a *AncientCatalogueLogic) GetClassDataById(id int) (data models.AnCatalogClass, err error) {
	return models.NewAnCatalogClass().GetClassDataById(id)
}
