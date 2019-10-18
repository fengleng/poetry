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

//古籍-正文内容
type AncientBookContent struct {
	Id             int64  `orm:"column(id);auto"`
	BookId         int64  `orm:"column(book_id)"`
	CatalogId      int64  `orm:"column(catalog_id)"`
	Content        string `orm:"column(content)"`
	Translation    string `orm:"column(translation)"`
	TranslationId  int    `orm:"column(translation_id)"`
	TranslationUrl string `orm:"column(translation_url)"`
	AuthorId       int64  `orm:"column(author_id)"`
	SongUrl        string `orm:"column(song_url)"`
	SongFilePath   string `orm:"column(song_file_path)"`
	AddDate        int64  `orm:"column(add_date)"`
	UpdateDate     int64  `orm:"column(update_date)"`
}

func init() {
	orm.RegisterModel(new(AncientBookContent))
}

func (a *AncientBookContent) TableName() string {
	return AncientBookContentTable
}

func NewAncientBookContent() *AncientBookContent {
	return new(AncientBookContent)
}

var bookContentFields = []string{"id", "book_id", "catalog_id", "content", "translation", "author_id", "song_url", "song_file_path"}

//根据目录ID查询正文内容
func (a *AncientBookContent) GetBookContentByCataLogId(id int) (data AncientBookContent, err error) {
	_, err = orm.NewOrm().QueryTable(AncientBookContentTable).Filter("catalog_id", id).All(&data, bookContentFields...)
	return
}

//根据ID查询正文内容
func (a *AncientBookContent) GetBookContentById(id int) (data AncientBookContent, err error) {
	_, err = orm.NewOrm().QueryTable(AncientBookContentTable).Filter("id", id).All(&data, bookContentFields...)
	return
}
