/*
@Time : 2019/9/4 18:40
@Author : zxr
@File : catalogCategoryModel
@Software: GoLand
*/
package models

import (
	"github.com/astaxie/beego/orm"
)

//古籍-书名目录类
type AncientCatalogue struct {
	Id               int64  `orm:"column(id);auto"`
	BookId           int64  `orm:"column(book_id)"`
	CatalogTitle     string `orm:"column(catalog_title)"`
	CatalogCatgoryId int64  `orm:"column(catalog_catgory_id)"`
	LinkUrl          string `orm:"column(link_url)"`
	Sort             int    `orm:"column(sort)"`
	AddDate          int64  `orm:"column(add_date)"`
}

func init() {
	orm.RegisterModel(new(AncientCatalogue))
}

func (a *AncientCatalogue) TableName() string {
	return AncientCatalogueTable
}

func NewAncientCatalogue() *AncientCatalogue {
	return new(AncientCatalogue)
}

//根据URL数组 查询目录列表
func (a *AncientCatalogue) GetCatalogListByPaths(paths []string) (data []AncientCatalogue, err error) {
	fields := []string{"id", "book_id", "catalog_title"}
	_, err = orm.NewOrm().QueryTable(AncientCatalogueTable).Filter("link_url__in", paths).All(&data, fields...)
	return
}
