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
