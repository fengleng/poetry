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
	fields := []string{"id", "book_id", "catalog_title", "link_url"}
	_, err = orm.NewOrm().QueryTable(AncientCatalogueTable).Filter("link_url__in", paths).All(&data, fields...)
	return
}

//根据bookId和目录分类ID查询目录具体列表
func (a *AncientCatalogue) GetCatalogListByBookIdCids(bookId int, categoryIds []int) (data []AncientCatalogue, err error) {
	fields := []string{"id", "book_id", "catalog_title", "link_url", "catalog_catgory_id"}
	if len(categoryIds) > 0 {
		_, err = orm.NewOrm().QueryTable(AncientCatalogueTable).Filter("book_id", bookId).Filter("catalog_catgory_id__in", categoryIds).All(&data, fields...)
	} else {
		_, err = orm.NewOrm().QueryTable(AncientCatalogueTable).Filter("book_id", bookId).All(&data, fields...)
	}
	return
}

//根据目录ID查询目录信息
func (a *AncientCatalogue) GetDataById(id int) (data AncientCatalogue, err error) {
	fields := []string{"id", "book_id", "catalog_title", "link_url", "catalog_catgory_id"}
	_, err = orm.NewOrm().QueryTable(AncientCatalogueTable).Filter("id", id).All(&data, fields...)
	return
}

//根据bookId和id查询比id小的目录信息，用于获取上一章的内容
func (a *AncientCatalogue) GetLogLtIdByBookId(bookId, id int64, offset, limit int) (data []AncientCatalogue, err error) {
	fields := []string{"id", "book_id", "catalog_title", "link_url", "catalog_catgory_id"}
	_, err = orm.NewOrm().QueryTable(AncientCatalogueTable).Filter("book_id", bookId).Filter("id__lt", id).Limit(limit, offset).OrderBy("-id").All(&data, fields...)
	return
}

//根据bookId和id查询比id大的目录信息，用于获取下一章的内容
func (a *AncientCatalogue) GetLogGtIdByBookId(bookId, id int64, offset, limit int) (data []AncientCatalogue, err error) {
	fields := []string{"id", "book_id", "catalog_title", "link_url", "catalog_catgory_id"}
	_, err = orm.NewOrm().QueryTable(AncientCatalogueTable).Filter("book_id", bookId).Filter("id__gt", id).Limit(limit, offset).OrderBy("id").All(&data, fields...)
	return
}
