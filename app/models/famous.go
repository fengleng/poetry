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

//名句Model
type Famous struct {
	Id           int64  `orm:"column(id);auto"`
	CatId        int    `orm:"column(cat_id)"`
	Content      string `orm:"column(content)"`
	ContentCrc32 uint32 `orm:"column(content_crc32)"`
	PoetryTitle  string `orm:"column(poetry_title)"`
	PoetryId     int64  `orm:"column(poetry_id)"`
	AuthorId     int64  `orm:"column(author_id)"`
	Sort         int    `orm:"column(sort)"`
	SourceUrl    string `orm:"column(source_url)"`
	SourceCrc32  uint32 `orm:"column(source_crc32)"`
	AddDate      int64  `orm:"column(add_date)"`
	UpdateDate   int64  `orm:"column(update_date)"`
}

func init() {
	orm.RegisterModel(new(Famous))
}

func (f *Famous) TableName() string {
	return FamousSentenceTable
}

func NewFamous() *Famous {
	return new(Famous)
}
